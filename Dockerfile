# Build the cipher service binary
FROM golang:1.15.10-stretch as base

# add the working directory for the project
WORKDIR /go/src/goginsvc
# WORKDIR /app

# Copy the service code
COPY config config
COPY controllers controllers
COPY environment environment
COPY models models
COPY routes routes
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum

# building service binary at path discovergy/www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o go-gin ./main.go
# RUN go build ./cmd/watermark


## Using the multi-stage image to just run the binary
FROM alpine:3.7 as final

# Init working directory to root /
WORKDIR /

# Copy just the binary from the base image
COPY --from=base /go/src/goginsvc/go-gin .

# just an indication that this port will be exposed by this container
EXPOSE 8080

# command to run at the immediate start of the container
ENTRYPOINT ["./go-gin"]

# RUN ["go run", "watermark.go"]

# docker build -t watermark_images . -f images/watermark/Dockerfile
# docker run -d  -p 3333:8081 --name watermark_container watermark_images

