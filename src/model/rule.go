package model

type Rule struct {
	// rule condition
	RuleID int `json:"rule_id"`
	MinVersion int `json:"min_version"`
	MaxVersion int `json:"max_version"`
	MinUserDID int `json:"min_user_did"`
	MaxUserDID int `json:"max_user_did"`
	// 需要补全
	// apk or ipa file link
	GrayLink string `json:"gray_link"`
}

func GetAllRules() *[]Rule {
	rules := []Rule{}

	rules = append(rules, Rule{
		MinVersion: 10,
		MaxVersion: 20,
		MinUserDID: 10,
		MaxUserDID: 20,
		GrayLink:   "https://dldir1.qq.com/weixin/android/weixin001.apk",
	})

	rules = append(rules, Rule{
		MinVersion: 20,
		MaxVersion: 33,
		MinUserDID: 20,
		MaxUserDID: 33,
		GrayLink:   "https://dldir1.qq.com/weixin/android/weixin002.apk",
	})

	return &rules

}
