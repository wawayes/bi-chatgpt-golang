package requests

type UserRegisterRequest struct {
	UserAccount   string `json:"userAccount"`
	UserPassword  string `json:"userPassword"`
	CheckPassword string `json:"checkPassword"`
}
