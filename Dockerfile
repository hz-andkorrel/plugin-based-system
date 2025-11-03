FROM golang:1.25-alpine AS builder
WORKDIR /src

# Copy the full repository so relative replace paths in plugins/broker/go.mod
# (for example replace ../common) resolve correctly during dependency download.
COPY . .

# Use the broker module's directory to download its dependencies.
WORKDIR /src/plugins/broker
RUN go mod download

# Build the broker binary. Adjust GOOS/GOARCH as needed.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags "-s -w" -o /bin/server ./

FROM alpine:3.18 AS final
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /bin/server /usr/local/bin/server

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/server"]
