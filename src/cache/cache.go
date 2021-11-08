package cache

import (
	"GaryReleaseProject/src/model"
	"sync"
)



const (
	CaCheSize = 100
)

var(
	ruleCache sync.Map
)

func InitCache(){
	// 预热cache，读取CaCheSize个规则
}

func AddRule(rule *model.Rule){
	// 首先检查Status字段是否为0
}

func ModifyRule(ruleID int, newStatus int){
	// 别忘记检查Status字段是否合法
}

func MatchRule(cr *model.CReport) *model.RMessage{
	var rm model.RMessage
	//
	ruleIDs := getAllRuleIds(cr)
	for i := 0; i < len(ruleIDs); i++ {
		// 默认其他都比完了，只剩下白名单
	}
	return &rm
}

func getAllRuleIds(cr *model.CReport)[]int{
	// 在数据库查询所有符合条件的ruleid
	// 返回切片要求一定是优先级顺序
	return nil
}