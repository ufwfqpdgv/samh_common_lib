package common_modle

type Cartoon struct {
	CartoonId                 int    `json:"cartoon_id" xorm:"not null pk INT(11)"`
	CartoonName               string `json:"cartoon_name" xorm:"VARCHAR(50)"`
	PinyinName                string `json:"PinYin_name" xorm:"VARCHAR(100)"`
	CartoonDesc               string `json:"cartoon_desc" xorm:"VARCHAR(2500)"`
	CartoonStatusId           int    `json:"cartoon_status_id" xorm:"INT(11)"`
	CartoonAuthorListName     string `json:"cartoon_author_list_name" xorm:"VARCHAR(250)"`
	CartoonTypeListName       string `json:"cartoon_type_list_name" xorm:"VARCHAR(500)"`
	CartoonAreaId             int    `json:"cartoon_area_id" xorm:"INT(11)"`
	TodayViewNum              int64  `json:"today_view_num" xorm:"BIGINT(20)"`
	YesterdayViewNum          int64  `json:"yesterday_view_num" xorm:"BIGINT(20)"`
	DayBeforeYesterdayViewNum int64  `json:"day_before_yesterday_view_num" xorm:"BIGINT(20)"`
	TotalViewNum              int64  `json:"total_view_num" xorm:"BIGINT(20)"`
	UpdateTime                string `json:"update_time" xorm:"VARCHAR(20)"`
	CurrentMonthTotalViewNum  int64  `json:"current_month_total_view_num" xorm:"BIGINT(20)"`
	LastMonthTotalViewNum     int64  `json:"last_month_total_view_num" xorm:"BIGINT(20)"`
	IsRecommend               int    `json:"is_recommend" xorm:"INT(11)"`
	CartoonImgAddr            string `json:"cartoon_img_addr" xorm:"VARCHAR(300)"`
	CartoonColorId            int    `json:"cartoon_color_id" xorm:"INT(11)"`
	CartoonLetterId           int    `json:"cartoon_letter_id" xorm:"INT(11)"`
	LatestCartoonTopicId      string `json:"latest_cartoon_topic_id" xorm:"VARCHAR(50)"`
	LatestCartoonTopicNewid   string `json:"latest_cartoon_topic_newid" xorm:"not null VARCHAR(50)"`
	LatestCartoonTopicName    string `json:"latest_cartoon_topic_name" xorm:"VARCHAR(100)"`
	Bbsid                     string `json:"bbsid" xorm:"VARCHAR(80)"`
	CartoonTag                string `json:"cartoon_tag" xorm:"not null VARCHAR(150)"`
	LocalhostCreateTime       string `json:"localhost_create_time" xorm:"VARCHAR(20)"`
	CreateTime                string `json:"create_time" xorm:"VARCHAR(20)"`
	ComicId                   string `json:"comic_id" xorm:"VARCHAR(15)"`
	CopyrightDesc             string `json:"copyright_desc" xorm:"not null VARCHAR(500)"`
	CartoonNoticeInfo         string `json:"cartoon_notice_info" xorm:"VARCHAR(500)"`
	SeoTitle                  string `json:"seo_title" xorm:"VARCHAR(100)"`
	SeoKeywords               string `json:"seo_keywords" xorm:"VARCHAR(100)"`
	SeoDesc                   string `json:"seo_desc" xorm:"VARCHAR(500)"`
	Xiaoshuodata              string `json:"xiaoshuodata" xorm:"VARCHAR(250)"`
	OperatingTime             string `json:"operating_time" xorm:"VARCHAR(20)"`
	Magazine                  string `json:"magazine" xorm:"VARCHAR(50)"`
	Downnum                   int    `json:"downnum" xorm:"INT(11)"`
	EditorPerson              string `json:"editor_person" xorm:"not null VARCHAR(100)"`
	CartoonNewid              string `json:"cartoon_newid" xorm:"not null VARCHAR(50)"`
	CartoonShortdesc          string `json:"cartoon_shortdesc" xorm:"VARCHAR(50)"`
	JideId                    string `json:"jide_id" xorm:"VARCHAR(50)"`
	LeshenId                  string `json:"leshen_id" xorm:"VARCHAR(50)"`
	KxdmId                    string `json:"kxdm_id" xorm:"VARCHAR(50)"`
	Mh160Id                   string `json:"mh160_id" xorm:"VARCHAR(50)"`
	ShowType                  int    `json:"show_type" xorm:"not null TINYINT(4)"`
	Isoutlink                 int    `json:"isoutlink" xorm:"not null INT(11)"`
	Pingfen                   string `json:"pingfen" xorm:"not null DECIMAL(10,8)"`
	PingfenAll                int    `json:"pingfen_all" xorm:"not null INT(11)"`
	Pingfen1                  int    `json:"pingfen_1" xorm:"not null INT(11)"`
	Pingfen2                  int    `json:"pingfen_2" xorm:"not null INT(11)"`
	Pingfen3                  int    `json:"pingfen_3" xorm:"not null INT(11)"`
	Pingfen4                  int    `json:"pingfen_4" xorm:"not null INT(11)"`
	Pingfen5                  int    `json:"pingfen_5" xorm:"not null INT(11)"`
	Shoucang                  int    `json:"shoucang" xorm:"not null INT(11)"`
	Tuijian                   int    `json:"tuijian" xorm:"not null INT(11)"`
	Xihuan                    int    `json:"xihuan" xorm:"not null INT(11)"`
	Yanwu                     int    `json:"yanwu" xorm:"not null INT(11)"`
	Share                     int    `json:"share" xorm:"not null INT(11)"`
	Dashang                   int    `json:"dashang" xorm:"not null INT(11)"`
	Yuepiao                   int    `json:"yuepiao" xorm:"not null INT(11)"`
	Zongpiao                  int    `json:"zongpiao" xorm:"not null INT(11)"`
	SiteStatus                string `json:"site_status" xorm:"not null CHAR(50)"`
	ChargeStatus              string `json:"charge_status" xorm:"not null CHAR(20)"`
	ComicFeature              string `json:"comic_feature" xorm:"not null VARCHAR(50)"`
	UpdateRule                string `json:"update_rule" xorm:"not null VARCHAR(50)"`
	DefaultPrice              int    `json:"default_price" xorm:"not null INT(11)"`
	Readtype                  int    `json:"readtype" xorm:"not null TINYINT(4)"`
	CountRate                 string `json:"count_rate" xorm:"not null DECIMAL(5,3)"`
	SettledRate               string `json:"settled_rate" xorm:"not null DECIMAL(5,5)"`
	TradefeeRate              string `json:"tradefee_rate" xorm:"not null DECIMAL(5,3)"`
	SharecropRate             string `json:"sharecrop_rate" xorm:"not null DECIMAL(5,3)"`
	CopyrightType             int    `json:"copyright_type" xorm:"not null TINYINT(4)"`
	HumanType                 int    `json:"human_type" xorm:"not null TINYINT(4)"`
	CartoonSeries             string `json:"cartoon_series" xorm:"not null VARCHAR(50)"`
	ComicAlias                string `json:"comic_alias" xorm:"not null VARCHAR(50)"`
	UploaderUid               int    `json:"uploader_Uid" xorm:"not null INT(11)"`
	UploaderUname             string `json:"uploader_Uname" xorm:"not null VARCHAR(50)"`
	ShieldingChannel          string `json:"shielding_channel" xorm:"VARCHAR(200)" DEFAULT ''`
}
