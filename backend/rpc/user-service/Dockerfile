FROM golang:1.10.3 as builder

WORKDIR /go/src/github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service

COPY . .

ENV http_proxy=$HTTP_PROXY
ENV https_proxy=$HTTP_PROXY

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
#get all the necessary dependencies
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /go/src/github.com/lucamatsumoto/go-sales-manager/backend/rpc/user-service .

CMD ["./user-service"]

