package v1

type PageRequest struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"pageSize"`
}

type PageResponse struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"pageSize"`
	Total    int64 `json:"total"`
	List     any   `json:"list"`
}

type Warn struct {
	DateTime string `json:"DateTime"`
	TraceID  string `json:"traceID"`
	Error    string `json:"error"`
}
