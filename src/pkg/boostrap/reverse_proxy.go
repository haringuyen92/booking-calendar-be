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
	//gin.Any("/api/staffs/*proxyPath", reverseProxy)
	//gin.Any("/api/courses/*proxyPath", reverseProxy)
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
	originalDirector := proxy.Director

	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/api/"+serviceName)
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Host = remote.Host
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		// Xóa tất cả các CORS header từ response của service được proxy
		resp.Header.Del("Access-Control-Allow-Origin")
		resp.Header.Del("Access-Control-Allow-Credentials")
		resp.Header.Del("Access-Control-Allow-Methods")
		resp.Header.Del("Access-Control-Allow-Headers")
		resp.Header.Del("Access-Control-Expose-Headers")
		return nil
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
