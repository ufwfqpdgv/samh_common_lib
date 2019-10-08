package comic_modle

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

// XXX httpClient 的初始化可以在这里定义跟做，然后下面的所有的utils.HttpGet、utils.HttpPost在最前面再加一个参数httpClient就可以了，这样就不用在 utils 里加 httpCleint 相关的初始化及根据 url 分类

func InternalFetchCartoonApi(rq *InternalFetchCartoonRequest) (rsp *InternalFetchCartoonResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalFetchCartoonResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode         `json:"status"`
		Msg  string                        `json:"msg"`
		Data *InternalFetchCartoonResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/cartoon"
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

func InternalGetCurrentReadingChapterApi(rq *InternalGetCurrentReadingChapterRequest) (rsp *InternalGetCurrentReadingChapterResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalGetCurrentReadingChapterResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode                     `json:"status"`
		Msg  string                                    `json:"msg"`
		Data *InternalGetCurrentReadingChapterResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/get_current_reading_chapter"
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

func InternalGetUserBookNumberApi(rq *InternalGetUserBookNumberRequest) (rsp *InternalGetUserBookNumberResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalGetUserBookNumberResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode              `json:"status"`
		Msg  string                             `json:"msg"`
		Data *InternalGetUserBookNumberResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/get_user_book_number"
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

func InternalFetchTicketsApi(rq *InternalFetchTicketsRequest) (rsp *InternalFetchTicketsResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalFetchTicketsResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode         `json:"status"`
		Msg  string                        `json:"msg"`
		Data *InternalFetchTicketsResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/tickets"
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

func InternalFetchBatchTicketsApi(rq *InternalBatchFetchTicketsRequest) (rsp *InternalBatchFetchTicketsResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalBatchFetchTicketsResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode              `json:"status"`
		Msg  string                             `json:"msg"`
		Data *InternalBatchFetchTicketsResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/batch_tickets"
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

func InternalAddReadingTicketApi(rq *InternalAddReadingTicketRequest) (rsp *InternalAddReadingTicketResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalAddReadingTicketResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode             `json:"status"`
		Msg  string                            `json:"msg"`
		Data *InternalAddReadingTicketResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/add_reading_ticket"
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

func InternalGetAllBuyVipGiveReadingTicketNumberApi(rq *InternalGetAllBuyVipGiveReadingTicketNumberRequest) (rsp *InternalGetAllBuyVipGiveReadingTicketNumberResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalGetAllBuyVipGiveReadingTicketNumberResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode                                `json:"status"`
		Msg  string                                               `json:"msg"`
		Data *InternalGetAllBuyVipGiveReadingTicketNumberResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/get_all_buy_vip_give_reading_ticket_number"
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

func InternalGetGoldVipUnlockChapterNumberApi(rq *InternalGetGoldVipUnlockChapterNumberRequest) (rsp *InternalGetGoldVipUnlockChapterNumberResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalGetGoldVipUnlockChapterNumberResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode                          `json:"status"`
		Msg  string                                         `json:"msg"`
		Data *InternalGetGoldVipUnlockChapterNumberResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/get_gold_vip_unlock_chapter_number"
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

func InternalGetReadingPageSeeAdRemianNumberApi(rq *InternalGetReadingPageSeeAdRemianNumberRequest) (rsp *InternalGetReadingPageSeeAdRemianNumberResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalGetReadingPageSeeAdRemianNumberResponse{}

	rsp2 := &struct {
		Code base.SamhResponseCode                            `json:"status"`
		Msg  string                                           `json:"msg"`
		Data *InternalGetReadingPageSeeAdRemianNumberResponse `json:"data"`
	}{}
	rsp2.Code = base.SamhResponseCode_ServerError
	addr := url + "/internal/api/v1/comics/get_reading_page_see_ad_remian_number"
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
