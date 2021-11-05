package update_service

import (
	"GaryReleaseProject/src/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func Hit(c *gin.Context) {
	var respUrl string
	// 后期Hit函数的参数可能会变为url；
	// cast.ToInt : nil string will be trans to 0
	// 要有异常处理，Query要不要设置默认值？
	// url有两个“userDID会怎么样？”...
	appVersion := c.Query("appVersion")
	userDID := c.Query("userDID")
	// 对应数据库操作
	rules := model.GetAllRules()

	for index := 0; index < len(*rules); index++ {

		if cast.ToInt(userDID) < (*rules)[index].MinUserDID || cast.ToInt(userDID) > (*rules)[index].MaxUserDID {
			continue
		}

		if cast.ToInt(appVersion) < (*rules)[index].MinVersion || cast.ToInt(appVersion) > (*rules)[index].MaxVersion {
			continue
		}

		respUrl = (*rules)[index].GrayLink
		break
	}

	c.JSON(200, gin.H{"downloadUrl": respUrl})
}

