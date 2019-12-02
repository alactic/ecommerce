package router

import (
	"log"
	"net/http"
	"github.com/alactic/ecommerce/customerservice/routes/routerindex"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type server struct{}

func Router() {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api").Subrouter()
	routerindex.Routerindex(v1)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe("0.0.0.0:8808", handlers.CORS(headers, methods, origins)(router)))

}