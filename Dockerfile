# Version: 0.0.2
FROM golang:1.14.4-alpine3.12 AS builder
RUN apk --no-cache add build-base
COPY . /code
RUN mkdir -p /usr/local/go/src/github.com/gongluck && \
    ln -s /code /usr/local/go/src/github.com/gongluck/govideo_server && \
    cd /usr/local/go/src/github.com/gongluck/govideo_server && \
    go env -w GOPROXY=https://goproxy.cn && \
    CGO_ENABLED=1 go build -a

FROM alpine:3.12
RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++
COPY --from=builder /usr/local/go/src/github.com/gongluck/govideo_server/govideo_server /govideo_server/app
COPY --from=builder /usr/local/go/src/github.com/gongluck/govideo_server/templates /govideo_server/templates
#RUN mkdir /govideo_server/videos
WORKDIR /govideo_server

CMD ["/govideo_server/app"]

EXPOSE 80
