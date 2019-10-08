package vip_modle

import (
	"samh_common_lib/utils"
	"samh_common_lib/utils/log"

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

func InitUser(rq *InternalInitRequest) (rsp *InternalInitResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalInitResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data *InternalInitResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/init"
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

func FetchUserPrivilegeInfo(rq *InternalUserPrivilegeInfoRequest) (rsp *InternalUserPrivilegeInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalUserPrivilegeInfoResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode              `json:"status"`
		Msg  string                             `json:"msg"`
		Data *InternalUserPrivilegeInfoResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/user_privilege_info"
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

func VipReward(rq *VipRewardRequest) (rsp *VipRewardResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &VipRewardResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode `json:"status"`
		Msg  string                `json:"msg"`
		Data *VipRewardResponse    `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/reward"
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

func InternalRechargeApi(rq *InternalRechargeRequest) (rsp *InternalRechargeResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalRechargeResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode     `json:"status"`
		Msg  string                    `json:"msg"`
		Data *InternalRechargeResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/recharge"
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

func InternalOperationApi(rq *InternalOperationRequest) (rsp *InternalOperationResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalOperationResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode      `json:"status"`
		Msg  string                     `json:"msg"`
		Data *InternalOperationResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/operation"
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

func InternalAllInfoApi(rq *InternalAllInfoRequest) (rsp *InternalAllInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalAllInfoResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode    `json:"status"`
		Msg  string                   `json:"msg"`
		Data *InternalAllInfoResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/all_info"
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

func DefaultInternalAllInfoResponse() (rsp *InternalAllInfoResponse) {
	rsp = &InternalAllInfoResponse{
		VipLevel:          0,
		GoldExpireTime:    0,
		DiamondExpireTime: 0,
		Diamonds:          0,
		Golds:             0,
	}
	return
}

func InternalBatchAllInfoApi(rq *InternalBatchAllInfoRequest) (rsp *InternalBatchAllInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalBatchAllInfoResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode         `json:"status"`
		Msg  string                        `json:"msg"`
		Data *InternalBatchAllInfoResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/batch_all_info"
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

func InternalIsRechargedEffectiveVipApi(rq *InternalIsRechargedEffectiveVipRequest) (rsp *InternalIsRechargedEffectiveVipResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalIsRechargedEffectiveVipResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode                    `json:"status"`
		Msg  string                                   `json:"msg"`
		Data *InternalIsRechargedEffectiveVipResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/is_recharged_effective_vip"
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

func InternalUserPrivilegeInfoApi(rq *InternalUserPrivilegeInfoRequest) (rsp *InternalUserPrivilegeInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalUserPrivilegeInfoResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode              `json:"status"`
		Msg  string                             `json:"msg"`
		Data *InternalUserPrivilegeInfoResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/user_privilege_info"
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
