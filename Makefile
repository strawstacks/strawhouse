.PHONY: protoc
protoc:
	protoc --go_out=./strawhouse-proto --go-grpc_out=./strawhouse-proto ./strawhouse-proto/**/*.proto

.PHONY: bench
bench:
	go test -v -benchmem -bench . ./strawhouse-driver/...

.PHONY: release
release:
	mkdir -p ./local/release/
	env GOOS=linux GOARCH=amd64 go build -o ./local/release/strawhousebackd_linux_amd64 ./backend/
	env GOOS=linux GOARCH=arm64 go build -o ./local/release/strawhousebackd_linux_arm64 ./backend/
	env GOOS=darwin GOARCH=amd64 go build -o ./local/release/strawc_darwin_amd64 ./command/
	env GOOS=darwin GOARCH=arm64 go build -o ./local/release/strawc_darwin_arm64 ./command/
