package sqldev

type SqlQueryForm struct {
	Project   int64  `json:"project"`
	Iid       string `json:"iid"`
	Sql       string `json:"sql"`
	Owner     string `json:"owner"`
	Schema    string `json:"schema"`
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
	NeedTotal bool   `json:"need_total"`
}
