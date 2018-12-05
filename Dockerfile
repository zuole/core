FROM golang:1.9 as builder
COPY --from=dosenv /go/src/github.com /go/src/github.com
WORKDIR /go/src/github.com/DOSNetwork/core
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go build -ldflags "-linkmode external -extldflags -static" -a -o clientNode main.go

FROM scratch
COPY --from=dosenv  /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/github.com/DOSNetwork/core/clientNode /
COPY --from=builder /go/src/github.com/DOSNetwork/core/onChain.json /
COPY --from=builder /go/src/github.com/DOSNetwork/core/offChain.json /
COPY --from=builder /go/src/github.com/DOSNetwork/core/credential /credential/
CMD ["/clientNode"]