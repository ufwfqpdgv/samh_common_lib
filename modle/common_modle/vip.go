package common_modle

type VipLevelCode int

const (
	VipLevelCode_None VipLevelCode = iota
	VipLevelCode_Gold
	VipLevelCode_Diamond
)

var vipLevelCodeToMsg = map[VipLevelCode]string{
	VipLevelCode_Gold:    "黄金会员",
	VipLevelCode_Diamond: "钻石会员",
}

func FetchVipLevelCodeToMsg(code VipLevelCode) string {
	return vipLevelCodeToMsg[code]
}

type VipTypeCode int

const (
	VipTypeCode_Gold VipTypeCode = iota + 1
	VipTypeCode_Diamond
	VipTypeCode_Activity
)
