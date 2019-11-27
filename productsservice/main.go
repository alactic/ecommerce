package main;

import (
	"log"
	"net/http"
	"context"
	
	proto "github.com/alactic/proto/user"
	"github.com/micro/go-micro"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)


func main() {
	route := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	service := micro.NewService(micro.Name("clientuser"))
	service.Init()

	client := proto.NewUserService("user", service.Client())

	id := &proto.Request{Id: int64(45)}
	r, err := client.UserDetails(context.Background(), id)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Println("Created: ", r)

	log.Fatal(http.ListenAndServe("0.0.0.0:8806", handlers.CORS(
		headers, methods, origins)(route)))
}