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

type ImageHandler struct {
	ImageUsecase manage_product.UseCase
}

func NewImageHandler(e *echo.Echo, imageusecase manage_product.UseCase) {
	handler := &ImageHandler{ImageUsecase: imageusecase}

	e.GET("/images", handler.FetchImage)
	e.GET("/images/:id", handler.GetById)
	e.POST("/images", handler.Store)
	e.PUT("/images/:id", handler.Update)
	e.DELETE("/images/:id", handler.Delete)
}

func (ph *ImageHandler) FetchImage(c echo.Context) error {
	listEl, err := ph.ImageUsecase.FetchImage()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *ImageHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.ImageUsecase.GetByIdImage(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *ImageHandler) Update(c echo.Context) error {
	var image_ Image

	err := c.Bind(&image_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForImageValid(&image_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ImageUsecase.UpdateImage(&image_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, image_)
}

func (ph *ImageHandler) Store(c echo.Context) error {
	var image_ Image

	err := c.Bind(&image_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForImageValid(&image_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ImageUsecase.StoreImage(&image_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, image_)
}

func (ph *ImageHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.ImageUsecase.DeleteImage(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForImageValid(p *Image) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}