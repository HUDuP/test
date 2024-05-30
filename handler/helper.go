package handler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Println(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, e)
	}
}
