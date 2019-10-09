package samh_common_lib

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ufwfqpdgv/log"
)

func SendResponse(c *gin.Context, retCode SamhResponseCode, rsp interface{}) {
	rspData := &SamhResponse{}
	rspData.Code = retCode
	rspData.Msg = responseCodeToMsg[retCode]
	if retCode == SamhResponseCode_Succ {
		rspData.Data = rsp
	}

	c.JSON(http.StatusOK, rspData)
}

func SendResponseWithMsg(c *gin.Context, retCode SamhResponseCode, msg string, rsp interface{}) {
	rspData := &SamhResponse{}
	rspData.Code = retCode
	rspData.Msg = msg
	if retCode == SamhResponseCode_Succ {
		rspData.Data = rsp
	}

	c.JSON(http.StatusOK, rspData)
}

func SendResponseCustom(c *gin.Context, code interface{}, msg string, rsp interface{}) {
	rspData := &SamhResponseCustom{}
	rspData.Code = code
	rspData.Msg = msg
	switch code.(type) {
	case SamhResponseCode:
		if code == SamhResponseCode_Succ {
			rspData.Data = rsp
		}
	case bool:
		if code == true {
			rspData.Data = rsp
		}
	case int, int64:
		if code == 0 {
			rspData.Data = rsp
		}
	default:
		rspData.Data = rsp
	}

	c.JSON(http.StatusOK, rspData)
}

func SendResponseCustom2(c *gin.Context, code interface{}, msg string, rsp interface{}) {
	rspData := &SamhResponseCustom{
		Code: code,
		Msg:  msg,
		Data: rsp,
	}

	c.JSON(http.StatusOK, rspData)
}

func NowFunc() string {
	pc, _, _, _ := runtime.Caller(1)
	return "NowFunc:" + runtime.FuncForPC(pc).Name() + " "
}

func NowFuncError() string {
	pc, _, _, _ := runtime.Caller(1)
	return "NowFunc:" + runtime.FuncForPC(pc).Name() + " Error "
}

func RecoverFunc(c *gin.Context) {
	if rec := recover(); rec != nil {
		c.Header("Content-Type", "text/json; charset=utf-8")
		c.String(http.StatusInternalServerError, "[]")
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, false)
		// utils.CheckErrSendEmail(fmt.Errorf("recovery:%s\nstack:%s", rec, string(buf[:n])))
		log.Panicf("recovery:%s\nstack:%s", rec, string(buf[:n]))
	}
}

func FetchResponseCodesMsg(retCode SamhResponseCode) string {
	return responseCodeToMsg[retCode]
}
