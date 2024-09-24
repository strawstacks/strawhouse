.PHONY: protoc
protoc:
	protoc --go_out=./proto --go-grpc_out=./proto ./proto/**/*.proto

.PHONY: bench
bench:
	go test -v -benchmem -bench . backend/...