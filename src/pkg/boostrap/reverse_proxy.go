package boostrap

import (
	"booking-calendar-server-backend/internal/core/constant"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ReverseProxy(gin *gin.Engine) {
	gin.Any("/api/stores/*proxyPath", reverseProxy)
	gin.Any("/api/bookings/*proxyPath", reverseProxy)
	gin.Any("/api/users/*proxyPath", reverseProxy)
}

func reverseProxy(c *gin.Context) {
	serviceName := extractServiceName(c.Request.URL.Path)
	target, err := forwardTo(serviceName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	}
	remote, err := url.Parse(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse remote URL"})
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		originalPath := req.URL.Path
		originalRawQuery := req.URL.RawQuery

		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = originalPath
		req.URL.RawQuery = originalRawQuery

		req.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))
		req.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
	if err := recover(); err != nil {
		fmt.Printf("Proxy error: %v\n", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Proxy error occurred"})
	}
}

func extractServiceName(path string) string {
	parts := strings.Split(path, "/")
	apiIndex := -1
	for i, part := range parts {
		if part == "api" {
			apiIndex = i
			break
		}
	}
	if apiIndex == -1 || apiIndex+1 >= len(parts) {
		return ""
	}
	return parts[apiIndex+1]
}

func forwardTo(fromURL string) (string, error) {
	service := constant.ServicesProperties[constant.ServiceID(fromURL)]
	if service.IsValid() {
		return fmt.Sprintf("http://%s:%d", service.Host, service.Port), nil
	}
	return "", errors.New("invalid service")

}
