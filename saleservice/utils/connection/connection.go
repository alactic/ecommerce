package connection

import (
	"fmt"
	"os"

	"gopkg.in/couchbase/gocb.v1"
)

func Connection() *gocb.Bucket {
	var bucket *gocb.Bucket
	var bucketName string

	fmt.Println("host bucket1 :: ", os.Getenv("COUCHBASE_BUCKET"))
	fmt.Println("host COUCHBASE_HOST :: ", os.Getenv("COUCHBASE_HOST"))
	fmt.Println("host COUCHBASE_USERNAME :: ", os.Getenv("COUCHBASE_USERNAME"))
	fmt.Println("host COUCHBASE_PASSWORD :: ", os.Getenv("COUCHBASE_PASSWORD"))

	cluster, _ := gocb.Connect(os.Getenv("COUCHBASE_HOST"))
    bucketName = os.Getenv("COUCHBASE_BUCKET")
 


	//  cluster, _ := gocb.Connect("http://localhost")
	// cluster, _ := gocb.Connect("http://192.168.0.107:31053")
	// cluster, _ := gocb.Connect("couchbase://192.168.0.100")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: os.Getenv("COUCHBASE_USERNAME"),
		Password: os.Getenv("COUCHBASE_PASSWORD"),
	})
	// bucket, error = cluster.OpenBucket("default", "")
	bucket, _ = cluster.OpenBucket(bucketName, "")
	fmt.Println("bucketName :: ", bucketName)
	fmt.Println("host bucket:: ", bucket)


	return bucket
}
