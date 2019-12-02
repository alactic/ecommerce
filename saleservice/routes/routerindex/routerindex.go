package routerindex

import (
	"github.com/alactic/ecommerce/saleservice/routes/sales"
	"github.com/gorilla/mux"
)

func Routerindex(router *mux.Router) {
	sales.Sales(router)
}
