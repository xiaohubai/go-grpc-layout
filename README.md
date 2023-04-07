# [安装go环境](https://xiaohubai.github.io/docs/env)
[]()
# 安装kratos
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade

## make init


## 开发步骤
### 添加并编写 Proto 文件
kratos proto add api/admin/v1/admin.proto

### 生成proto代码
kratos proto client api/admin/v1/admin.proto

### 生成service代码
kratos proto server api/admin/v1/admin.proto -t internal/service

## 生成配置文件代码
``` go
    1. 更新conf.yaml ,conf.proto
    2. kratos proto client conf/conf.proto
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




proto:
    http:定义入参 反参  返回码 ,参数校验
    grpc 定义入参 反参  返回码 ,参数校验


gin用于前端交互,正常数据库请求,和图片 视频请求.文件请求.

grpc 作为服务端给 别人提供服务.

grpc 作为客户端,向别的微服务请求.

internal目录下
    去除conf,根目录下新增conf/conf.proto
    新增consts定义常量,避免魔数
    新增model定义数据库模型,方便后期工具生成
TODO
    权限 和认证
    //定义code码 和异常错误返回结果.
   // zap日志

TODO
 接受信号,处理grpc和http

TODO
1.服务实例是否携带配置
    消耗大,
2.读取远程配置 到全局实例

