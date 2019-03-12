package usecase

import (
	"MPPL-Modul-2/models"
	"MPPL-Modul-2/product"
	"time"
)

type productUsecase struct {
	productRepo    product.Repository
	contextTimeout time.Duration
}

func (*productUsecase) Fetch() (res []*models.Product, err error) {
	panic("implement me")
}

func (*productUsecase) GetById(id uint) (*models.Product, error) {
	panic("implement me")
}

func (*productUsecase) Update(p *models.Product) error {
	panic("implement me")
}

func (*productUsecase) Store(p *models.Product) error {
	panic("implement me")
}

func (*productUsecase) Delete(id uint) error {
	panic("implement me")
}

func NewProductUsecase(repository product.Repository, duration time.Duration) product.UseCase {
	return &productUsecase{productRepo: repository, contextTimeout: duration}
}