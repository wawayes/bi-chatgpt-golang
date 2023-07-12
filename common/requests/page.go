package requests

type Page struct {
	Total    int `json:"total"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}
