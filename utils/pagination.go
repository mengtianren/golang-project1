package utils

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Pagination 分页结构体
type Pagination struct {
	Page     int    `form:"page,default=1" json:"page"`
	PageSize int    `form:"page_size,default=10" json:"page_size"`
	Sort     string `form:"sort,default=id desc" json:"sort"` // 新增排序字段
}

// GetOffset 计算偏移量
func (p *Pagination) GetOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.PageSize
}

// GetLimit 获取每页数量
func (p *Pagination) GetLimit() int {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	return p.PageSize
}

// GetTotal 获取总记录数
func (p *Pagination) GetTotal(db *gorm.DB) int64 {
	var total int64 = 0
	db.Count(&total)
	return total
}

// GetTotalPages 计算总页数
func (p *Pagination) GetTotalPages(total int64) int {
	if total == 0 {
		return 1
	}
	return int(math.Ceil(float64(total) / float64(p.PageSize)))
}

// Paginate 分页查询作用域
func (p *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := p.GetOffset()
		limit := p.GetLimit()
		return db.Offset(offset).Limit(limit).Order(p.Sort)
	}
}

// PagedResult 分页结果结构体
type PagedResult struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// NewPagedResult 创建分页结果
func NewPagedResult(c *gin.Context, db *gorm.DB, result interface{}) (*PagedResult, error) {
	var pagination Pagination
	if err := c.ShouldBind(&pagination); err != nil {
		return nil, err
	}
	total := pagination.GetTotal(db)
	fmt.Println(total)
	err := db.Scopes(pagination.Paginate()).Find(&result).Error
	if err != nil {
		// 这里可以添加错误处理逻辑，例如记录日志
		return nil, err
	}

	return &PagedResult{
		Data:       result,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: pagination.GetTotalPages(total),
	}, nil
}
