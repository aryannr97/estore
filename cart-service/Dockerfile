FROM golang:alpine

WORKDIR /go/src/order-server
COPY . .

RUN CGO_ENABLED=0 GOOS=${GOOS} go build -mod vendor cmd/main.go

EXPOSE 9090
CMD ["./main"]