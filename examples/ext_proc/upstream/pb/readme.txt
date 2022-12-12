protoc --go_out=./gen --go_opt=paths=source_relative \
    --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
    pb/greeter.proto

buf build -o gen/moego.bin --as-file-descriptor-set pb