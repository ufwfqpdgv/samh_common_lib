package base

type SamhBaseRequest struct {
	Uid        int64          `form:"uid" json:"uid" binding:"required"`
	DeviceId   string         `form:"udid" json:"udid" binding:"required"`
	DeviceType DeviceTypeCode `form:"ud_type" json:"ud_type" binding:"-"`
}

type SamhResponse struct {
	Code SamhResponseCode `json:"status"`
	Msg  string           `json:"msg"`
	Data interface{}      `json:"data,omitempty"`
	// 上msg 是调试文本，下为给 app 显示的提示
	// 客户端在在动作的地方如果需要显示的话，就会显示返回里的 msg，如有特殊的提示框啥的话，就直接特殊处理以 code 判断，tip相关暂不加了
	// Tip_type    SamhTipCode      `json:"tip_type"`
	// Tip_content interface{}      `json:"tip_content,omitempty"`
}

type SamhResponseCustom struct {
	Code interface{} `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//SamhResponseCode 状态返回码
type SamhResponseCode int

const (
	SamhResponseCode_Succ SamhResponseCode = 0

	//1000~1999 通用
	SamhResponseCode_ServerError               SamhResponseCode = 1000
	SamhResponseCode_Param_Less                SamhResponseCode = 1001
	SamhResponseCode_Param_Invalid             SamhResponseCode = 1002
	SamhResponseCode_Data_Invalid              SamhResponseCode = 1003
	SamhResponseCode_Data_NotExist             SamhResponseCode = 1004
	SamhResponseCode_User_NotExist             SamhResponseCode = 1005
	SamhResponseCode_Upload_Save_Image_Fail    SamhResponseCode = 1006
	SamhResponseCode_Upload_Check_Image_Fail   SamhResponseCode = 1007
	SamhResponseCode_Upload_Check_Image_Format SamhResponseCode = 1008
	SamhResponseCode_AccountNotExist           SamhResponseCode = 1024
	SamhResponseCode_Repeat_Request_Invalid    SamhResponseCode = 1009

	//2000~2999 运营
	SamhResponseCode_Activity_NotExist                              SamhResponseCode = 2000
	SamhResponseCode_Activity_NotStarted                            SamhResponseCode = 2001
	SamhResponseCode_Activity_Finish                                SamhResponseCode = 2002
	SamhResponseCode_NoDiamondBuyFinancialCardPurchaseQualification SamhResponseCode = 2003
	SamhResponseCode_NoDiamondGetBenefitQualification               SamhResponseCode = 2004
	SamhResponseCode_DiamondGetBenefited                            SamhResponseCode = 2005
	SamhResponseCode_CreateGroupEnough                              SamhResponseCode = 2006
	SamhResponseCode_CreateGroupMemberEnough                        SamhResponseCode = 2007
	SamhResponseCode_GroupExist                                     SamhResponseCode = 2008
	SamhResponseCode_MemberExist                                    SamhResponseCode = 2009
	SamhResponseCode_SerialError                                    SamhResponseCode = 2010
	SamhResponseCode_JoinGroupMemberEnough                          SamhResponseCode = 2011
	SamhResponseCode_JoinTooFast                                    SamhResponseCode = 2012
	SamhResponseCode_AnswerExist                                    SamhResponseCode = 2013
	SamhResponseCode_SpecialCharacter                               SamhResponseCode = 2014
	SamhResponseCode_UidAlreadyJoinNewUserPackage                   SamhResponseCode = 2015
	SamhResponseCode_IsOldDevice                                    SamhResponseCode = 2016
	SamhResponseCode_HaveOverNewUserPackageVipTime                  SamhResponseCode = 2017

	//3000~3999 会员
	SamhResponseCode_VipUpgradeOrdering          SamhResponseCode = 3000
	SamhResponseCode_LessDiamond                 SamhResponseCode = 3001
	SamhResponseCode_NoBindOneOfSinaQQWeixnPhone SamhResponseCode = 3002
	SamhResponseCode_LessGolds                   SamhResponseCode = 3003
	SamhResponseCode_Processed_Order             SamhResponseCode = 3004
	SamhResponseCode_NonProbationaryVip          SamhResponseCode = 3005

	//4000~4999 漫画
	SamhResponseCode_LessMonthlyTicket                             SamhResponseCode = 4000
	SamhResponseCode_LessRecommendedTicket                         SamhResponseCode = 4001
	SamhResponseCode_NoBookId                                      SamhResponseCode = 4002
	SamhResponseCode_CanCollectSelfBook                            SamhResponseCode = 4003
	SamhResponseCode_BookNotExistOrDisable                         SamhResponseCode = 4004
	SamhResponseCode_CanCancelCollectSelfBook                      SamhResponseCode = 4005
	SamhResponseCode_NoPermissionsBookNotExistOrDisable            SamhResponseCode = 4006
	SamhResponseCode_ComicNotExistOrDisable                        SamhResponseCode = 4007
	SamhResponseCode_BookComicLimit                                SamhResponseCode = 4008
	SamhResponseCode_BuyedRelatedChapters                          SamhResponseCode = 4009
	SamhResponseCode_DonotNeedBuyFreeChapter                       SamhResponseCode = 4010
	SamhResponseCode_TouristsCanAppreciate                         SamhResponseCode = 4011
	SamhResponseCode_ComicCannotFreeRead                           SamhResponseCode = 4012
	SamhResponseCode_LessMoney                                     SamhResponseCode = 4013
	SamhResponseCode_GoldReadTimesRemain                           SamhResponseCode = 4014
	SamhResponseCode_GoldReadTimesDone                             SamhResponseCode = 4015
	SamhResponseCode_BookmarkAlreadyExists                         SamhResponseCode = 4016
	SamhResponseCode_AdReadTimesRemain                             SamhResponseCode = 4017
	SamhResponseCode_ReadingTicketNumberNotEnough                  SamhResponseCode = 4018
	SamhResponseCode_ReachedReadingPageSeeAdDailyRewardNumberLimit SamhResponseCode = 4019
)

var responseCodeToMsg = map[SamhResponseCode]string{
	SamhResponseCode_Succ: "请求成功",

	//1000~1999 通用
	SamhResponseCode_ServerError:               "服务器错误",
	SamhResponseCode_Param_Less:                "参数不足",
	SamhResponseCode_Param_Invalid:             "参数无效",
	SamhResponseCode_Data_Invalid:              "数据无效",
	SamhResponseCode_Data_NotExist:             "数据不存在",
	SamhResponseCode_User_NotExist:             "用户不存在",
	SamhResponseCode_Upload_Save_Image_Fail:    "保存图片失败",
	SamhResponseCode_Upload_Check_Image_Fail:   "检查图片失败",
	SamhResponseCode_Upload_Check_Image_Format: "图片格式或大小有问题",
	SamhResponseCode_AccountNotExist:           "账号不存在",
	SamhResponseCode_Repeat_Request_Invalid:    "重复请求，无效",

	//2000~2999 运营
	SamhResponseCode_Activity_NotExist:                              "活动不存在",
	SamhResponseCode_Activity_NotStarted:                            "活动没开始",
	SamhResponseCode_Activity_Finish:                                "活动已结束",
	SamhResponseCode_NoDiamondBuyFinancialCardPurchaseQualification: "没购买资格，已购买过此套餐或已购买够2个套餐",
	SamhResponseCode_NoDiamondGetBenefitQualification:               "没有领取收益资格",
	SamhResponseCode_DiamondGetBenefited:                            "已领取过收益",
	SamhResponseCode_CreateGroupEnough:                              "已经创建了3个队伍",
	SamhResponseCode_CreateGroupMemberEnough:                        "队伍已经有4名成员",
	SamhResponseCode_GroupExist:                                     "团队4名成员不能和其他队伍完全相同哦！",
	SamhResponseCode_MemberExist:                                    "不可重复加入同一个团队哦！",
	SamhResponseCode_SerialError:                                    "请填写正确的邀请码！",
	SamhResponseCode_JoinGroupMemberEnough:                          "每个用户最多加入3个团队！",
	SamhResponseCode_JoinTooFast:                                    "操作速度太快！",
	SamhResponseCode_AnswerExist:                                    "已答题！",
	SamhResponseCode_SpecialCharacter:                               "含有特殊字符！",
	SamhResponseCode_UidAlreadyJoinNewUserPackage:                   "uid已经参与过新用户大礼包活动",
	SamhResponseCode_IsOldDevice:                                    "老设备",
	SamhResponseCode_HaveOverNewUserPackageVipTime:                  "拥有超过新用户大礼包的会员时间",

	//3000~3999 会员
	SamhResponseCode_VipUpgradeOrdering:          "已有升级订单进行中",
	SamhResponseCode_LessDiamond:                 "宝石不足",
	SamhResponseCode_NoBindOneOfSinaQQWeixnPhone: "没有绑定微博、qq、微信、手机号至少一个",
	SamhResponseCode_LessGolds:                   "金币不足",
	SamhResponseCode_Processed_Order:             "已处理过的订单",
	SamhResponseCode_NonProbationaryVip:          "非试用会员",

	//4000~4999 漫画
	SamhResponseCode_LessMonthlyTicket:                             "月票不足",
	SamhResponseCode_LessRecommendedTicket:                         "推荐票不足",
	SamhResponseCode_NoBookId:                                      "书单 id 不能为空",
	SamhResponseCode_CanCollectSelfBook:                            "不能收藏自己创建的书单",
	SamhResponseCode_BookNotExistOrDisable:                         "书单不存在或已被禁用！",
	SamhResponseCode_CanCancelCollectSelfBook:                      "不能操作取消收藏自己创建的书单！",
	SamhResponseCode_NoPermissionsBookNotExistOrDisable:            "没有操作此书单的权限或书单被禁用",
	SamhResponseCode_ComicNotExistOrDisable:                        "漫画不存在或已被禁用！",
	SamhResponseCode_BookComicLimit:                                "最多100本漫画哦~",
	SamhResponseCode_BuyedRelatedChapters:                          "您已购买相关章节",
	SamhResponseCode_DonotNeedBuyFreeChapter:                       "无需购买免费章节",
	SamhResponseCode_TouristsCanAppreciate:                         "游客账号暂不能打赏！",
	SamhResponseCode_ComicCannotFreeRead:                           "此漫画暂时无法通过免费读观看！",
	SamhResponseCode_LessMoney:                                     "余额不足，请前去充值中心充值",
	SamhResponseCode_GoldReadTimesRemain:                           "您今日的黄金会员免费阅读次数还剩%v章",
	SamhResponseCode_GoldReadTimesDone:                             "您今日的黄金会员免费阅读次数已用完",
	SamhResponseCode_BookmarkAlreadyExists:                         "这个书签已存在",
	SamhResponseCode_AdReadTimesRemain:                             "您今日的看广告免费阅读次数还剩%v章",
	SamhResponseCode_ReadingTicketNumberNotEnough:                  "阅读券数量不足",
	SamhResponseCode_ReachedReadingPageSeeAdDailyRewardNumberLimit: "已达阅读页看广告每日奖励限制",
}

type SamhDataStatusCode int

const (
	SamhDataStatusCode_Normal  SamhDataStatusCode = 0
	SamhDataStatusCode_Deleted SamhDataStatusCode = 1

	SamhDataStatusCode_Online  SamhDataStatusCode = 0
	SamhDataStatusCode_Offline SamhDataStatusCode = 1

	SamhDataStatusCode_NotReward SamhDataStatusCode = 0
	SamhDataStatusCode_Rewarded  SamhDataStatusCode = 1
)

type DeviceTypeCode int

const (
	DeviceTypeCode_Iphone  DeviceTypeCode = 1
	DeviceTypeCode_Android DeviceTypeCode = 2
	DeviceTypeCode_Huawei  DeviceTypeCode = 3
)

type ClientTypeCode int

const (
	ClientTypeCode_Ios ClientTypeCode = iota + 1
	ClientTypeCode_Android
	ClientTypeCode_Huawei
)
