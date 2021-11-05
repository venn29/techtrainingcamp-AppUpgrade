package deploy_service

import (
	"GaryReleaseProject/src/cache"
	"GaryReleaseProject/src/model"
	"github.com/gin-gonic/gin"
)


func DeployRule(c *gin.Context){
	// 前端接受的信息处理成Rule结构体
}

// 添加一条新规则
func addRule(rule *model.Rule){
	cache.AddRule(rule)
}

// 修改规则状态
func modifyRule(ruleID int,newStatus int){
	cache.ModifyRule(ruleID,newStatus)
}