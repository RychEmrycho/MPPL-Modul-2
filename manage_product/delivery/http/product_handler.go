package http

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/manage_product/delivery/utils"
	. "MPPL-Modul-2/models/manage_product"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductUsecase manage_product.UseCase
}

func NewProductHandler(e *echo.Echo, productusecase manage_product.UseCase) {
	handler := &ProductHandler{ProductUsecase: productusecase}
	r := mux.NewRouter()

	r.HandleFunc("/products/{id}", handler.update).Methods("PUT")

	e.GET("/products", handler.FetchProduct)
	e.GET("/products/:id", handler.GetById)
	e.POST("/products", handler.Store)
	e.PUT("/products/:id", handler.Update)
	e.DELETE("/products/:id", handler.Delete)
}

func (ph *ProductHandler) update(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var product_ Product
	err := json.NewDecoder(req.Body).Decode(&product_)

	if err != nil {
		dd, _ := json.Marshal(err)
		_, _ = rw.Write(dd)
		return
	}

	id_ := params["id"]
	idUint64, _ := strconv.ParseUint(id_, 10, 32)
	id := uint(idUint64)

	if ok, err := isRequestForProductValid(&product_); !ok {
		dd, _ := json.Marshal(err)
		_, _ = rw.Write(dd)
		return
	}

	err = ph.ProductUsecase.UpdateProduct(id, &product_)

	if err != nil {
		dd, _ := json.Marshal(err)
		_, _ = rw.Write(dd)
		return
	}

	dd, _ := json.Marshal(product_)
	_, _ = rw.Write(dd)

}

func (ph *ProductHandler) FetchProduct(c echo.Context) error {
	listEl, err := ph.ProductUsecase.FetchProduct()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *ProductHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.ProductUsecase.GetByIdProduct(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *ProductHandler,) Update(c echo.Context) error {
	var product_ Product
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)
	product_.ID =id
	err = c.Bind(&product_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForProductValid(&product_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ProductUsecase.UpdateProduct(id,&product_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product_)
}

func (ph *ProductHandler) Store(c echo.Context) error {
	var product_ Product

	err := c.Bind(&product_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForProductValid(&product_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ProductUsecase.StoreProduct(&product_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product_)
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.ProductUsecase.DeleteProduct(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForProductValid(p *Product) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}