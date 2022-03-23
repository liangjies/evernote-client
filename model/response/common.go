package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type NoteResult struct {
	List interface{} `json:"list"`
}

type AddResult struct {
	ID interface{} `json:"id"`
}
