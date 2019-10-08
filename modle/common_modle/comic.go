package common_modle

type ReadingTicketTypeCode int

const (
	ReadingTicketTypeCode_None ReadingTicketTypeCode = iota
	ReadingTicketTypeCode_Common
	ReadingTicketTypeCode_Exclusive
)

type ReadingTicketGetTypeCode int

const (
	ReadingTicketGetTypeCode_Old ReadingTicketGetTypeCode = iota
	ReadingTicketGetTypeCode_GoldExchange
	ReadingTicketGetTypeCode_BuyVipGive
	ReadingTicketGetTypeCode_SeeReadingPageAd
)
