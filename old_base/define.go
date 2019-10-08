package old_base

type SamhBaseRequest struct {
	Uid      int64  `form:"uid" json:"uid" binding:"required"`
	DeviceId string `form:"udid" json:"udid" binding:"required"`
}

type SamhResponse struct {
	Code SamhResponseCode `json:"status"`
	Msg  string           `json:"msg"`
	Data interface{}      `json:"data,omitempty"`
}

//SamhResponseCode 状态返回码
type SamhResponseCode int

const (
	SamhResponseCode_Ok   SamhResponseCode = 1
	SamhResponseCode_Succ SamhResponseCode = 1

	SamhResponseCode_ServerError            SamhResponseCode = 0
	SamhResponseCode_Error                  SamhResponseCode = 0
	SamhResponseCode_ParameterError         SamhResponseCode = 2
	SamhResponseCode_Param_Invalid          SamhResponseCode = 2
	SamhResponseCode_Param_Less             SamhResponseCode = 2
	SamhResponseCode_SelfError              SamhResponseCode = 3
	SamhResponseCode_PermissionDenied       SamhResponseCode = 4
	SamhResponseCode_ContentTypeError       SamhResponseCode = 5
	SamhResponseCode_AlreadyExisted         SamhResponseCode = 6
	SamhResponseCode_NoExists               SamhResponseCode = 100
	SamhResponseCode_Exists                 SamhResponseCode = 101
	SamhResponseCode_ForbidChange           SamhResponseCode = 102
	SamhResponseCode_UserNoExists           SamhResponseCode = 103
	SamhResponseCode_UserExists             SamhResponseCode = 104
	SamhResponseCode_AlreadyGetPrice        SamhResponseCode = 105
	SamhResponseCode_ClientTimeDifference   SamhResponseCode = 106
	SamhResponseCode_ServiceStop            SamhResponseCode = 107
	SamhResponseCode_Relogin                SamhResponseCode = 201
	SamhResponseCode_SystemError            SamhResponseCode = 500
	SamhResponseCode_ContentCheck           SamhResponseCode = 2001
	SamhResponseCode_RePost                 SamhResponseCode = 2002
	SamhResponseCode_UserMust               SamhResponseCode = 2003
	SamhResponseCode_IpMust                 SamhResponseCode = 2004
	SamhResponseCode_NeedAuthCode           SamhResponseCode = 2005
	SamhResponseCode_AuthCodeError          SamhResponseCode = 2006
	SamhResponseCode_LevelLimit             SamhResponseCode = 2007
	SamhResponseCode_NeedPhone              SamhResponseCode = 2008
	SamhResponseCode_ContainsSensitiveWords SamhResponseCode = 3000
)

var responseCodeToMsg = map[SamhResponseCode]string{
	SamhResponseCode_Ok: "执行成功",

	SamhResponseCode_Error:                  "执行错误",
	SamhResponseCode_ParameterError:         "参数错误",
	SamhResponseCode_SelfError:              "自定义错误",
	SamhResponseCode_PermissionDenied:       "权限验证不通过",
	SamhResponseCode_ContentTypeError:       "请求类型出错",
	SamhResponseCode_AlreadyExisted:         "验证过期",
	SamhResponseCode_NoExists:               "操作对象不存在",
	SamhResponseCode_Exists:                 "操作对象存在",
	SamhResponseCode_ForbidChange:           "禁止更新",
	SamhResponseCode_UserNoExists:           "用户不存在",
	SamhResponseCode_UserExists:             "用户存在",
	SamhResponseCode_AlreadyGetPrice:        "任务已经领取",
	SamhResponseCode_ClientTimeDifference:   "服务端与客户端时间不一致",
	SamhResponseCode_ServiceStop:            "服务停止",
	SamhResponseCode_Relogin:                "重新登录",
	SamhResponseCode_SystemError:            "系统错误",
	SamhResponseCode_ContentCheck:           "内容审核中",
	SamhResponseCode_RePost:                 "重复提交",
	SamhResponseCode_UserMust:               "用户被封禁",
	SamhResponseCode_IpMust:                 "IP被土封禁",
	SamhResponseCode_NeedAuthCode:           "需要验证码",
	SamhResponseCode_AuthCodeError:          "验证码错误",
	SamhResponseCode_LevelLimit:             "等级限制",
	SamhResponseCode_NeedPhone:              "需要绑定手机",
	SamhResponseCode_ContainsSensitiveWords: "包含敏感词",
}

type SamhDataStatusCode int

const (
	SamhDataStatusCode_Deleted SamhDataStatusCode = 0
	SamhDataStatusCode_Normal  SamhDataStatusCode = 1
)
