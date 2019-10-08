package operation_modle

import (
	. "samh_common_lib/modle/common_modle"

	"github.com/ufwfqpdgv/samh_common_lib/base"
)

var (
	url     string
	timeOut int
)

type InternalActivityRequest struct {
	base.SamhBaseRequest
	FetchActivityType FetchActivityTypeCode `form:"fetch_activity_type" json:"fetch_activity_type" binding:"-"`
	ActivityId        int64                 `form:"activity_id" json:"activity_id" binding:"-"`
	ActivityType      ActivityTypeCode      `form:"activity_type" json:"activity_type" binding:"-"`
	Area              string                `form:"area,default=CN" json:"area,default=CN" binding:"-"`
}

type InternalActivityResponse struct {
	ActivityArr []*Activity `json:"activity_arr"`
}
