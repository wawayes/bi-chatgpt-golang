package requests

type LoginRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword" form:"userPassword"`
}
