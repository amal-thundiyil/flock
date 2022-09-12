grpc-server:
	@ go run server/main.go

grpc-client:
	@ go run client/main.go

grpc-proto:
	@ protoc --go-grpc_out=. --go_out=. *.proto
