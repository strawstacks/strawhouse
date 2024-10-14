.PHONY: protoc
protoc:
	protoc --go_out=./strawhouse-proto --go-grpc_out=./strawhouse-proto ./strawhouse-proto/**/*.proto

.PHONY: bench
bench:
	go test -v -benchmem -bench . ./strawhouse-driver/...

.PHONY: release
release:
	mkdir -p ./.local/release/
	@env GOOS=linux GOARCH=amd64 go build -o ./.local/release/strawhousebackd_linux_amd64 ./strawhouse-backend/
	@env GOOS=linux GOARCH=arm64 go build -o ./.local/release/strawhousebackd_linux_arm64 ./strawhouse-backend/
	@if [ "$(shell uname -s)" = "Darwin" ]; then \
	  if [ "$(shell uname -m)" = "x86_64" ]; then \
		env GOOS=darwin GOARCH=amd64 go build -o ./.local/release/strawc_darwin_amd64 ./strawhouse-command/; \
	  elif [ "$(shell uname -m)" = "arm64" ]; then \
		env GOOS=darwin GOARCH=arm64 go build -o ./.local/release/strawc_darwin_arm64 ./strawhouse-command/; \
	  fi; \
	fi
