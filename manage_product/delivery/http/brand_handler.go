package http

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/manage_product/delivery/utils"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type BrandHandler struct {
	BrandUsecase manage_product.UseCase
}

func NewBrandHandler(e *echo.Echo, brandusecase manage_product.UseCase) {
	handler := &BrandHandler{BrandUsecase: brandusecase}

	e.GET("/brands", handler.FetchBrand)
	e.GET("/brands/:id", handler.GetById)
	e.POST("/brands", handler.Store)
	e.PUT("/brands/:id", handler.Update)
	e.DELETE("/brands/:id", handler.Delete)
}

func (ph *BrandHandler) FetchBrand(c echo.Context) error {
	listEl, err := ph.BrandUsecase.FetchBrand()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *BrandHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.BrandUsecase.GetByIdBrand(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *BrandHandler) Update(c echo.Context) error {
	var brand_ Brand

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	brand_.ID = id

	err = c.Bind(&brand_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForBrandValid(&brand_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.BrandUsecase.UpdateBrand(&brand_, id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, brand_)
}

func (ph *BrandHandler) Store(c echo.Context) error {
	var brand_ Brand

	err := c.Bind(&brand_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForBrandValid(&brand_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.BrandUsecase.StoreBrand(&brand_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, brand_)
}

func (ph *BrandHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.BrandUsecase.DeleteBrand(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForBrandValid(p *Brand) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}