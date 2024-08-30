all:
	protoc --go_out=. --go-grpc_out=. ninep.proto

