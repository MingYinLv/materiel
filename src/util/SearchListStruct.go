package util

type SearchFilter struct {
	size    int64  // 每页数量
	keyword string // 查询关键字
	page    int64  // 页码
	count   int64  // 总数量
}
