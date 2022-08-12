proto:
	protoc ./pkg/auth/pb/*.proto --go_out=.
	protoc ./pkg/auth/pb/*.proto --go-grpc_out=.

	protoc ./pkg/song/pb/*.proto --go_out=.
	protoc ./pkg/song/pb/*.proto --go-grpc_out=.
server:
	go run cmd/main.go