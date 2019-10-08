package user_modle

import "time"

var (
	url     string
	timeOut int
)

type UserLog struct {
	Lid      int64     `json:"lid" xorm:"not null pk autoincr BIGINT(20)"`
	Luid     int64     `json:"luid" xorm:"not null default 0 BIGINT(20)"`
	Ltarget  int       `json:"ltarget" xorm:"not null default 0 INT(11)"`
	Lcontent string    `json:"lcontent" xorm:"not null default '0' VARCHAR(50)"`
	Ltime    time.Time `json:"ltime" xorm:"not null TIMESTAMP"`
	Laction  int       `json:"laction" xorm:"not null default 0 TINYINT(4)"`
	Lstatus  int       `json:"lstatus" xorm:"not null default 0 TINYINT(4)"`
	Ltid     int64     `json:"ltid" xorm:"not null default 0 BIGINT(20)"`
	Lnumber  int       `json:"lnumber" xorm:"not null default 0 INT(11)"`
	Lnumber2 int       `json:"lnumber2" xorm:"not null default 0 INT(11)"`
}

type GetInfoRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type GetInfoResponse struct {
	Uid               int64       `json:"Uid"`
	NickName          string      `json:"Uname"`
	Sex               int         `json:"Usex"`
	RegIp             string      `json:"Uregip"`
	RegTime           int64       `json:"Uregtime"`
	UTpye             int         `json:"Utpye"`
	LoginType         string      `json:"type"`
	Level             int         `json:"Ulevel"`
	Mail              string      `json:"Umail"`
	Sign              string      `json:"Usign"`
	OpenId            string      `json:"openid"`
	GoldVipTime       int64       `json:"goldVipTime"`
	GoldVipOverDay    int64       `json:"goldVipoverday"`
	DiamondVipTime    int64       `json:"diamondViptime"`
	DiamondVipOverDay int64       `json:"diamondVipoverday"`
	AreaCode          string      `json:"Uareacode"`
	TagId             string      `json:"Utag_ids"`
	Birthday          int64       `json:"Ubirthday"`
	IsVip             int         `json:"isvip"`
	Fans              int         `json:"Cfans"`
	Focus             int         `json:"Cfocus"`
	Ticket            int         `json:"Cticket"`
	Exp               int         `json:"Cexp"`
	Diamonds          int         `json:"diamonds"`
	Coins             int         `json:"coins"`
	BookNum           int         `json:"book_num"`
	Ismkxq            int         `json:"ismkxq"`
	IsNewUser         int         `json:"isnewuser"`
	LimitBook         int         `json:"limitbook"`
	LimitBookSize     int         `json:"limitbooksize"`
	LimitCollect      int         `json:"limitcollect"`
	LimitFocus        int         `json:"limitfocus"`
	NextLevelExp      float64     `json:"nextlevelexp"`
	Recommend         int         `json:"recommend"`
	RoleInfo          interface{} `json:"roleinfo"`
	AuthData          Auth        `json:"auth_data"`
	CommerceAuth      Auth        `json:"commerceauth"`
	CommunityData     Auth        `json:"community_data"`
	TaskData          Auth        `json:"task_data"`
	IsGuest           bool        `json:"is_guest"`
}

type Auth struct {
	Appid       int    `json:"appid"`
	Expiry      int64  `json:"expiry"`
	AuthCode    string `json:"authcode"`
	ImageDomain string `json:"imagedomain"`
	ImageLimit  string `json:"imagelimit"`
}

type GetBindInfoRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type GetBindInfoResponse struct {
	QQ     string `json:"qq"`
	Weixin string `json:"weixin"`
	Mkxq   string `json:"mkxq"`
	Sina   string `json:"sina"`
	Mobile string `json:"mobile"`
}

type IsNewUserRequest struct {
	Uid      int64  `form:"uid" json:"uid" binding:"required"`
	DeviceId string `form:"udid" json:"udid" binding:"required"`
}

type GetBatchInfoRequest struct {
	UidArr []int64 `form:"uids" json:"uids" binding:"required"`
}

type GetBatchInfoResponse struct {
	UserInfoArr []*GetInfoResponse
}

type IsOldDeviceRequest struct {
	DeviceId string `form:"udid" json:"udid" binding:"required"`
}

type IsGuestRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}
