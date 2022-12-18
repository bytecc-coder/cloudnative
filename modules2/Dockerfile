FROM golang:1.18 AS builder
LABEL author=bytecc-coder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn \
    VERSION=v0.0.1
WORKDIR /var
COPY httpserver/main.go .
RUN go build /var/main.go

FROM alpine
COPY --from=builder /var/main /
EXPOSE 8080
ENTRYPOINT ["/main"]

