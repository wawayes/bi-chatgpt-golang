package requests

type RegisterRequest struct {
	UserAccount   string `json:"userAccount"`
	UserPassword  string `json:"userPassword"`
	CheckPassword string `json:"checkPassword"`
}
