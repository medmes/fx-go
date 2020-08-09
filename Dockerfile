#Builder Image for CI
FROM golang:1.14-alpine3.12 AS builderBase
RUN apk add --no-cache git
# Set the Current Working Directory inside the container (temporary folder)
WORKDIR /tmp/fx-go
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
# Run Unit tests
RUN CGO_ENABLED=0 go test -v
# Build the Go app
RUN go build -o ./out/fx-go .

# Start fresh from a smaller image (slimmer)
#Runtime Image
FROM alpine:3.12
RUN apk add ca-certificates
COPY --from=builderBase /tmp/fx-go/out/fx-go /app/fx-go

# This container exposes port 8080 to the outside world
EXPOSE 8080
# Run the binary program produced by `go install`
CMD ["./out/fx-go"]