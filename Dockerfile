# Stage 1: Build go binary
FROM golang:1.23-bookworm AS builder

# Set the working directory inside the container
WORKDIR /opt

# Copy the source code
COPY ./strawhouse-backend ./strawhouse-backend
COPY ./strawhouse-command ./strawhouse-command
COPY ./strawhouse-driver ./strawhouse-driver
COPY ./strawhouse-proto ./strawhouse-proto
COPY ./go.work ./go.work
COPY ./go.work.sum ./go.work.sum
COPY ./Makefile ./Makefile

# Install dependencies and build
RUN apt update > /dev/null && \
    apt install -y protobuf-compiler > /dev/null && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    export PATH="$PATH:$(go env GOPATH)/bin" && \
    go build -o ./.local/strawhousebackd ./strawhouse-backend

# Stage 2: Create the final image
FROM debian:bookworm

# Copy the compiled binary from the builder stage
COPY --from=builder /opt/.local/strawhousebackd /usr/local/bin/strawhousebackd

# Command to run the application
WORKDIR /opt
CMD ["strawhousebackd"]