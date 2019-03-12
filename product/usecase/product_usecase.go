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

func (pu *productUsecase) Fetch() (res []*models.Product, err error) {
	res, err = pu.productRepo.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetById(id uint) (*models.Product, error) {
	res, err := pu.productRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) Update(p *models.Product) error {
	return pu.productRepo.Update(p)
}

func (pu *productUsecase) Store(p *models.Product) error {
	return pu.productRepo.Store(p)
}

func (pu *productUsecase) Delete(id uint) error {
	return pu.productRepo.Delete(id)
}

func NewProductUsecase(repository product.Repository, timeout time.Duration) product.UseCase {
	return &productUsecase{productRepo: repository, contextTimeout: timeout}
}
