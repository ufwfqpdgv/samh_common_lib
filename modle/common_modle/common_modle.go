package common_modle

type UserLogAction int

const (
	UserLogAction_NewGiveAppreciate   UserLogAction = 1000
	UserLogAction_BuyComicChapter     UserLogAction = 1001
	UserLogAction_BuyDiamonds         UserLogAction = 1002
	UserLogAction_BuyPurificationCard UserLogAction = 1003
	UserLogAction_DiamondsUpgradeVip  UserLogAction = 1004
)
