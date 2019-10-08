package user_modle

import (
	"samh_common_lib/base"
	"samh_common_lib/utils"
	"samh_common_lib/utils/log"
	"strconv"
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

func GetInfo(rq *GetInfoRequest) (rsp *GetInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &GetInfoResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data *GetInfoResponse      `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/userapi/internal/v1/getInfo"
	retCode = utils.HttpGet(addr, rq, rsp2, timeOut, nil, 3)
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

func GetBindInfo(rq *GetBindInfoRequest) (rsp *GetBindInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &GetBindInfoResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data *GetBindInfoResponse  `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/userapi/internal/v1/getBindInfo"
	retCode = utils.HttpGet(addr, rq, rsp2, timeOut, nil, 3)
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

func IsNewUser(rq *IsNewUserRequest) (rsp bool, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = false

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data bool                  `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/userapi/internal/v1/isNewUser"
	retCode = utils.HttpGet(addr, rq, rsp2, timeOut, nil, 3)
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

func GetBatchInfo(rq *GetBatchInfoRequest) (rsp *GetBatchInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &GetBatchInfoResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data []*GetInfoResponse    `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/userapi/internal/v1/getBatchInfo"
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
	rsp.UserInfoArr = rsp2.Data

	return
}

func DefaultGetInfoResponse() (rsp *GetInfoResponse) {
	rsp = &GetInfoResponse{
		LimitBook:     100,
		LimitCollect:  100,
		LimitFocus:    300,
		LimitBookSize: 100,
		Exp:           0,
		Fans:          0,
		Focus:         0,
	}
	return
}

func IsOldDevice(rq *IsOldDeviceRequest) (rsp bool, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = false

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data bool                  `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/userapi/internal/v1/isOldDevice"
	retCode = utils.HttpGet(addr, rq, rsp2, timeOut, nil, 3)
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

func IsGuest(rq *IsGuestRequest) (rsp bool, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = false

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data bool                  `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/userapi/internal/v1/isGuest"
	retCode = utils.HttpGet(addr, rq, rsp2, timeOut, nil, 3)
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

func IsGuest2(uid int64) (rsp bool) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	if len(strconv.FormatInt(uid, 10)) > 15 {
		rsp = true
	} else {
		rsp = false
	}

	return
}
