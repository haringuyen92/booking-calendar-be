package user_routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(
	group *gin.RouterGroup,
	rUser *UserRouter,
	rMessage *MessageRouter,
) {
	rUser.Register(group)
	rMessage.Register(group)
}
