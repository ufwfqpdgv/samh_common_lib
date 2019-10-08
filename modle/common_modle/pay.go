package common_modle

type PayTypeCode int

const (
	PayTypeCode_AndroidAlipay PayTypeCode = 3
	PayTypeCode_AppleStore    PayTypeCode = 4
	PayTypeCode_AndroidWechat PayTypeCode = 9
	PayTypeCode_AndroidQQ     PayTypeCode = 16
)

type DiscountTypeCode int

const (
	DiscountTypeCode_None DiscountTypeCode = iota
	DiscountTypeCode_Vip
)
