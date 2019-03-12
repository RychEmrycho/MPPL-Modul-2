package product

import (
	"MPPL-Modul-2/models"
	"context"
)

type UseCase interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []*models.Product, nextCursor string, err error)
	GetById(ctx context.Context, id uint) (*models.Product, error)
	Update(ctx context.Context, p *models.Product) error
	Store(ctx context.Context, p *models.Product) error
	Delete(ctx context.Context, id uint) error
}