package dto

import (
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
)

func BindQuery[T any](ctx *gin.Context) (*T, error) {
	req := new(T)
	if err := defaults.Set(req); err != nil {
		return nil, err
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func BindBody[T any](ctx *gin.Context) (*T, error) {
	req := new(T)
	if err := defaults.Set(req); err != nil {
		return nil, err
	}
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, err
	}
	return req, nil
}

type ListReq struct {
	PageReq
	CommonQueryReq
}
