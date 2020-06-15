# Version: 0.0.1
#FROM alpine
FROM golang

WORKDIR /home/redis
RUN wget http://download.redis.io/releases/redis-4.0.2.tar.gz
RUN tar xzf redis-4.0.2.tar.gz
WORKDIR /home/redis/redis-4.0.2
RUN make
RUN make install

WORKDIR /home/govideo_server
COPY . /home/govideo_server
RUN go env -w GOPROXY=https://goproxy.cn
RUN go build
EXPOSE 80

ENTRYPOINT ["./start.sh"]
