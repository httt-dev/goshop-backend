package repository

import (
	"context"
	"timewise/model"
)

type CateRepo interface {
	SaveCate(ctx context.Context,cate model.Cate ) (model.Cate , error)
	DeleteCate(ctx context.Context,cateId string) error
	UpdateCate(ctx context.Context, cate model.Cate) error
	SelectCateById(ctx context.Context, cateId string) (model.Cate, error)
	SelectCates(ctx context.Context)([]model.Cate , error)
}