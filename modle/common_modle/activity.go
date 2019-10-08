package common_modle

type Activity struct {
	ActivityId int64  `json:"activity_id" xorm:"not null pk autoincr BIGINT(20)"`
	Content    string `json:"content" xorm:"default '' VARCHAR(1000)"`
	CreateTime int64  `json:"create_time" xorm:"default 0 BIGINT(20)"`
	EndTime    int64  `json:"end_time" xorm:"default 0 BIGINT(20)"`
	LinkUrl    string `json:"link_url" xorm:"default '活动链接，一般是h5网页' VARCHAR(200)"`
	StartTime  int64  `json:"start_time" xorm:"default 0 BIGINT(20)"`
	Status     int    `json:"status" xorm:"default 0 comment('0-正常，1-已删除') INT(11)"`
	Title      string `json:"title" xorm:"default '' VARCHAR(200)"`
	Type       int    `json:"type" xorm:"default 0 comment('1-会员,2-日常,1000+为三方的，如1000-会员充值') INT(11)"`
	UpdateTime int64  `json:"update_time" xorm:"default 0 BIGINT(20)"`
	UseStatus  int    `json:"use_status" xorm:"default 0 comment('0-进行中，1-已下线') INT(11)"`
	//
	RewardRuleArr []*RewardRule `json:"reward_rule_arr"`
}

type FetchActivityTypeCode int

const (
	FetchActivityTypeCode_All FetchActivityTypeCode = iota
	FetchActivityTypeCode_Id
	FetchActivityTypeCode_Type
)

type ActivityTypeCode int

const (
	ActivityTypeCode_Vip             ActivityTypeCode = 1
	ActivityTypeCode_Daily           ActivityTypeCode = 2
	ActivityTypeCode_ThirdVipRechage ActivityTypeCode = 1000
)
