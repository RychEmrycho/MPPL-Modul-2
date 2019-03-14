package usecase

import (
	. "MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"golang.org/x/sync/errgroup"
)

type productUsecase struct {
	repoBrand    RepositoryBrand
	repoCategory RepositoryCategory
	repoColor    RepositoryColor
	repoImage    RepositoryImage
	repoProduct  RepositoryProduct
	repoReview   RepositoryReview
	repoSize     RepositorySize
	repoVariant  RepositoryVariant
}

func (pu *productUsecase) FetchBrand() (res []*Brand, err error) {
	res, err = pu.repoBrand.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdBrand(id uint) (*Brand, error) {
	res, err := pu.repoBrand.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateBrand(p *Brand) error {
	return pu.repoBrand.Update(p)
}

func (pu *productUsecase) StoreBrand(p *Brand) error {
	return pu.repoBrand.Store(p)
}

func (pu *productUsecase) DeleteBrand(id uint) error {
	return pu.repoBrand.Delete(id)
}

func (pu *productUsecase) FetchCategory() (res []*Category, err error) {
	res, err = pu.repoCategory.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdCategory(id uint) (*Category, error) {
	res, err := pu.repoCategory.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateCategory(p *Category) error {
	return pu.repoCategory.Update(p)
}

func (pu *productUsecase) StoreCategory(p *Category) error {
	return pu.repoCategory.Store(p)
}

func (pu *productUsecase) DeleteCategory(id uint) error {
	return pu.repoCategory.Delete(id)
}

func (pu *productUsecase) FetchColor() (res []*Color, err error) {
	res, err = pu.repoColor.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdColor(id uint) (*Color, error) {
	res, err := pu.repoColor.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateColor(p *Color) error {
	return pu.repoColor.Update(p)
}

func (pu *productUsecase) StoreColor(p *Color) error {
	return pu.repoColor.Store(p)
}

func (pu *productUsecase) DeleteColor(id uint) error {
	return pu.repoColor.Delete(id)
}

func (pu *productUsecase) FetchImage() (res []*Image, err error) {
	res, err = pu.repoImage.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdImage(id uint) (*Image, error) {
	res, err := pu.repoImage.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateImage(p *Image) error {
	return pu.repoImage.Update(p)
}

func (pu *productUsecase) StoreImage(p *Image) error {
	return pu.repoImage.Store(p)
}

func (pu *productUsecase) DeleteImage(id uint) error {
	return pu.repoImage.Delete(id)
}

func (pu *productUsecase) FetchReview() (res []*Review, err error) {
	res, err = pu.repoReview.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdReview(id uint) (*Review, error) {
	res, err := pu.repoReview.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateReview(p *Review) error {
	return pu.repoReview.Update(p)
}

func (pu *productUsecase) StoreReview(p *Review) error {
	return pu.repoReview.Store(p)
}

func (pu *productUsecase) DeleteReview(id uint) error {
	return pu.repoReview.Delete(id)
}

func (pu *productUsecase) FetchSize() (res []*Size, err error) {
	res, err = pu.repoSize.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdSize(id uint) (*Size, error) {
	res, err := pu.repoSize.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateSize(p *Size) error {
	return pu.repoSize.Update(p)
}

func (pu *productUsecase) StoreSize(p *Size) error {
	return pu.repoSize.Store(p)
}

func (pu *productUsecase) DeleteSize(id uint) error {
	return pu.repoSize.Delete(id)
}

func (pu *productUsecase) FetchVariant() (res []*Variant, err error) {
	res, err = pu.repoVariant.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdVariant(id uint) (*Variant, error) {
	res, err := pu.repoVariant.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateVariant(p *Variant) error {
	return pu.repoVariant.Update(p)
}

func (pu *productUsecase) StoreVariant(p *Variant) error {
	return pu.repoVariant.Store(p)
}

func (pu *productUsecase) DeleteVariant(id uint) error {
	return pu.repoVariant.Delete(id)
}

func (pu *productUsecase) FetchProduct() (res []*Product, err error) {
	res, err = pu.repoProduct.Fetch()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) GetByIdProduct(id uint) (*Product, error) {
	res, err := pu.repoProduct.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pu *productUsecase) UpdateProduct(p *Product) error {
	return pu.repoProduct.Update(p)
}

func (pu *productUsecase) StoreProduct(p *Product) error {
	return pu.repoProduct.Store(p)
}

func (pu *productUsecase) DeleteProduct(id uint) error {
	return pu.repoProduct.Delete(id)
}

func (pu *productUsecase) provideProductDetails(data []*Product) ([]*Product, error) {

	g := errgroup.Group{}
	// Get brand's id
	mapBrands := map[uint]Brand{}

	for _, p := range data {
		mapBrands[p.BrandID] = Brand{}
	}

	// using goroutine to fetch brand's details
	chanBrand := make(chan *Brand)
	for brandID := range mapBrands {
		g.Go(func() error {
			res, err := pu.repoBrand.GetById(brandID)
			if err != nil {
				return err
			}

			chanBrand <- res
			return nil
		})
	}

	go func() {
		g.Wait()
		close(chanBrand)
	}()

	for brand := range chanBrand {
		if brand != nil {
			mapBrands[brand.ID] = *brand
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge brand to product
	for index, item := range data {
		if b, ok := mapBrands[item.Brand.ID]; ok {
			data[index].Brand = b
		}
	}

	return data, nil
}

func NewProductUsecase(
	repoBrand RepositoryBrand,
	repoCategory RepositoryCategory,
	repoColor RepositoryColor,
	repoImage RepositoryImage,
	repoProduct RepositoryProduct,
	repoReview RepositoryReview,
	repoSize RepositorySize,
	repoVariant RepositoryVariant) UseCase {
	return &productUsecase{
		repoBrand,
		repoCategory,
		repoColor,
		repoImage,
		repoProduct,
		repoReview,
		repoSize,
		repoVariant,
	}
}
