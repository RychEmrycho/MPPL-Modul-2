package utils

import (
	. "MPPL-Modul-2/manage_product/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}