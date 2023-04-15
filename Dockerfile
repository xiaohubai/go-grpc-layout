FROM debian:stable-slim
LABEL MAINTAINER="xiaohubai@outlook.com"

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /app

COPY bin/server /app
COPY configs/rbac_model.conf /app/configs

EXPOSE 8000
EXPOSE 9000

CMD ["./server","-env", "remote", "-chost", "172.21.0.2:8500", "-ctype", "consul" ,"-cpath", "prod/config.yaml"]
