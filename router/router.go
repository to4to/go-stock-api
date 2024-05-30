package router

import (
	"github.com/to4to/go-stock-api/middleware"
	"github.com/gorilla/mux"
)


func Router()*mux.Router{


	router:=mux.NewRouter()

	router.HandleFunc("api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS")
router.HandleFunc("api/stock", middleware.GetAllStock).Methods("GET", "OPTIONS")

return router
}