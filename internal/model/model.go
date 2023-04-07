package model

type PageResp struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
