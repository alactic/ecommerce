package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	proto "github.com/alactic/ecommerce/proto/user"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
)

func main() {
	// route := mux.NewRouter()
	// headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	// origins := handlers.AllowedOrigins([]string{"*"})

	service := micro.NewService(micro.Name("clientuser1"))
	service.Init()

	client := proto.NewUserService("user1", service.Client())

	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "ecommerce product page",
		})

	})

	g.GET("/test", func(ctx *gin.Context) {
		id := &proto.Request{Id: int64(45)}
		r, err := client.UserDetails(context.Background(), id)
		if err != nil {
			log.Fatalf("Could not greet: %v", err)
		}
		log.Println("Created: ", r)
		ctx.JSON(http.StatusOK, gin.H{
			"response": fmt.Sprint(r),
		})

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	// log.Fatal(http.ListenAndServe("0.0.0.0:8806", handlers.CORS(
	// 	headers, methods, origins)(route)))
}
