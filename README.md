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