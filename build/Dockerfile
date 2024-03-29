FROM golang:1.17-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ARG project
ENV PROJECT $project

# Set the Go environment
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOARCH amd64

# Build the Go app
RUN go build -o ./out/app $project

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/out/app /app/app

# This container exposes port 8080 and 8081 to the outside world
EXPOSE 8080
EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["/app/app"]