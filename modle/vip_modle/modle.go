package vip_modle

import (
	. "github.com/ufwfqpdgv/samh_common_lib/modle/common_modle"

	"github.com/ufwfqpdgv/samh_common_lib/base"
)

var (
	url     string
	timeOut int
)

type InternalInitRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalInitResponse struct {
	Uid               int64 `json:"uid,omitempty"`
	GoldExpireTime    int64 `json:"gold_expire_time,omitempty"`
	DiamondExpireTime int64 `json:"diamond_expire_time,omitempty"`
	Diamonds          int   `json:"diamonds,omitempty"`
	Golds             int   `json:"golds,omitempty"`
}

type InternalUserPrivilegeInfoRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalUserPrivilegeInfoResponse struct {
	*VipPrivilege
}

type VipRewardRequest struct {
	base.SamhBaseRequest
	RewardRule   *RewardRule `form:"reward_rule" json:"reward_rule"`
	ActivityId   int64       `form:"activity_id" json:"activity_id"`
	RewardRuleId int64       `form:"reward_rule_id" json:"reward_rule_id"`
}

type VipRewardResponse struct {
}

type InternalRechargeRequest struct {
	base.SamhBaseRequest
	ActivityId    int64       `form:"activity_id" json:"activity_id"`
	RewardRuleId  int64       `form:"reward_rule_id" json:"reward_rule_id"`
	VipComboId    int64       `form:"vip_combo_id" json:"vip_combo_id"`
	ThirdId       int64       `form:"third_id" json:"third_id"`
	PayType       PayTypeCode `form:"pay_type" json:"pay_type"`
	IsPaypal      int         `form:"ispaypal" json:"ispaypal"`
	ClientType    string      `form:"client-type" json:"client-type"`
	ClientVersion string      `form:"client-version" json:"client-version"`
	Content       string      `form:"content" json:"content"`
}

type InternalRechargeResponse struct {
	Extra interface{} `json:"extra"`
}

type InternalOperationRequest struct {
	Uid               int64            `form:"uid" json:"uid" binding:"required"`
	VipOperation      VipOperationCode `form:"type" json:"type" binding:"required"`
	ActivityId        int64            `form:"activity_id" json:"activity_id"`
	RewardRuleId      int64            `form:"reward_rule_id" json:"reward_rule_id"`
	RewardRuleItem    *RewardRule      `form:"reward_rule" json:"reward_rule"`
	TaskId            int64            `form:"task_id" json:"task_id"`
	OperationRuleItem *OperationRule   `form:"operation_rule" json:"operation_rule" binding:"required"`
	UserLog           *UserLog2        `form:"user_log" json:"user_log"`
	CoinUser          *CoinUser2       `form:"coin_user" json:"coin_user"`
}

type InternalOperationResponse struct {
}

type VipOperationCode int

const (
	VipOperationCode_Recharge VipOperationCode = iota + 1
	VipOperationCode_Upgrade
	VipOperationCode_ActivityRecharge
	VipOperationCode_ActivityReward
	VipOperationCode_TaskReward
	VipOperationCode_ComicCost
)

type OperationRule struct {
	GoldTime    int64 `json:"gold_time"`
	DiamondTime int64 `json:"diamond_time"`
	Diamonds    int   `json:"diamonds"`
	Golds       int   `json:"golds"`
}

type UserLog2 struct {
	Ltarget  int    `form:"ltarget" json:"ltarget"`
	Lcontent string `form:"lcontent" json:"lcontent"`
	Laction  int    `form:"laction" json:"laction"`
	Ltid     int64  `form:"ltid" json:"ltid"`
	Lnumber  int    `form:"lnumber" json:"lnumber"`
	Lnumber2 int    `form:"lnumber2" json:"lnumber2"`
}

type CoinUser2 struct {
	Cgotcoin    int    `form:"Cgotcoin" json:"Cgotcoin"`
	Cremaincoin int    `form:"Cremaincoin" json:"Cremaincoin"`
	Cgotmode    int    `form:"Cgotmode" json:"Cgotmode"`
	Cgotdes     string `form:"Cgotdes" json:"Cgotdes"`
	Siteid      int    `form:"Siteid" json:"Siteid"`
	Platformid  int    `form:"Platformid" json:"Platformid"`
}

type CoinUser3 struct {
	Number   int    `form:"number" json:"number"` //number 金币数量
	ComicId  int64  `form:"comic_id" json:"comic_id"`
	TargetId int    `form:"target_id" json:"target_id"` //targetid 目标（漫画章节，其他操作）
	Content  string `form:"content" json:"content"`
	Lid      int    `form:"lid" json:"lid"` //lid 购买章节订单id
}

type InternalAllInfoRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalAllInfoResponse struct {
	Uid               int64        `json:"uid,omitempty"`
	VipLevel          VipLevelCode `json:"vip_level"`
	GoldExpireTime    int64        `json:"gold_expire_time,omitempty"`
	DiamondExpireTime int64        `json:"diamond_expire_time,omitempty"`
	Diamonds          int          `json:"diamonds,omitempty"`
	Golds             int          `json:"golds,omitempty"`
}

type InternalBatchAllInfoRequest struct {
	UidArr []int64 `form:"uid_arr" json:"uid_arr" binding:"required"`
}

type InternalBatchAllInfoResponse struct {
	VipInfoArr []*InternalAllInfoResponse `json:"vip_info_arr"`
}

type InternalIsRechargedEffectiveVipRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalIsRechargedEffectiveVipResponse struct {
	IsRechargedEffectiveVip bool `json:"is_recharged_effective_vip"`
}
