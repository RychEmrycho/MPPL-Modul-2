package usecase

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
)

type productUsecase struct {
	productRepo manage_product.Repository
}

func (pu *productUsecase) Fetch() (res []*Product, err error) {
	res, err = pu.productRepo.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetById(id uint) (*Product, error) {
	res, err := pu.productRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) Update(p *Product) error {
	return pu.productRepo.Update(p)
}

func (pu *productUsecase) Store(p *Product) error {
	return pu.productRepo.Store(p)
}

func (pu *productUsecase) Delete(id uint) error {
	return pu.productRepo.Delete(id)
}

func NewProductUsecase(repository manage_product.Repository) manage_product.UseCase {
	return &productUsecase{productRepo: repository}
}
