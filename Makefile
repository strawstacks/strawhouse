TAG := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'github.com/bsthun/gut.Version=$(TAG)'
GOFLAGS := -trimpath -ldflags "$(LDFLAGS)"

.PHONY: bench
bench:
	go test -v -benchmem -bench . ./strawhouse-driver/...

.PHONY: protoc
protoc:
	protoc --go_out=./driver --go-grpc_out=./driver ./proto/**/*.proto

.PHONY: release
release:
	mkdir -p ./.local/release/
	cd backend && go get -u github.com/strawst/strawhouse-go@latest
	cd backend && env GOWORK=off GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o ../.local/release/strawhousebackd_linux_amd64 .
	cd backend && env GOWORK=off GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -o ../.local/release/strawhousebackd_linux_arm64 .
	env GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o ./.local/release/strawc_linux_amd64 ./command/
	env GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -o ./.local/release/strawc_linux_arm64 ./command/
	env GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -o ./.local/release/strawc_darwin_amd64 ./command/
	env GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) -o ./.local/release/strawc_darwin_arm64 ./command/
	env GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -o ./.local/release/strawc_windows_amd64.exe ./command/
	env GOOS=windows GOARCH=arm64 go build $(GOFLAGS) -o ./.local/release/strawc_windows_arm64.exe ./command/
