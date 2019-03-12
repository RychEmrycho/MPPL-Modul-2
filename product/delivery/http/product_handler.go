package http

import (
	"MPPL-Modul-2/models"
	"MPPL-Modul-2/product"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
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
}

func (ph *ProductHandler) FetchProduct(c echo.Context) error {
	listEl, _ := ph.ProductUsecase.Fetch()

	return c.JSON(http.StatusOK, listEl)
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
