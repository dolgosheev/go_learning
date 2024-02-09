#/etc/bin
mkdir -p ./gen/go/
protoc -I proto proto/**/**/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/