package cache

import (
	"GaryReleaseProject/src/model"
	"GaryReleaseProject/src/update_service"
	"sync"
)

type RMessage struct {
	// message return ro App
	DownloadUrl string `json:"download_url"`
	UpdateVersionCode string `json:"update_version_code"`
	Md5 string `json:"md5"`
	Title string `json:"title"`
	UpdateTips string `json:"update_tips"`
}
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

func MatchRule(cr *update_service.CReport) *RMessage{
	var rm RMessage
	//
	ruleIDs := getAllRuleIds(cr)
	for i := 0; i < len(ruleIDs); i++ {
		// 默认其他都比完了，只剩下白名单
	}
	return &rm
}

func getAllRuleIds(cr *update_service.CReport)[]int{
	// 在数据库查询所有符合条件的ruleid
	// 返回切片要求一定是优先级顺序
	return nil
}