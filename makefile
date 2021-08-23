run-server:
	go run server/main.go

run-client:
	go run client/main.go

grpc:
	protoc --go_out=. --go-grpc_out=. server/proto/regulator/regulator.proto

grpc-micro:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. \
	--go_out=. --go_opt=Mserver/proto/regulator/regulator.proto=server/proto/regulator/ \
	server/proto/regulator/regulator.proto
