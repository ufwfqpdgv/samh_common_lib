package pay_modle

import (
	"github.com/ufwfqpdgv/samh_common_lib/utils"
	"github.com/ufwfqpdgv/samh_common_lib/utils/log"

	"github.com/ufwfqpdgv/samh_common_lib/base"
)

func Init(urlRq string, timeOutRq int) {
	var err error
	url, timeOut, err = utils.CheckUrlTimeout(urlRq, timeOutRq)
	if err != nil {
		log.Panic(err)
	}
	err = utils.CheckServerConnect(urlRq)
	if err != nil {
		// log.Panic(err) // 不能用Panic，如整个服务器重启，那互相依赖的服务谁也启动不了
		log.Error(err)
	}
}

func InternalPrepay(rq *InternalPrepayRequest) (rsp *InternalPrepayResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	// rsp = &InternalPrepayResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode   `json:"status"`
		Msg  string                  `json:"msg"`
		Data *InternalPrepayResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/app_api/pay/prepay"
	retCode = utils.HttpPost(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != base.SamhResponseCode_Succ {
		log.Error(base.NowFuncError())
		return
	}
	if rsp2.Code != base.SamhResponseCode_Succ {
		log.Error(base.NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func InternalGetProducts(rq *InternalGetProductsRequest) (rsp *InternalGetProductsResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalGetProductsResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode        `json:"status"`
		Msg  string                       `json:"msg"`
		Data *InternalGetProductsResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/app_api/pay/get_products"
	retCode = utils.HttpPost(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != base.SamhResponseCode_Succ {
		log.Error(base.NowFuncError())
		return
	}
	if rsp2.Code != base.SamhResponseCode_Succ {
		log.Error(base.NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}
