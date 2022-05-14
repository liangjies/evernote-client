package request

// 查询及排序结构体
type SearchNoteParams struct {
	SearchKey  string `json:"searchKey"`  // 搜索词
	NotebookId uint   `json:"notebookId"` // 笔记所属笔记本
}
