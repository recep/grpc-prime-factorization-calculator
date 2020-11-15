# Prime Factorization Calculator   
gRPC API written in Go

## Run  
```
go run server/server.go
go run client/client.go
```

## Demo  
![Image of Demo](https://user-images.githubusercontent.com/10357501/99191299-fda7e980-277c-11eb-9bf8-8e7f96d2237c.png)  

### Proto File Update  
```
protoc --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	./number.proto
```
**number.pb.go** -> protobuf  
**number_grpc.pb.go** -> grpc client and server functions  

### Resources  
https://grpc.io/docs/languages/go/quickstart/  
https://developers.google.com/protocol-buffers/docs/overview   
https://github.com/yakuter/go-grpc-protobuf