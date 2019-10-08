package samh_common_lib

type RewardRule struct {
	Coin                int                   `json:"coin" xorm:"default 0 INT(11)"`
	Content             string                `json:"content" xorm:"default '' VARCHAR(1000)"`
	CreateTime          int64                 `json:"create_time" xorm:"default 0 BIGINT(20)"`
	Exp                 int                   `json:"exp" xorm:"default 0 INT(11)"`
	Rechange            int                   `json:"rechange" xorm:"default 0 comment('单位是分，充值多少') INT(11)"`
	RewardRuleId        int64                 `json:"reward_rule_id" xorm:"not null pk autoincr BIGINT(20)"`
	Status              int                   `json:"status" xorm:"default 0 comment('0-正常，1-已删除') INT(11)"`
	Title               string                `json:"title" xorm:"default '' VARCHAR(200)"`
	Type                int                   `json:"type" xorm:"default 0 comment('1-不需要充值什么的，直接请求奖励如分享得奖、新用户直接得奖,2-需要充值之类的，直接请求VIP服务那边并把返回的data放到extra_data里(一类活动就得对应一个type)') INT(11)"`
	UpdateTime          int64                 `json:"update_time" xorm:"default 0 BIGINT(20)"`
	VipTime             int64                 `json:"vip_time" xorm:"default 0 comment('单位是s') BIGINT(20)"`
	VipType             VipLevelCode          `json:"vip_type" xorm:"default 0 comment('1-黄金，2-钻石') INT(11)"`
	AndroidThirdId      int64                 `json:"android_third_id" xorm:"default 0 comment('来源三方的id，如虚拟商品、good_id') BIGINT(20)"`
	IosThirdId          int64                 `json:"ios_third_id" xorm:"default 0 comment('来源三方的id，如虚拟商品、good_id') BIGINT(20)"`
	ReadingTicketNumber int                   `json:"reading_ticket_number" xorm:"not null default 0 comment('阅读券数量') TINYINT(4)"`
	ReadingTicketType   ReadingTicketTypeCode `json:"reading_ticket_type" xorm:"not null default 0 comment('阅读券类型。0-无，即无此奖励；1-通用券；2-某个漫画的专属券') TINYINT(4)"`
}

type ReadingTicketTypeCode int

const (
	ReadingTicketTypeCode_None ReadingTicketTypeCode = iota
	ReadingTicketTypeCode_Common
	ReadingTicketTypeCode_Exclusive
)
