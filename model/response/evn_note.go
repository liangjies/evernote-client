package response

type SysNoteSearchResponse struct {
	List      interface{} `json:"list"`
	SearchKey string      `json:"searchKey"`
	Total     int64       `json:"total"`
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
}
