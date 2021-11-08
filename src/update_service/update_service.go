package update_service

import (
	"GaryReleaseProject/src/cache"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}


func DealCRport(c *gin.Context) {

	cr:= cache.CReport{
		DevicePlatform : c.Query("device_platform"),
		DeviceId : c.Query("device_id"),
		// ToInt会把空串转为0
		OsApi : cast.ToInt(c.Query("os_api")),
		Channel : c.Query("channel"),
		UpdateVersionCode : c.Query("update_version_code"),
		CpuArch : cast.ToInt(c.Query("cpu_arch")),
	}

	rm := cache.MatchRule(&cr)
	c.JSON(200,*rm)
}

