package handle

import (
	"github.com/ufwfqpdgv/samh_common_lib/utils/log"

	"github.com/ufwfqpdgv/samh_common_lib/modle/user_modle"

	"github.com/ufwfqpdgv/samh_common_lib/base"

	"github.com/xormplus/xorm"
)

// func WriteUserLog(samhUserLogDB *xorm.Engine, uid int64, target int, content string, action int, status int, number int, number2 int, tid int64) (retCode base.SamhResponseCode) {
func WriteUserLog(samhUserLogDB *xorm.Engine, userLog *user_modle.UserLog) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ

	affected, err := samhUserLogDB.Table("user_log").Insert(userLog)
	if err != nil ||
		affected == 0 {
		log.Error(err)
		retCode = base.SamhResponseCode_ServerError
		return
	}

	return
}

// func WriteUserLogBySession(session *xorm.Session, uid int64, target int, content string, action int, status int, number int, number2 int, tid int64) (retCode base.SamhResponseCode) {
func WriteUserLogBySession(session *xorm.Session, userLog *user_modle.UserLog) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ

	affected, err := session.Table("samh_user_log.user_log").Insert(userLog)
	if err != nil ||
		affected == 0 {
		log.Error(err)
		retCode = base.SamhResponseCode_ServerError
		return
	}

	return
}
