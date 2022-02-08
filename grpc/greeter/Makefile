client:
	go run client.go
server:
	go run server.go
protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greeter/greeter.proto

