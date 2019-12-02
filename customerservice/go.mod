module github.com/alactic/ecommerce/customerservice

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/kr/pretty v0.1.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20191126235420-ef20fe5d7933 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/couchbase/gocb.v1 v1.6.4
	gopkg.in/couchbase/gocbcore.v7 v7.1.15 // indirect
	gopkg.in/couchbaselabs/gocbconnstr.v1 v1.0.4 // indirect
	gopkg.in/couchbaselabs/gojcbmock.v1 v1.0.3 // indirect
	gopkg.in/couchbaselabs/jsonx.v1 v1.0.0 // indirect
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
