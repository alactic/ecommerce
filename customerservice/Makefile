build:
	# protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=.  proto/user/user.proto
	sudo docker build -t ecommerce-customerservice .

run: 
	sudo docker rm -f customerservice
	sudo docker run -it -p 8812:8807 -e COUCHBASE_HOST=http://192.168.0.101:30886/ -e COUCHBASE_BUCKET=default -e COUCHBASE_USERNAME=elvis -e COUCHBASE_PASSWORD=password --network="bridge" --name customerservice ecommerce-customerservice
	# sudo docker run -d -p 1111:8805  --network="bridge" --name userservice2 ecommerce-userservice

go:
	env COUCHBASE_HOST=http://localhost/ COUCHBASE_BUCKET=travel-sample COUCHBASE_USERNAME=softloft COUCHBASE_PASSWORD=password go run *.go
	# env COUCHBASE_HOST=http://192.168.0.101:30886/ COUCHBASE_BUCKET=default COUCHBASE_USERNAME=elvis COUCHBASE_PASSWORD=password go run *.go
    # env COUCHBASE_HOST=http://10.102.187.20/ COUCHBASE_BUCKET=default COUCHBASE_USERNAME=elvis COUCHBASE_PASSWORD=password go run *.go

network:
	sudo docker network ls

clean:
	go clean -modcache

fixuuid:
	go mod edit -replace=github.com/satori/go.uuid@v1.2.0=github.com/satori/go.uuid@master
	go mod tidy
	go build