# Stage 1: Build go binary
FROM golang:1.23.1-bookworm AS builder

# Set the working directory
WORKDIR /opt

# Copy the source code
COPY ./backend ./backend
COPY ./command ./command
COPY ./driver ./driver
COPY ./proto ./proto
COPY ./go.work ./go.work
COPY ./go.work.sum ./go.work.sum

# Install dependencies and build
RUN export PATH="$PATH:$(go env GOPATH)/bin" && \
    go build -o ./.local/strawhousebackd ./backend

# Stage 2: Create the final image
FROM alpine:3

# Install dependencies
RUN apk add --no-cache ca-certificates gcompat libstdc++

# Copy binary
COPY --from=builder /opt/.local/strawhousebackd /usr/local/bin/strawhousebackd

# Entrypoint
WORKDIR /opt
CMD ["strawhousebackd"]