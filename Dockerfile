FROM golang:1.20 AS builder

WORKDIR /src
COPY . /src

RUN export GOPROXY=https://goproxy.io \
    && go mod tidy \
    && mkdir -p bin \
    && make build

FROM alpine:3.15
LABEL MAINTAINER="xiaohubai@outlook.com"
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY --from=builder /src/bin/server /app
COPY --from=0 /src/configs/configs.yaml /app

EXPOSE 8000
EXPOSE 9000

CMD ["./server","-conf","configs.yaml"]
