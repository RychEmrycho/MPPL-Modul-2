package usecase

import (
	"MPPL-Modul-2/models"
	"MPPL-Modul-2/product"
	"context"
	"time"
)

type productUsecase struct {
	productRepo    product.Repository
	contextTimeout time.Duration
}

func NewProductUsecase(repository product.Repository, duration time.Duration) product.UseCase {
	return &productUsecase{productRepo: repository, contextTimeout: duration}
}

func (productUsecase) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

func (productUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []*models.Product, nextCursor string, err error) {
	panic("implement me")
}

func (productUsecase) GetById(ctx context.Context, id uint) (*models.Product, error) {
	panic("implement me")
}

func (productUsecase) Store(ctx context.Context, p *models.Product) error {
	panic("implement me")
}

func (productUsecase) Update(ctx context.Context, p *models.Product) error {
	panic("implement me")
}
