package v1

type PageResponse struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"pageSize"`
	Total    int64 `json:"total"`
	List     any   `json:"list"`
}
