FROM golang:1.20 AS builder

WORKDIR /src
COPY . /src

RUN export GOPROXY=https://goproxy.io && make build

FROM alpine:3.15
LABEL MAINTAINER="xiaohubai@outlook.com"

RUN apk --no-cache add tzdata

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /src/bin/server /app
COPY --from=builder /src/configs/configs.yaml /app

EXPOSE 8000
EXPOSE 9000

CMD ["./server","-conf","configs.yaml"]
