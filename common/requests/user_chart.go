package requests

type UserChartQuery struct {
	Key         string `json:"Id"`
	UserID      string `json:"userID"`
	ChartId     string `json:"chartId"`
	UserAccount string `json:"userAccount"`
	Goal        string `json:"goal"`
	ChartData   string `json:"chartData"`
	GenChart    string `json:"genChart"`
	GenResult   string `json:"genResult"`
	FreeCount   string `json:"freeCount"`
}
