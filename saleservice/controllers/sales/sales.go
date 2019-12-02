package sales

import (
	"context"
	"encoding/json"
	"log"
	"fmt"
	"net/http"

	proto "github.com/alactic/ecommerce/proto/user"
	"github.com/alactic/ecommerce/saleservice/models/sales"
	"github.com/alactic/ecommerce/saleservice/utils/connection"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/couchbase/gocb.v1"
)

type Sale = sales.Sales
var bucket *gocb.Bucket = connection.Connection()

func GetUserDetails(response http.ResponseWriter, request *http.Request) {
	service := micro.NewService(micro.Name("saleservice"))
	service.Init()

	client := proto.NewUserService("user1", service.Client())

	id := &proto.Request{Id: int64(45)}
	r, err := client.UserDetails(context.Background(), id)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	response.Write([]byte(`{"response": "` + fmt.Sprint(r) + `" }`))
}

//router.HandleFunc("/sale", CreateSaleEndpoint).Methods("POST")
func CreateSaleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var sale Sale
	_ = json.NewDecoder(request.Body).Decode(&sale)
	id := uuid.Must(uuid.NewV4()).String()
	sale.Type = "sale"
	sale.Id = id
	_, err := bucket.Insert(id, sale, 0)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	response.Write([]byte(`{ "id": "` + id + `"}`))
}

//router.HandleFunc("/sale/{id}", GetSaleEndpoint).Methods("GET")
func GetSaleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	routerParams := mux.Vars(request)
	var sale Sale
	sale.Id = routerParams["id"]
	_, err := bucket.Get(routerParams["id"], &sale)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(sale)
}

//router.HandleFunc("/sale", GetSalesEndpoint).Methods("GET")
func GetSalesEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var sales []Sale
	query := gocb.NewN1qlQuery("SELECT META().id, " + bucket.Name() + ".* FROM " + bucket.Name() + " WHERE type = 'sale'")
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	var row Sale
	for rows.Next(&row) {
		sales = append(sales, row)
	}
	json.NewEncoder(response).Encode(sales)
}
