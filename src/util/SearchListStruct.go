package util

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strings"
)

type SearchFilter struct {
	Size    int64  // 每页数量
	Keyword string // 查询关键字
	Page    int64  // 页码
	SortBy  string // 排序字段
	Order   string // 排序顺序
	Type    int64  // 查询类型
}

func GetDefaultSearchFilter() SearchFilter {
	return SearchFilter{
		10,
		"",
		1,
		"id",
		"asc",
		0,
	}
}

func GetSearchFilter(c *gin.Context) SearchFilter {
	searchFilter := GetDefaultSearchFilter()
	if val, ok := c.GetQuery("page"); ok && govalidator.IsInt(val) {
		searchFilter.Page, _ = govalidator.ToInt(val)
	}
	if val, ok := c.GetQuery("size"); ok && govalidator.IsInt(val) {
		searchFilter.Size, _ = govalidator.ToInt(val)
	}
	if val, ok := c.GetQuery("type"); ok && govalidator.IsInt(val) {
		searchFilter.Type, _ = govalidator.ToInt(val)
	}
	if val, ok := c.GetQuery("sortby"); ok && strings.TrimSpace(val) != "" {
		searchFilter.SortBy = val
	}
	if val, ok := c.GetQuery("order"); ok && strings.TrimSpace(val) != "" {
		searchFilter.Order = val
	}
	if val, ok := c.GetQuery("keyword"); ok && strings.TrimSpace(val) != "" {
		searchFilter.Keyword = val
	}

	return searchFilter
}