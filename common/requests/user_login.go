package requests

type UserLoginRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword" form:"userPassword"`
}
