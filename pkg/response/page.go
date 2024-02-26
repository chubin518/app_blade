package response

type PageResult[T any] struct {
	List     []T `json:"list"`
	Total    int `json:"total"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

func Page[T any](list []T, total int, page, size int) *PageResult[T] {
	return &PageResult[T]{
		List:     list,
		Total:    total,
		PageNo:   page,
		PageSize: size,
	}
}
