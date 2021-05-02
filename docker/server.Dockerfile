FROM golang:1.16 as builder

WORKDIR /root/go/src/github.com/thejasn/grpc-web-istio-demo/
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o bin/server ./cmd/server/server.go

FROM scratch
WORKDIR /bin/
COPY --from=builder /root/go/src/github.com/thejasn/grpc-web-istio-demo/bin/server .
ENTRYPOINT [ "/bin/server" ]
EXPOSE 9000
