package proto

//go:generate bash -c "protoc --go_out=../pkg --go_opt=paths=source_relative --go-grpc_out=../pkg --go-grpc_opt=paths=source_relative $(find . -iname \"*.proto\")"
