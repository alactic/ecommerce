package sales

import (
	"github.com/alactic/ecommerce/saleservice/controllers/sales"
	"github.com/alactic/ecommerce/saleservice/middlewares/authentication"
	"github.com/gorilla/mux"
)

func Sales(router *mux.Router) {
	router.HandleFunc("/sale-user", authentication.AuthMiddleware(sales.GetUserDetails)).Methods("GET")
	router.HandleFunc("/sales", authentication.AuthMiddleware(sales.CreateSaleEndpoint)).Methods("POST")
	router.HandleFunc("/sales/{id}", authentication.AuthMiddleware(sales.GetSaleEndpoint)).Methods("GET")
	router.HandleFunc("/sales", authentication.AuthMiddleware(sales.GetSalesEndpoint)).Methods("GET")
}
