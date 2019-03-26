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

type VariantHandler struct {
	VariantUsecase manage_product.UseCase
}

func NewVariantHandler(e *echo.Echo, variantusecase manage_product.UseCase) {
	handler := &VariantHandler{VariantUsecase: variantusecase}

	e.GET("/variants", handler.FetchVariant)
	e.GET("/variants/:id", handler.GetById)
	e.POST("/variants", handler.Store)
	e.PUT("/variants/:id", handler.Update)
	e.DELETE("/variants/:id", handler.Delete)
}

func (ph *VariantHandler) FetchVariant(c echo.Context) error {
	listEl, err := ph.VariantUsecase.FetchVariant()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *VariantHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.VariantUsecase.GetByIdVariant(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *VariantHandler) Update(c echo.Context) error {
	var variant_ Variant
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	variant_.ID = id
	err = c.Bind(&variant_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForVariantValid(&variant_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.VariantUsecase.UpdateVariant(&variant_,id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, variant_)
}

func (ph *VariantHandler) Store(c echo.Context) error {
	var variant_ Variant

	err := c.Bind(&variant_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForVariantValid(&variant_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.VariantUsecase.StoreVariant(&variant_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, variant_)
}

func (ph *VariantHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.VariantUsecase.DeleteVariant(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForVariantValid(p *Variant) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}
