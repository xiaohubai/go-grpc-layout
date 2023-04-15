# 目录结构
- gin 作为前端最直接交互处理数据,图片,视频,文件请求.
- grpc 作为服务端为别人提供服务;作为客户端,向别人的微服务请求.

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
│   ├── data            // 数据访问层,实现biz的repo
│   │    ├── model      // 数据库实体对象，以及输入与输出数据结构定义
│   │    └── gen        // gorm代码生成
│   ├── errors          // 业务错误码
│   ├── server          // http和grpc实例的创建和配置
│   └── service         // 接收/解析用户输入参数的入口
├── logs                // 日志
├── pkg                 // 公共组件
│   └── utils           // 工具
├── scripts             // 脚本
└── third_party         // api依赖的第三方proto
```
## 大仓
``` go
├── api                 // 定义proto
│   └── user            // 用户服务
│       ├── grpc        // 生成grpc相关代码
│       └── http        // 生成http相关代码
├── cmd                 // main入口和wire注入
├── configs             // 配置相关,proto生成
├── deploy              // 部署相关
├── docs                // 文档
├── internal            // 业务逻辑
│   └── user            // 用户服务
│       ├── biz         // 业务组装层和定义repo接口
│       ├── consts      // 常量定义,避免魔数
│       ├── data        // 数据访问层,实现biz的repo
│       │    ├── model  // 数据库实体对象，以及输入与输出数据结构定义
│       │    └── gen    // gorm代码生成
│       ├── errors      // 业务错误码
│       ├── server      // http和grpc实例的创建和配置
│       └── service     // 接收/解析用户输入参数的入口
├── logs                // 日志
├── pkg                 // 公共组件
│   └── utils           // 工具
├── scripts             // 脚本
└── third_party         // api依赖的第三方proto
```

## 主要功能
- gin web框架（github.com/gin-gonic/gin）
- jwt认证（github.com/golang-jwt/jwt）
- casbin鉴权（github.com/casbin/casbin/v2）
- gorm数据库组件及gentool代码生成（gorm.io/gorm,gorm.io/gen）
- viper实时解析检测配置文件（github.com/spf13/viper）
- swagger 接口文档生成 （github.com/swaggo/swag）
- redis组件 （github.com/go-redis/redis）
- zap日志定制化 （go.uber.org/zap）
- 参数校验（github.com/envoyproxy/protoc-gen-validate/validate）
- jaeger全链路监控opentelemetry（go.opentelemetry.io/otel）
- prometheus埋点 （github.com/prometheus/client_golang）
- 分布式接口限流 （github.com/go-redis/redis_rate/v9）
- consul服务注册/发现，远程配置文件 (github.com/hashicorp/consul)



## 组件访问接口
``` go
- api           172.21.0.2:8000
- jaeger        172.21.0.2:16686
- consul:       172.21.0.2:8500
- prometheus    172.21.0.2:9090
- grafana       172.21.0.2:3000 （admin admin）
```
