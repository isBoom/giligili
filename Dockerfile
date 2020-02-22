FROM golang:alpine as base
ADD go/pkg /go/pkg
ADD go/giligili /giligili
WORKDIR /giligili
ENV GOPROXY="https://mirrors.aliyun.com/goproxy/"
ENV GO111MODULE="on"
RUN go build -o api_server

FROM alpine:3.7
RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf
WORKDIR /www
COPY --from=base /giligili/api_server /www
ADD go/giligili/conf /www/conf
ADD .env .
RUN chmod +x /www/api_server
CMD ["./api_server"]