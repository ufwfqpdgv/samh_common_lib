package samh_common_lib

import "github.com/ufwfqpdgv/utils"

func Init(urlRq string, timeOutRq int) {
	var err error
	url, timeOut, err = utils.CheckUrlTimeout(urlRq, timeOutRq)
	if err != nil {
		Panic(err)
	}
	err = CheckServerConnect(urlRq)
	if err != nil {
		// Panic(err) // 不能用Panic，如整个服务器重启，那互相依赖的服务谁也启动不了
		Error(err)
	}
}

func InitUser(rq *InternalInitRequest) (rsp *InternalInitResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalInitResponse{}

	rsp2 := &struct {
		Code SamhResponseCode      `json:"status"`
		Msg  string                `json:"msg"`
		Data *InternalInitResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/init"
	retCode = HttpGet(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func FetchUserPrivilegeInfo(rq *InternalUserPrivilegeInfoRequest) (rsp *InternalUserPrivilegeInfoResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalUserPrivilegeInfoResponse{}

	rsp2 := &struct {
		Code SamhResponseCode                   `json:"status"`
		Msg  string                             `json:"msg"`
		Data *InternalUserPrivilegeInfoResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/user_privilege_info"
	retCode = HttpGet(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func VipReward(rq *VipRewardRequest) (rsp *VipRewardResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &VipRewardResponse{}

	rsp2 := &struct {
		Code SamhResponseCode   `json:"status"`
		Msg  string             `json:"msg"`
		Data *VipRewardResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/reward"
	retCode = HttpPost(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func InternalRechargeApi(rq *InternalRechargeRequest) (rsp *InternalRechargeResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalRechargeResponse{}

	rsp2 := &struct {
		Code SamhResponseCode          `json:"status"`
		Msg  string                    `json:"msg"`
		Data *InternalRechargeResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/recharge"
	retCode = HttpPost(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func InternalOperationApi(rq *InternalOperationRequest) (rsp *InternalOperationResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalOperationResponse{}

	rsp2 := &struct {
		Code SamhResponseCode           `json:"status"`
		Msg  string                     `json:"msg"`
		Data *InternalOperationResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/operation"
	retCode = HttpPost(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func InternalAllInfoApi(rq *InternalAllInfoRequest) (rsp *InternalAllInfoResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalAllInfoResponse{}

	rsp2 := &struct {
		Code SamhResponseCode         `json:"status"`
		Msg  string                   `json:"msg"`
		Data *InternalAllInfoResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/all_info"
	retCode = HttpGet(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
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

func InternalBatchAllInfoApi(rq *InternalBatchAllInfoRequest) (rsp *InternalBatchAllInfoResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalBatchAllInfoResponse{}

	rsp2 := &struct {
		Code SamhResponseCode              `json:"status"`
		Msg  string                        `json:"msg"`
		Data *InternalBatchAllInfoResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/batch_all_info"
	retCode = HttpGet(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func InternalIsRechargedEffectiveVipApi(rq *InternalIsRechargedEffectiveVipRequest) (rsp *InternalIsRechargedEffectiveVipResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalIsRechargedEffectiveVipResponse{}

	rsp2 := &struct {
		Code SamhResponseCode                         `json:"status"`
		Msg  string                                   `json:"msg"`
		Data *InternalIsRechargedEffectiveVipResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/is_recharged_effective_vip"
	retCode = HttpGet(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}

func InternalUserPrivilegeInfoApi(rq *InternalUserPrivilegeInfoRequest) (rsp *InternalUserPrivilegeInfoResponse, retCode SamhResponseCode) {
	Debug(NowFunc())
	defer Debug(NowFunc() + " end")

	retCode = SamhResponseCode_Succ
	rsp = &InternalUserPrivilegeInfoResponse{}

	rsp2 := &struct {
		Code SamhResponseCode                   `json:"status"`
		Msg  string                             `json:"msg"`
		Data *InternalUserPrivilegeInfoResponse `json:"data"`
	}{}
	rsp2.Code = SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/vip/user_privilege_info"
	retCode = HttpGet(addr, rq, rsp2, timeOut, nil, 3)
	if retCode != SamhResponseCode_Succ {
		Error(NowFuncError())
		return
	}
	if rsp2.Code != SamhResponseCode_Succ {
		Error(NowFuncError())
		retCode = rsp2.Code
		return
	}
	rsp = rsp2.Data

	return
}
