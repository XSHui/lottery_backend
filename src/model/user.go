package model

type LogInRequest struct {
	Action      string `json:"Action" comment: "LogIn"`
	PhoneNumber uint64 `json:"PhoneNumber" comment: "user phone number"`
}

type LogInResponse struct {
	Action  string `json:"Action" comment:"LogInResponse"`
	RetCode int    `json:"RetCode" comment:"return code"`
	Message string `json:"Message" comment:"return message"`
}

// TODO: Verification code
type UserExistRequest struct {
	Action      string `json:"Action" comment: "UserExist"`
	PhoneNumber uint64 `json:"PhoneNumber" comment: "user phone number"`
}

type UserExistResponse struct {
	Action  string `json:"Action" comment:"UserExistResponse"`
	RetCode int    `json:"RetCode" comment:"return code"`
	Message string `json:"Message" comment:"return message"`
	Exist   bool `json:"Exist" comment:"user exist"`
}
