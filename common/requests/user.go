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

type AddRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserName     string `json:"userName"`
	UserAvatar   string `json:"userAvatar"`
	UserRole     string `json:"userRole"`
	FreeCount    string `json:"freeCount"`
}
