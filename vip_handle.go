package samh_common_lib

import (
	"github.com/ufwfqpdgv/log"
	"github.com/ufwfqpdgv/utils"
)

func Init(urlRq string, timeOutRq int) {
	var err error
	url, timeOut, err = utils.CheckUrlTimeout(urlRq, timeOutRq)
	if err != nil {
		log.Panic(err)
	}
	err = utils.CheckServerConnect(urlRq)
	if err != nil {
		// Panic(err) // 不能用Panic，如整个服务器重启，那互相依赖的服务谁也启动不了
		log.Error(err)
	}
}

func InitUser(rq *InternalInitRequest) (rsp *InternalInitResponse, retCode SamhResponseCode) {
	log.Debug(NowFunc())
	defer log.Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalInitResponse{}

	rsp2 := &struct {
		Code SamhResponseCode      `json:"status"`
		Msg  string                `json:"msg"`
		Data *InternalInitResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/init"
	err := utils.HttpGet(addr, rq, rsp2, timeOut, nil, 3)
	if err != nil {
		retCode = SamhResponseCode_ServerError
		log.Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		log.Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}
