protoc --go_out=./gen --go_opt=paths=source_relative \
    --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
    pb/greeter.proto

# buf build -o gen/moego.bin --as-file-descriptor-set pb

# we need googleapis
git clone https://github.com/googleapis/googleapis
GOOGLEAPIS_DIR=<your-local-googleapis-folder>
protoc -I${GOOGLEAPIS_DIR} -I. --include_imports --include_source_info \
    --descriptor_set_out=proto.pb pb/greeter.proto