package router

import (
	"github.com/gorilla/mux"
	"github.com/to4to/go-stock-api/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("api/stock", middleware.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("api/deleteStock",middleware.DeleteStock).Methods("POST","OPTIONS")

	return router
}
