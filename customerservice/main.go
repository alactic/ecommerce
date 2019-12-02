package main

import (
	"fmt"
	 "github.com/alactic/ecommerce/customerservice/utils/router"
	// "gopkg.in/couchbase/gocb.v1"
)

// var bucket *gocb.Bucket

func main() {
	fmt.Println("starting application")
	// bucket = connection.Connection()
	router.Router()
}