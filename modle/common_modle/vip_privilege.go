package common_modle

type VipPrivilege struct {
	AdditionalCollectionExpansion  int          `json:"additional_collection_expansion" xorm:"default 0 comment('收藏额外扩充，多少部漫画，-1表无限制') INT(11)"`
	AdvancedBarrageBackgroup       int          `json:"advanced_barrage_backgroup" xorm:"default 0 comment('高级弹幕背景，可用数量，-1表无限制') INT(11)"`
	CreateTime                     int64        `json:"create_time" xorm:"default 0 BIGINT(20)"`
	FreeDailyRecommendation        int          `json:"free_daily_recommendation" xorm:"default 0 comment('每日免费推荐票，倍数') INT(11)"`
	FreeGold                       int          `json:"free_gold" xorm:"default 0 comment('每月免费领取金币数') INT(11)"`
	HasVipExclusiveCustomerService int          `json:"has_vip_exclusive_customer_service" xorm:"default 0 comment('VIP专属客服，0-无，1-有') TINYINT(1)"`
	HeadPendant                    string       `json:"head_pendant" xorm:"default '头像挂件' VARCHAR(200)"`
	IdentityNameplate              string       `json:"identity_nameplate" xorm:"default '身份铭牌图标' VARCHAR(200)"`
	IsRemoveExternalAds            int          `json:"is_remove_external_ads" xorm:"default 0 comment('去除外接广告，0-无，1-全部') TINYINT(1)"`
	MonthlyComplementSignedChance  int          `json:"monthly_complement_signed_chance" xorm:"default 0 comment('每月补签机会，次数') INT(11)"`
	MonthlyFreeTicket              int          `json:"monthly_free_ticket" xorm:"default 0 comment('每月免费月票，倍数') INT(11)"`
	SignRewardDouble               int          `json:"sign_reward_double" xorm:"default 0 comment('签到奖励翻倍，0-否，1-是') INT(11)"`
	Status                         int          `json:"status" xorm:"default 0 comment('0-正常，1-已删除') INT(11)"`
	TaskExperienceSpeed            int          `json:"task_experience_speed" xorm:"default 0 comment('任务经验加速，倍数') INT(11)"`
	TaskGoldReward                 int          `json:"task_gold_reward" xorm:"default 0 comment('任务金币奖励，倍数') INT(11)"`
	UpdateTime                     int64        `json:"update_time" xorm:"default 0 BIGINT(20)"`
	UserNicknameColor              int          `json:"user_nickname_color" xorm:"default 0 comment('用户昵称颜色，1-红色') INT(11)"`
	VipComicFreeRead               int          `json:"vip_comic_free_read" xorm:"default 0 comment('付费漫画免费读,多少话/天，-1表无限制') INT(11)"`
	VipExclusiveHdPictureQuality   int          `json:"vip_exclusive_hd_picture_quality" xorm:"default 0 comment('VIP专属高清画质，0-无，1-有') TINYINT(1)"`
	VipExtraBookList               int          `json:"vip_extra_book_list" xorm:"default 0 comment('VIP额外书单，多少部漫画') INT(11)"`
	VipHeadPortraitMark            string       `json:"vip_head_portrait_mark" xorm:"default '会员头像标志' VARCHAR(200)"`
	VipLevel                       VipLevelCode `json:"vip_level" xorm:"default 0 comment('1-黄金会员，2-钻石会员') INT(11)"`
	VipPrivilegeId                 int64        `json:"vip_privilege_id" xorm:"not null pk autoincr BIGINT(20)"`
	//
	IsVip int   `json:"is_vip" xorm:"-"`
	Uid   int64 `json:"uid" xorm:"-"`
}
