FROM golang:1.13.0 as builder

RUN mkdir /build
ADD . /build
WORKDIR /build

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hello ./hello.go

FROM alpine:3.10

RUN apk --no-cache add ca-certificates

WORKDIR /usr/bin/
COPY --from=builder /build/hello .

ENTRYPOINT ["/usr/bin/hello"]
