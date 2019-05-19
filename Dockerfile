FROM golang:1.12 as builder

LABEL maintainer="Zichen Zhu <zic.zhu@gmail.com>"

WORKDIR /go/src/github.com/ziczhu/fibonacci_rest_api

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o /go/bin/fibonacci_rest_api .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/bin/fibonacci_rest_api .

CMD ["./fibonacci_rest_api", "start"]
