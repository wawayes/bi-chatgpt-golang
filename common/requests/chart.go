package requests

type GenRequest struct {
	Goal      string `json:"goal"`
	ChartType string `json:"chartType"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChartQueryRequest struct {
	Page

	Key       string `json:"key"`
	Goal      string `json:"goal"`
	ChartType string `json:"chartType"`
	UserId    int    `json:"userId"`
}

type AddRequest struct {
	Page

	Key       string `json:"key"`
	Goal      string `json:"goal"`
	ChartType string `json:"chartType"`
	UserId    string `json:"userId"`
}
