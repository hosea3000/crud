package dto

import "math"

type ListResp[T any] struct {
	List       []T      `json:"list"`
	Pagination PageResp `json:"pagination"`
}

type PageResp struct {
	Page       int `json:"pageNum"`
	PageSize   int `json:"pageSize"`
	TotalPages int `json:"totalPages"`
	TotalCount int `json:"totalCount"`
}

// 抄pocketbase的查询协议
type CommonQueryReq struct {
	Filter string `json:"filter" form:"filter"`
}

type PageReq struct {
	SortBy    string `json:"sortBy" form:"sortBy"`
	SortOrder string `json:"sortOrder" form:"sortOrder" default:"asc"`
	PageNum   int    `json:"pageNum" form:"pageNum" default:"1"`
	PageSize  int    `json:"pageSize" form:"pageSize" default:"20"`
}

// PageRes 分页返回结构，使用范型
type PageRes[T any] struct {
	Data       []T   `json:"data"`
	TotalCount int64 `json:"totalCount"`
	PageNum    int   `json:"pageNum"`
	PageSize   int   `json:"pageSize"`
}

func CalPageResp(page, pageSize, total int) PageResp {
	r := PageResp{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
		TotalCount: total,
	}
	return r
}
