package model

type Rule struct {
	// rule condition
	Platform string `json:"platform"`
	DownloadUrl string `json:"download_url"`
	UpdateVersionCode string `json:"update_version_code"`
	MD5 string `json:"md5"`
	DeviceIdList string `json:"device_id_list"`
	MaxUpdateVersionCode string `json:"max_update_version_code"`
	MinUpdateVersionCode string `json:"min_update_version_code"`
	MaxOsApi int `json:"max_os_api"`
	MinOsApi int `json:"min_os_api"`
	CpuArch int `json:"cpu_arch"`
	Channel string `json:"channel"`
	Title string `json:"title"`
	UpdateTips string `json:"update_tips"`
	//新添加的字段
	UpdateVersionCodeInt64 int64
	MaxUpdateVersionCodeInt64 int64
	MinUpdateVersionCodeInt64 int64
	RuleStatus int
}

func VersionToInt64(version string) int64{
	return 0
}
