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

type ReviewHandler struct {
	ReviewUsecase manage_product.UseCase
}

func NewReviewHandler(e *echo.Echo, reviewusecase manage_product.UseCase) {
	handler := &ReviewHandler{ReviewUsecase: reviewusecase}

	e.GET("/reviews", handler.FetchReview)
	e.GET("/reviews/:id", handler.GetById)
	e.POST("/reviews", handler.Store)
	e.PUT("/reviews/:id", handler.Update)
	e.DELETE("/reviews/:id", handler.Delete)
}

func (ph *ReviewHandler) FetchReview(c echo.Context) error {
	listEl, err := ph.ReviewUsecase.FetchReview()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (ph *ReviewHandler) GetById(c echo.Context) error {

	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := ph.ReviewUsecase.GetByIdReview(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (ph *ReviewHandler) Update(c echo.Context) error {
	var review_ Review
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)
	err = c.Bind(&review_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForReviewValid(&review_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	review_.ID = id
	err = ph.ReviewUsecase.UpdateReview(&review_,id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, review_)
}

func (ph *ReviewHandler) Store(c echo.Context) error {
	var review_ Review

	err := c.Bind(&review_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestForReviewValid(&review_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ph.ReviewUsecase.StoreReview(&review_)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, review_)
}

func (ph *ReviewHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	err = ph.ReviewUsecase.DeleteReview(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestForReviewValid(p *Review) (bool, error) {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}