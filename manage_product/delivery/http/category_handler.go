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

type CategoryHandler struct {
	CategoryUsecase manage_product.UseCase
}

func NewCategoryHandler(e *echo.Echo, categoryusecase manage_product.UseCase) {
	handler := &CategoryHandler{CategoryUsecase: categoryusecase}

	e.GET("/categories", handler.FetchCategory)
	e.GET("/categories/:id", handler.GetById)
	e.POST("/categories", handler.Store)
	e.PUT("/categories/:id", handler.Update)
	e.DELETE("/categories/:id", handler.Delete)
}

func (ph *CategoryHandler) FetchCategory(c echo.Context) error {
	listEl, err := ph.CategoryUsecase.FetchCategory()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *CategoryHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.CategoryUsecase.GetByIdCategory(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *CategoryHandler) Update(c echo.Context) error {
	var category_ Category
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)
	category_.ID = id
	err = c.Bind(&category_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForCategoryValid(&category_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.CategoryUsecase.UpdateCategory(&category_,id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, category_)
}

func (ph *CategoryHandler) Store(c echo.Context) error {
	var category_ Category

	err := c.Bind(&category_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForCategoryValid(&category_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.CategoryUsecase.StoreCategory(&category_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, category_)
}

func (ph *CategoryHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.CategoryUsecase.DeleteCategory(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForCategoryValid(p *Category) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}