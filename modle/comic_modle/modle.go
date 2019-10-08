package comic_modle

import (
	. "samh_common_lib/modle/common_modle"
)

var (
	url     string
	timeOut int
)

type InternalFetchCartoonRequest struct {
	CartoonId int `form:"cartoon_id" json:"cartoon_id" binding:"required"`
}

type InternalFetchCartoonResponse struct {
	*Cartoon
}

type InternalGetCurrentReadingChapterRequest struct {
	Uid     int64 `form:"uid" json:"uid" binding:"required"`
	ComicId int64 `form:"comic_id" json:"comic_id" binding:"required"`
}

type InternalGetCurrentReadingChapterResponse struct {
	CurrentReadingChapterId int `json:"current_reading_chapter_id"`
}

type InternalGetUserBookNumberRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalGetUserBookNumberResponse struct {
	BookNumber int64 `json:"book_number"`
}

type InternalFetchTicketsRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalFetchTicketsResponse struct {
	MonthlyTicket             int                   `json:"monthly_ticket,omitempty"`
	RecommendedTicket         int                   `json:"recommended_ticket,omitempty"`
	TotalAdFreeReadDailyNum   int                   `json:"total_ad_free_read_daily_num"`
	RemainAdFreeReadDailyNum  int                   `json:"remain_ad_free_read_daily_num"`
	TotalReadingTicketNumber  int64                 `json:"total_reading_ticket_number"`
	CommonReadingTicketNumber int64                 `json:"common_reading_ticket_number"`
	ComicReadingTicketArr     []*ComicReadingTicket `json:"comic_reading_ticket_arr,omitempty"`
	RemainChapter             int                   `json:"remain_chapter,omitempty"`
	TotalChapter              int                   `json:"total_chapter,omitempty"`
}

type ComicReadingTicket struct {
	ComicId     int `json:"comic_id" xorm:"not null default 0 comment('所属漫画ID') index BIGINT(20)"`
	TotalNumber int `json:"total_number" xorm:"not null default 0 comment('总数量') INT(11)"`
}

type InternalBatchFetchTicketsRequest struct {
	Uid []int64 `form:"uid_arr" json:"uid_arr" binding:"required"`
}

type InternalBatchFetchTicketsResponse struct {
	TicketsArr []*InternalFetchTicketsResponse `json:"tickets_arr"`
}

type InternalAddReadingTicketRequest struct {
	Uid                         int64                    `form:"uid" json:"uid" binding:"required"`
	TicketType                  ReadingTicketTypeCode    `form:"ticket_type" json:"ticket_type" binding:"required"`
	Number                      int                      `form:"number" json:"number" binding:"required"`
	ComicId                     int                      `form:"comic_id" json:"comic_id" binding:"-"`
	Reading_ticket_overdue_time int64                    `form:"reading_ticket_overdue_time" json:"reading_ticket_overdue_time" binding:"-"`
	GetType                     ReadingTicketGetTypeCode `json:"get_type" binding:"-"`
}

type InternalAddReadingTicketResponse struct {
}

type InternalGetAllBuyVipGiveReadingTicketNumberRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalGetAllBuyVipGiveReadingTicketNumberResponse struct {
	Number int `json:"number,omitempty"`
}

type InternalGetGoldVipUnlockChapterNumberRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalGetGoldVipUnlockChapterNumberResponse struct {
	RemainChapter int `json:"remain_chapter,omitempty"`
	TotalChapter  int `json:"total_chapter,omitempty"`
}

type InternalGetReadingPageSeeAdRemianNumberRequest struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type InternalGetReadingPageSeeAdRemianNumberResponse struct {
	SeeAdRemainNumber int `json:"see_ad_remain_number,omitempty"`
}
