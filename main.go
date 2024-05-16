package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.73:3306)/banking")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepository
	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	rounter := mux.NewRouter()

	rounter.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	rounter.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":8000", rounter)

}
