package repository

import (
	"MPPL-Modul-2/models"
	"MPPL-Modul-2/product"
	"context"
	"github.com/jinzhu/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) product.Repository {
	return &productRepository{Conn}
}

func (productRepository) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

func (productRepository) Fetch(ctx context.Context, cursor string, num int64) (res []*models.Product, nextCursor string, err error) {
	panic("implement me")
}

func (productRepository) GetById(ctx context.Context, id uint) (*models.Product, error) {
	panic("implement me")
}

func (productRepository) Store(ctx context.Context, p *models.Product) error {
	panic("implement me")
}

func (productRepository) Update(ctx context.Context, p *models.Product) error {
	panic("implement me")
}
