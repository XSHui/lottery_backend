package model

import xmodel "lottery_backend/src/xorm/model"

type ListRecordRequest struct {
	Action string `json:"Action" comment: "ListRecord"`
	Offset int    `json:"Offset" commend:"offset"`
	Limit  int    `json:"Limit" comment:"limit"`
}

// TODO: return count(*)
type ListRecordResponse struct {
	Action  string `json:"Action" comment:"ListRecordResponse"`
	RetCode int    `json:"RetCode" comment:"return code"`
	Message string `json:"Message" comment:"return message"`
	// TODO: Join User Info
	DataSet []xmodel.Record `json:"DataSet" comment:"record info"`
}

type SubOneDayForRecordRequest struct {
	Action string `json:"Action" comment: "SubOneDayForRecord"`
}

type SubOneDayForRecordResponse struct {
	Action  string `json:"Action" comment:"SubOneDayForRecordResponse"`
	RetCode int    `json:"RetCode" comment:"return code"`
	Message string `json:"Message" comment:"return message"`
}
