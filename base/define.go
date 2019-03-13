package base

type SamhBaseRequest struct {
	Uid        int64          `form:"uid" json:"uid" binding:"required"`
	DeviceId   string         `form:"udid" json:"udid" binding:"required"`
	DeviceType DeviceTypeCode `form:"ud_type" json:"ud_type" binding:"-"`
}

type SamhResponse struct {
	Code SamhResponseCode `json:"status"`
	Msg  string           `json:"msg"`
	Data interface{}      `json:"data"`
}

//SamhResponseCode 状态返回码
type SamhResponseCode int

const (
	SamhResponseCode_Succ SamhResponseCode = 0

	//1000~1999 通用
	SamhResponseCode_ServerError   SamhResponseCode = 1000
	SamhResponseCode_Param_Less    SamhResponseCode = 1001
	SamhResponseCode_Param_Invalid SamhResponseCode = 1002
	SamhResponseCode_Data_Invalid  SamhResponseCode = 1003
	SamhResponseCode_Data_NotExist SamhResponseCode = 1004

	//2000~2999 运营
	SamhResponseCode_Activity_NotExist                              SamhResponseCode = 2000
	SamhResponseCode_Activity_NotStarted                            SamhResponseCode = 2001
	SamhResponseCode_Activity_Finish                                SamhResponseCode = 2002
	SamhResponseCode_NoDiamondBuyFinancialCardPurchaseQualification SamhResponseCode = 2003
	SamhResponseCode_NoDiamondGetBenefitQualification               SamhResponseCode = 2004
	SamhResponseCode_DiamondGetBenefited                            SamhResponseCode = 2005

	//3000~3999 会员
	SamhResponseCode_VipUpgradeOrdering          SamhResponseCode = 3000
	SamhResponseCode_LessDiamond                 SamhResponseCode = 3001
	SamhResponseCode_NoBindOneOfSinaQQWeixnPhone SamhResponseCode = 3002
)

var responseCodeToMsg = map[SamhResponseCode]string{
	SamhResponseCode_Succ: "请求成功",

	//1000~1999 通用
	SamhResponseCode_ServerError:   "服务器错误",
	SamhResponseCode_Param_Less:    "参数不足",
	SamhResponseCode_Param_Invalid: "参数无效",
	SamhResponseCode_Data_Invalid:  "数据无效",
	SamhResponseCode_Data_NotExist: "数据不存在",

	//2000~2999 运营
	SamhResponseCode_Activity_NotExist:                              "活动不存在",
	SamhResponseCode_Activity_NotStarted:                            "活动没开始",
	SamhResponseCode_Activity_Finish:                                "活动已结束",
	SamhResponseCode_NoDiamondBuyFinancialCardPurchaseQualification: "没购买资格，已购买过此套餐或已购买够2个套餐",
	SamhResponseCode_NoDiamondGetBenefitQualification:               "没有领取收益资格",
	SamhResponseCode_DiamondGetBenefited:                            "已领取过收益",

	//3000~3999 会员
	SamhResponseCode_VipUpgradeOrdering:          "已有升级订单进行中",
	SamhResponseCode_LessDiamond:                 "钻石不足",
	SamhResponseCode_NoBindOneOfSinaQQWeixnPhone: "没有绑定微博、qq、微信、手机号至少一个",
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
	DeviceTypeCode_Huawei DeviceTypeCode = 1
)
