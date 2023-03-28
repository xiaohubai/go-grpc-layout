# [安装go环境](https://xiaohubai.github.io/docs/env)
[]()
# 安装kratos
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade

## make init

## 安装 goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl env check -i -f --verbose

## 开发步骤
### 添加并编写 Proto 文件
kratos proto add api/admin/v1/admin.proto

### 生成proto代码
kratos proto client api/admin/v1/admin.proto

### 生成service代码
kratos proto server api/admin/v1/admin.proto -t internal/service

## 生成配置文件代码
``` go
    1.更新configs/config.yaml，和对应的conf/conf.pb.go文件
    2.执行make config
```

## wire编写依赖关系
```go
//在wire.go文件下执行 wire

```


- api接口:      172.12.0.2:8888
- jaeger:       172.12.0.2:16686
- consul:       172.12.0.2:8500
- prometheus    172.12.0.2:9090
- grafana       172.12.0.2:3000 （admin admin）
- kibana        172.12.0.2:5601

//定义code码 和异常错误返回结果.
//权限 和认证


gin用于前端交互,正常数据库请求,和图片 视频请求.文件请求.

grpc 作为服务端给 别人提供服务.

grpc 作为客户端,向别的微服务请求.