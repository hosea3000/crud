package crud

import (
	"context"
	"github.com/hosea3000/crud/dto"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Get(ctx context.Context, id int64) (*T, error)
	List(ctx context.Context, req *dto.ListReq) (*dto.ListResp[T], error)
	Create(ctx context.Context, entity *T) (*T, error)
	Update(ctx context.Context, id int64, updates any) error
	Delete(ctx context.Context, id int64) error
}

type BaseRepoImpl[T any] struct {
	db *gorm.DB
}

func NewBaseRepo[T any](db *gorm.DB) BaseRepository[T] {
	return &BaseRepoImpl[T]{
		db: db,
	}
}

func (r *BaseRepoImpl[T]) Get(ctx context.Context, id int64) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).First(&entity, id).Error
	return &entity, err
}

func (r *BaseRepoImpl[T]) List(ctx context.Context, req *dto.ListReq) (*dto.ListResp[T], error) {
	var entities []T
	var count int64

	query := r.db.Model(new(T)).WithContext(ctx)

	if req.SortBy != "" {
		query = query.Order(req.SortBy + " " + req.SortOrder)
	}

	if req.Filter != "" {
		filteredQuery, err := ApplyFilter(req.Filter, query)
		if err != nil {
			return nil, err
		}
		query = filteredQuery
	}

	err := query.Count(&count).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(req.PageSize * (req.PageNum - 1)).
		Limit(req.PageSize).
		Find(&entities).Error

	return &dto.ListResp[T]{
		List:       entities,
		Pagination: dto.CalPageResp(req.PageNum, req.PageSize, int(count)),
	}, err
}

func (r *BaseRepoImpl[T]) Create(ctx context.Context, entity *T) (*T, error) {
	err := r.db.WithContext(ctx).Create(entity).Error
	return entity, err
}

func (r *BaseRepoImpl[T]) Update(ctx context.Context, id int64, updates any) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id = ?", id).Updates(updates).Error
}

func (r *BaseRepoImpl[T]) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(new(T), id).Error
}
