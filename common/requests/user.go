package requests

type LoginRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

type RegisterRequest struct {
	UserAccount   string `json:"userAccount"`
	UserPassword  string `json:"userPassword"`
	CheckPassword string `json:"checkPassword"`
}
