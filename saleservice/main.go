package main

import (
	"fmt"
	// "github.com/alactic/ecommerce/userservice/utils/connection"
	"github.com/alactic/ecommerce/saleservice/utils/router"
	"gopkg.in/couchbase/gocb.v1"
)

var bucket *gocb.Bucket

func main() {
	fmt.Println("starting application")
	// bucket = connection.Connection()
	router.Router()
}
