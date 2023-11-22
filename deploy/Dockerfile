FROM golang:1.20.3 AS builder

WORKDIR /app
COPY . /app

RUN export GOPROXY=https://goproxy.io && make build

FROM debian:stable-slim
LABEL MAINTAINER="xiaohubai@outlook.com"

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /app

COPY --from=builder /app/server /app
COPY --from=builder /app/rbac_model.conf /app

EXPOSE 8000
EXPOSE 9000

CMD ["./server","-env", "remote", "-chost", "127.0.0.1:8500", "-ctype", "consul" ,"-cpath", "prod/config.yaml"]
