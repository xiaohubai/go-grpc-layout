# 目录结构
- gin 作为前端最直接的交互用于处理数据,图片,视频,文件请求.
- grpc 作为服务端为别人提供服务;作为客户端向别人的微服务请求.

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
└── third               // api依赖的第三方proto
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
└── third               // api依赖的第三方proto
```

# 主要功能
- gin框架（github.com/gin-gonic/gin）
- kratos微服务框架 (github.com/go-kratos/kratos)
- jwt认证（github.com/golang-jwt/jwt）
- casbin鉴权（github.com/casbin/casbin/v2）
- gorm数据库组件及gentool代码生成（gorm.io/gorm,gorm.io/gen）
- viper实时解析检测配置文件（github.com/spf13/viper）
- swagger 接口文档生成 （github.com/swaggo/swag）
- redis组件 （github.com/go-redis/redis）
- zap日志定制化 （go.uber.org/zap）
- validate参数校验（github.com/envoyproxy/protoc-gen-validate/validate）
- jaeger全链路监控 opentelemetry+jaeger（go.opentelemetry.io/otel）
- prometheus埋点 （github.com/prometheus/client_golang）
- 分布式接口限流 （github.com/go-redis/redis_rate/v9）
- consul服务注册/发现，远程配置文件 (github.com/hashicorp/consul)

# 组件访问接口
``` go
- api           172.21.0.2:8000
- jaeger        172.21.0.2:16686
- consul:       172.21.0.2:8500
- prometheus    172.21.0.2:9090
- grafana       172.21.0.2:3000 （admin admin）
```
## 设计图
![](https://github.com/xiaohubai/go-grpc-layout/blob/master/docs/user-req-resp.png)
### 网关
``` go
服务管理
    服务发现 服务注册 健康检查
配置管理
    版本管理
API 元信息管理
    路由匹配(前缀 精准 正则 RESTful),文档 OpenAPI
流量管理
    灰度发布 流量复制 负载均衡
隔离保护
    限流、熔断、降级、缓存
访问控制
    统一鉴权、跨域、风控
可观测性
    QPS、P99 各埋点 上下游基础信息(容器id,环境,请求接口,返回码),
```
### 持续部署
```go
流水线(触发方式 代码检查 单侧 质量红线 构建镜像 人工确认 自动发布 通知方式)
制品仓库
部署容器(手动)

```
# TODO

## 基建
- [x] gin HTTP框架
- [x] grpc RPC框架
- [x] 入参校验validate
- [x] 统一业务错误码和请求返回格式数据
- [x] 权限
- [x] 认证
- [x] 接口限流
- [x] 跨域
- [x] 部署(快速脚本构建,运行)
- [x] 读取本地或远程配置文件及监听
- [x] 捕获panic和运行pkg包中error发送邮件告警,日志记录,jaeger上报
- [x] gorm的gen自动化生成
- [x] 服务注册和发现
- [x] jenkins流水线
- [x] 埋点metric(使用prometheus和上报每一接口的请求用于记录QPS,P99等指标)
- [x] trace使用jaeger(基于opentelemetry标准)
- [x] 日志使用zap
- [x] grafana看板
- [x] 捕获用户请求和结果到kafka,
- [ ] 捕捉瞬时异常的cpu 内存堆栈, 等收集信息到kafka
- [ ] kafka->es
- [ ] openAPI和swagger
- [ ] 写一个热点缓存中间件(使用到redis分布式锁):多个用户请求相同,只一个用户获取热点缓存,返回多个用户请求.singleflight
- [x] 产生的painc和pkg包的error发送邮件告警
- [ ] grpc的中间件和gin补齐
- [x] kafka通用消费处理器(根据配置文件的topic对应的func,自动匹配处理器,链式执行)
- [x] kafka生产,消费数据失败,埋点量级


## 业务
- [x] 字典序
- [x] 获取全部路由
- [x] 获取角色路由
- [x] 获取layout配置信息
- [x] 获取 设置用户信息
- [ ] 注册
- [x] 获取验证码
- [x] 下发token和验证

## 注意事项:
- **Note**:servive层只处理 解析入参,组装 biz层需要的入参,调取biz层获取结果, 返回结果.一般用数据库model+分页信息
- **Note**: gorm 做更新操作 要特别注意 默认值, 再不确定更新那个struct字段时,要求请求参数全部有值,gorm进行map[string]interface{}指定全部请求参数更新.

- **gorm** 统计 加上 delete_at  is null