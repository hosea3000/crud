package crud

import (
	"crud-api/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Handler[T any] struct {
	db   *gorm.DB
	repo BaseRepository[T]
}

func NewHandler[T any](db *gorm.DB) *Handler[T] {
	return &Handler[T]{
		db:   db,
		repo: NewBaseRepo[T](db),
	}
}

func (h *Handler[T]) Create(ctx *gin.Context) {
	var entity T
	if err := ctx.ShouldBindJSON(&entity); err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	res, err := h.repo.Create(ctx, &entity)
	if err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}
	dto.HandleSuccess(ctx, res)
}

func (h *Handler[T]) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	var entity T
	if err := ctx.ShouldBindJSON(&entity); err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}
	if err := h.repo.Update(ctx, id, &entity); err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}
	dto.HandleSuccess(ctx, nil)
}

func (h *Handler[T]) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	if err := h.repo.Delete(ctx, id); err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}
	dto.HandleSuccess(ctx, nil)
}

func (h *Handler[T]) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	entity, err := h.repo.Get(ctx, id)
	if err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	dto.HandleSuccess(ctx, entity)
}

func (h *Handler[T]) List(ctx *gin.Context) {
	req, err := dto.BindQuery[dto.ListReq](ctx)
	if err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, dto.CreateRequestParamsError(err).Error(), nil)
		return
	}

	entityList, err := h.repo.List(ctx, req)
	if err != nil {
		dto.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}
	dto.HandleSuccess(ctx, entityList)
}
