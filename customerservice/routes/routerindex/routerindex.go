package routerindex

import (
	"github.com/alactic/ecommerce/customerservice/routes/customers"
	"github.com/gorilla/mux"
)

func Routerindex(router *mux.Router) {
	customers.Customers(router)
}
