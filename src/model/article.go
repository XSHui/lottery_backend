package model

import xmodel "lottery_backend/src/xorm/model"

type SubmitArticleRequest struct {
	Action      string `json:"Action" comment: "SubmitArticle"`
	UserId 		string `json:"UserId" comment: "user id"`
	Text        string `json:"Text" comment: "article"`
}

type SubmitArticleResponse struct {
	Action  string `json:"Action" comment:"SubmitArticleResponse"`
	RetCode int    `json:"RetCode" comment:"return code"`
	Message string `json:"Message" comment:"return message"`
}

type ListArticleRequest struct {
	Action      string `json:"Action" comment: "ListArticle"`
	Offset      int `json:"Offset" commend:"offset"`
	Limit       int `json:"Limit" comment:"limit"`
}

type ListArticleResponse struct {
	Action  string `json:"Action" comment:"SubmitArticleResponse"`
	RetCode int    `json:"RetCode" comment:"return code"`
	Message string `json:"Message" comment:"return message"`
	// TODO: Join User Info
	DataSet []xmodel.Article `json:"DataSet" comment:"article info"`
}
