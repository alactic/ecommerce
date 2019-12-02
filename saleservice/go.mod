module github.com/alactic/ecommerce/saleservice

go 1.13

require (
	github.com/alactic/ecommerce v0.0.0-20191127170529-11aca05ec427
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/micro/go-micro v1.17.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20191108234033-bd318be0434a
	golang.org/x/net v0.0.0-20191126235420-ef20fe5d7933 // indirect
	google.golang.org/grpc v1.25.1
	gopkg.in/couchbase/gocb.v1 v1.6.4
	gopkg.in/couchbase/gocbcore.v7 v7.1.15 // indirect
	gopkg.in/couchbaselabs/gocbconnstr.v1 v1.0.4 // indirect
	gopkg.in/couchbaselabs/gojcbmock.v1 v1.0.3 // indirect
	gopkg.in/couchbaselabs/jsonx.v1 v1.0.0 // indirect
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
