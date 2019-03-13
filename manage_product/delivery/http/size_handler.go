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

type SizeHandler struct {
	SizeUsecase manage_product.UseCase
}

func NewSizeHandler(e *echo.Echo, sizeusecase manage_product.UseCase) {
	handler := &SizeHandler{SizeUsecase: sizeusecase}

	e.GET("/sizes", handler.FetchSize)
	e.GET("/sizes/:id", handler.GetById)
	e.POST("/sizes", handler.Store)
	e.PUT("/sizes/:id", handler.Update)
	e.DELETE("/sizes/:id", handler.Delete)
}

func (ph *SizeHandler) FetchSize(c echo.Context) error {
	listEl, err := ph.SizeUsecase.FetchSize()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *SizeHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.SizeUsecase.GetByIdSize(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *SizeHandler) Update(c echo.Context) error {
	var size_ Size

	err := c.Bind(&size_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForSizeValid(&size_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.SizeUsecase.UpdateSize(&size_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, size_)
}

func (ph *SizeHandler) Store(c echo.Context) error {
	var size_ Size

	err := c.Bind(&size_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForSizeValid(&size_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.SizeUsecase.StoreSize(&size_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, size_)
}

func (ph *SizeHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.SizeUsecase.DeleteSize(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForSizeValid(p *Size) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}