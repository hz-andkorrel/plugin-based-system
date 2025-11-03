FROM golang:1.20-alpine AS builder
WORKDIR /src

# Copy go.mod and go.sum from the broker module and download deps there so
# module resolution happens correctly for the module at plugins/broker.
COPY plugins/broker/go.mod plugins/broker/go.sum ./
RUN go mod download

# Copy the rest of the repository; broker may import local packages.
COPY . .

# Build the broker binary. Adjust GOOS/GOARCH as needed.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags "-s -w" -o /bin/server ./plugins/broker

FROM alpine:3.18 AS final
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /bin/server /usr/local/bin/server

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/server"]
