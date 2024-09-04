all:
	mkdir -p generated
	protoc     --go_out=generated --go_opt=paths=source_relative,Mproto/ninep.proto=github.com/rminnich/guid9p/generated     --go-grpc_out=generated --go-grpc_opt=paths=source_relative,Mproto/ninep.proto=github.com/rminnich/guid9p/generated     proto/ninep.proto
