package http

import (
	"MPPL-Modul-2/models"
	"MPPL-Modul-2/product"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ProductHandler struct {
	ProductUsecase product.UseCase
}

func NewProductHandler(e *echo.Echo, productusecase product.UseCase) {
	handler := &ProductHandler{ProductUsecase: productusecase}

	e.GET("/products", handler.FetchProduct)
	e.GET("/products/:id", handler.GetById)
	e.POST("/products", handler.Store)
	e.POST("/products/:id", handler.Update)
	e.DELETE("/products/:id", handler.Delete)
}

func (ph *ProductHandler) FetchProduct(c echo.Context) error {
	listEl, err := ph.ProductUsecase.Fetch()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *ProductHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.ProductUsecase.GetById(id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *ProductHandler) Update(c echo.Context) error {
	var product_ models.Product

	err := c.Bind(&product_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&product_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ProductUsecase.Update(&product_)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product_)
}

func (ph *ProductHandler) Store(c echo.Context) error {
	var product_ models.Product

	err := c.Bind(&product_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&product_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ProductUsecase.Store(&product_)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product_)
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.ProductUsecase.Delete(id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestValid(p *models.Product) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
