package store_routers

import "github.com/gin-gonic/gin"

func RegisterRouters(
	group *gin.RouterGroup,
	rSetting *SettingRouter,
	rStore *StoreRouter,
) {
	rSetting.Register(group)
	rStore.Register(group)
}
