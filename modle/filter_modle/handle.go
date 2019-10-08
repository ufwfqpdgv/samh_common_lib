package filter_modle

import (
	"samh_common_lib/base"
	"samh_common_lib/utils"
	"samh_common_lib/utils/log"
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

func FilterByRegex(rq *FilterByRegexRequest) (rsp *FilterByRegexResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &FilterByRegexResponse{}

	addr := url + "/filter_by_regex"
	retCode = utils.HttpGet(addr, rq, rsp, timeOut, nil, 3)
	if retCode != base.SamhResponseCode_Succ {
		log.Error(base.NowFuncError())
		return
	}

	return
}
