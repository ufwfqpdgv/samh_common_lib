package common_modle

import "samh_common_lib/base"

type UserVip struct {
	UserVipId         int64                   `json:"user_vip_id" xorm:"not null pk autoincr BIGINT(20)"`
	Uid               int64                   `json:"uid" xorm:"not null default 0 comment('用户id') unique BIGINT(20)"`
	GoldExpireTime    int64                   `json:"gold_expire_time" xorm:"default 0 comment('黄金会员到期时间，当同一用户拥有多种会员时，如加钻石会员时间时得将黄金会员的时候后延至钻石会员结束的时间点') BIGINT(20)"`
	DiamondExpireTime int64                   `json:"diamond_expire_time" xorm:"default 0 comment('钻石会员到期时间，当同一用户拥有多种会员时，如加钻石会员时间时得将黄金会员的时候后延至钻石会员结束的时间点') BIGINT(20)"`
	Diamonds          int                     `json:"diamonds" xorm:"default 0 comment('宝石') INT(11)"`
	CreateTime        int64                   `json:"create_time" xorm:"default 0 BIGINT(20)"`
	UpdateTime        int64                   `json:"update_time" xorm:"default 0 BIGINT(20)"`
	Status            base.SamhDataStatusCode `json:"status" xorm:"default 0 comment('0-正常，1-已删除') INT(11)"`
}
