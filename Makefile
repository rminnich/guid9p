all:
	protoc  --go_out=. --go-grpc_out=. proto/ninep.proto

badmaybe:
	mkdir -p generated
	protoc     --go_out=generated --go_opt=paths=source_relative,Mproto/ninep.proto=github.com/rminnich/guid9p/generated     --go-grpc_out=generated --go-grpc_opt=paths=source_relative,Mproto/ninep.proto=github.com/rminnich/guid9p/generated     proto/ninep.proto

protoc1.3:
	 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
