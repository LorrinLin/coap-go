FROM ubuntu:16.04

MAINTAINER Yujia Lin <linyujia666@gmail.com>

# Get CoAP server
WORKDIR /root/coapserver

RUN go run server.go
EXPOSE 5683/udp
CMD ["sh", "/run.sh"]