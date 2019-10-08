package pay_modle

import (
	. "github.com/ufwfqpdgv/samh_common_lib/modle/common_modle"
)

var (
	url     string
	timeOut int
)

type InternalPrepayRequest struct {
	Uid       int64       `form:"uid" json:"uid"`
	DeviceId  string      `form:"udid" json:"udid"`
	AppId     int         `json:"app_id"` //飒漫画:6
	ProductId int         `json:"product_id"`
	PayType   PayTypeCode `json:"pay_type"`
	Supply    string      `json:"supply"`
	NotifyUrl string      `json:"notify_url"`
	ClientIp  string      `json:"client_ip"`
}

type InternalPrepayResponse interface{}

type InternalGetProductsRequest struct {
	Uid          int64            `form:"uid" json:"uid"`
	DeviceId     string           `form:"udid" json:"udid"`
	AppId        int              `form:"app_id" json:"app_id"` //飒漫画:6
	ProductType  string           `form:"product_type" json:"product_type"`
	Area         string           `form:"area" json:"area"`
	PayType      string           `form:"pay_type" json:"pay_type"`
	DiscountType DiscountTypeCode `form:"discount_type" json:"discount_type"`
}

type InternalGetProductsResponse struct {
	ProductArr []*Product `json:"product_list"`
}

type Product struct {
	ProductId int64  `json:"product_id"`
	Price     int    `json:"price"`
	Detail    string `json:"detail"`
	Extra     string `json:"extra"`
	Symbol    string `json:"symbol"`
}
