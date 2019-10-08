package old_base

import (
	"net/http"
	"runtime"

	"samh_common_lib/utils/log"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, retCode SamhResponseCode, rsp interface{}) {
	rspData := SamhResponse{}
	rspData.Code = retCode
	rspData.Msg = responseCodeToMsg[retCode]
	if retCode == SamhResponseCode_Ok {
		rspData.Data = rsp
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
