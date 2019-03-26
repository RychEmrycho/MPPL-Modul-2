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

type ColorHandler struct {
	ColorUsecase manage_product.UseCase
}

func NewColorHandler(e *echo.Echo, colorusecase manage_product.UseCase) {
	handler := &ColorHandler{ColorUsecase: colorusecase}

	e.GET("/colors", handler.FetchColor)
	e.GET("/colors/:id", handler.GetById)
	e.POST("/colors", handler.Store)
	e.PUT("/colors/:id", handler.Update)
	e.DELETE("/colors/:id", handler.Delete)
}

func (ph *ColorHandler) FetchColor(c echo.Context) error {
	listEl, err := ph.ColorUsecase.FetchColor()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *ColorHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.ColorUsecase.GetByIdColor(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *ColorHandler) Update(c echo.Context) error {
	var color_ Color
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	color_.ID = id
	err = c.Bind(&color_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForColorValid(&color_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ColorUsecase.UpdateColor(&color_, id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, color_)
}

func (ph *ColorHandler) Store(c echo.Context) error {
	var color_ Color

	err := c.Bind(&color_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForColorValid(&color_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ColorUsecase.StoreColor(&color_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, color_)
}

func (ph *ColorHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.ColorUsecase.DeleteColor(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForColorValid(p *Color) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}