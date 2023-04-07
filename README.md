# [安装go环境](https://xiaohubai.github.io/docs/env)
[]()
# 安装kratos
``` sh
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade
make init
```

# 开发步骤
- api添加并编写 Proto 文件
``` sh
kratos proto add api/grpc/v1/grpc.proto
kratos proto add api/http/v1/gin.proto
```

- api生成proto客户端代码
``` sh
kratos proto client api/grpc/v1/grpc.proto
kratos proto client api/http/v1/http.proto
```

- 生成service代码
``` sh
kratos proto server api/grpc/v1/grpc.proto -t internal/service
```
- 生成配置文件代码
``` sh
//编写configs.yaml,configs.proto
kratos proto client configs/configs.proto
```

## 组件访问接口
``` go
- api接口       172.12.0.2:8888
- jaeger        172.12.0.2:16686
- consul:       172.12.0.2:8500
- prometheus    172.12.0.2:9090
- grafana       172.12.0.2:3000 （admin admin）
- kibana        172.12.0.2:5601
```

# 目录结构
gin用于前端交互,正常数据库请求,和图片 视频请求.文件请求.
grpc 作为服务端给 别人提供服务.作为客户端,向别的微服务请求.

## 单仓
``` go
├── api                 // 定义proto,入参反参,参数校验
│   ├── grpc            // 生成grpc相关代码
│   └── http            // 生成http相关代码
├── cmd                 // main入口和wire注入
├── configs             // 配置相关,proto生成
├── deploy              // 部署相关
├── docs                // 文档
├── internal            // 业务逻辑
│   ├── biz             // 业务组装层和定义repo接口
│   ├── consts          // 常量定义,避免魔数
│   ├── dao             // 数据访问层,实现biz的repo
│   │   └── gen         // 工具生成gorm底层数据库交互代码
│   ├── errors          // 业务错误码
│   ├── model           // 数据库实体对象，以及输入与输出数据结构定义
│   ├── server          // http和grpc实例的创建和配置
│   └── service         // 接收/解析用户输入参数的入口
├── logs                // 日志
├── pkg                 // 公共组件
│   ├── utils           // 工具
├── scripts             // 脚本
└── third_party         // api 依赖的第三方proto
```

## 大仓
``` go
├── api                 // 定义proto
│   └── user            // 用户服务
│       └── grpc        // 生成grpc相关代码
│       └── http        // 生成http相关代码
├── cmd                 // main入口和wire注入
├── configs             // 配置相关,proto生成
├── deploy              // 部署相关
├── docs                // 文档
├── internal            // 业务逻辑
│   └── user            // 用户服务
│       ├── biz         // 业务组装层和定义repo接口
│       ├── consts      // 常量定义,避免魔数
│       ├── dao         // 数据访问层,实现biz的repo
│       │   └── gen     // 工具生成gorm底层数据库交互代码
│       ├── errors      // 业务错误码
│       ├── model       // 数据库实体对象，以及输入与输出数据结构定义
│       ├── server      // http和grpc实例的创建和配置
│       └── service     // 接收/解析用户输入参数的入口
├── logs                // 日志
├── pkg                 // 公共组件
│   ├── utils           // 工具
├── scripts             // 脚本
└── third_party         // api 依赖的第三方proto
```


TODO
- 权限 和认证
- 读取远程配置 到全局实例,监听并更新
- 接受信号,处理grpc和http