package model

type LotteryRequest struct {
	Action      string `json:"Action" comment: "Lottery"`
	UserId      string `json:"UserId" comment: "UserId"`
	PhoneNumber uint64 `json:"PhoneNumber" comment: "user phone number"`
}

type LotteryResponse struct {
	Action    string `json:"Action" comment:"LotteryResponse"`
	RetCode   int    `json:"RetCode" comment:"return code"`
	Message   string `json:"Message" comment:"return message"`
	Win       bool   `json:"Win" comment:"whether win"`
	PrizeName string `json:"PrizeName" comment:"prize"`
	PrizeId   string `json:"PrizeId" comment:"prize id"` // empty proze id for thank u
}
