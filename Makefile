proto:
	protoc -I api/proto/ api/proto/*.proto --go_out=. --go-grpc_out=.

clean:
	del pkg\api\*.pb.go

client:
	go run cmd/client/main.go

server:
	go run cmd/server/main.go