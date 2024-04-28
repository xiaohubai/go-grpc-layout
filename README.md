## 目录结构
- gin 作为前端最直接的交互用于处理数据,图片,视频,文件请求等特殊请求处理.
- grpc 作为服务端为别人提供服务;作为客户端请求别人的微服务使用.

### 单仓
``` go
.
├── api                             // 定义proto,生成相关代码
│   ├── any                         // 通用
│   ├── grpc                        // grpc
│   └── http                        // http
├── cmd
│   ├── wire.go                     // 定义wire
│   ├── cmd.go                      // app实例
│   └── wire_gen.go                 // 依赖注入
├── internal                        // 内部逻辑
│   ├── biz                         // 业务组装层和定义repo接口
│   ├── consts                      // 全局或常量定义,避免魔数
│   ├── ecode                       // 业务错误码
│   ├── server
│   │   ├── grpc.go                 // grpc服务
│   │   ├── server.go               // 依赖注入
│   │   └── http.go                 // http服务
│   ├── service                     // 请求入口(接收/校验输入参数,返回结果)
│   └── data                        // 数据访问层
│       ├── gen                     // gorm gen工具生成代码
│       ├── model                   // gorm gen工具生成表结构
│       ├── data.go                 // mysql redis es的实例连接
│       └── user.go                 // 用户相关的数据操作
├── pkg                             // 公共组件
│   ├── consul                      // 服务注册/发现,远程配置
│   ├── email                       // 邮件组件
│   ├── holmes                      // 系统崩溃捕获
│   ├── jwt                         // 认证组件
│   ├── kafka                       // 消息队列
│   ├── metric                      // 埋点组件
│   ├── middleware                  // 中间件
│   │   ├── casbin.go               // 鉴权
│   │   ├── cors.go                 // 跨域
│   │   ├── jwt.go                  // 认证
│   │   ├── limiter.go              // 限流
│   │   ├── metrics.go              // 埋点
│   │   ├── recovery.go             // panic处理
│   │   ├── tracing.go              // 全链路跟踪
│   │   └── operation.go            // 操作记录
│   ├── tracing                     // 全链路组件
│   ├── utils                       // 工具集
│   ├── viper                       // 配置解析
│   ├── zap                         // 日志组件
│   └── redis                       // redis组件
├── third_party                     // 第三方proto
│   ├── errors
│   ├── google
│   ├── validate
│   └── openapi
├── configs                         // 配置相关
├── deploy                          // 部署相关
├── docs                            // 文档
│   ├── wiki                        // 开发文档
│   └── openapi                     // OpenAPI3.0在线文档
├── scripts                         // 脚本
├── logs                            // 日志
├── rbac_model.conf                 // 鉴权配置
├── go.mod
├── go.sum
├── main.go                         // 程序入口
├── Makefile                        // make命令
├── LICENSE                         // 版权
└── README.md
```
### 大仓(根据单仓,下移某些模块)
``` go
.
├── api
│   └── user
│       ├──any
│       ├──grpc
│       └──http
├── cmd
│   └── user
├── configs
│   └── user
│       └──conf
└── internal
   └── user
        ├──biz
        ├──consts
        ├──data
        ├──ecode
        ├──server
        └──service
```
## 组件访问接口
``` go
- http          127.0.0.1:8000
- grpc          127.0.0.1:9000
- openAPI       127.0.0.1:8000/docs
- mysql         172.21.0.11:3306
- redis         172.21.0.12:6379
- consul        172.21.0.13:8500
- jaeger        172.21.0.14:16686
- prometheus    172.21.0.15:9090
- grafana       172.21.0.16:3000 (admin admin)
- node-exporter 172.21.0.17:9100
- pyroscope     172.21.0.18:4040
- elasticsearch 172.21.0.19:9200
- kibana        172.21.0.20:5601 (elastic 1qaz!QAZ)
- kafka         172.21.0.21:9092
- kafka-ui      172.21.0.22:8080
win可以127.0.0.1/localhost访问wsl/docker的应用。
wsl内可以使用回环和docker的定制ip 172.21.0.13:8500访问容器
容器之间可以使用docker的定制ip 172.21.0.13:8500访问容器
容器内访问wsl的应用，需要使用wsl的ip.192.168.94.170
github.com/hyperjumptech/grule-rule-engine

```

## 主要功能
- http（github.com/gin-gonic/gin）
- grpc (github.com/go-kratos/kratos)
- mysql (gorm.io/gorm,gorm.io/gen)
- redis (github.com/redis/go-redis/v9)
- kafka (github.com/Shopify/sarama)
- elasticsearch (github.com/elastic/go-elasticsearch/v8)
- OpenAPI3.0 Swagger (github.com/swaggo/swag,github.com/google/gnostic)
- 跨域 (github.com/xiaohubai/go-grpc-layout/pkg/middleware/cors.go)
- 统一错误码 (github.com/xiaohubai/go-grpc-layout/internal/ecode)
- 统一返回格式 (github.com/xiaohubai/go-grpc-layout/pkg/utils/response)
- 业务异常处理 (panic,组件产生错误->上传jaeger/日志记录->邮件告警)
- 认证 (github.com/golang-jwt/jwt)
- 鉴权 (github.com/casbin/casbin/v2)
- 埋点  (github.com/prometheus/client_golang)
- 限流 （github.com/go-redis/redis_rate/v10）
- 日志 （go.uber.org/zap,gopkg.in/natefinch/lumberjack.v2）
- 参数校验（github.com/envoyproxy/protoc-gen-validate/validate）
- 全链路监控 (go.opentelemetry.io/otel)
- 服务注册/发现 (github.com/hashicorp/consul)
- 远程配置文件 (github.com/spf13/viper,github.com/hashicorp/consul)
- 邮件发送 (github.com/jordan-wright/email)
- 系统崩溃捕获 (mosn.io/holmes)
- 实时性能分析火焰图 (github.com/grafana/pyroscope)

## 网关(先记录一下,待开发)
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
    限流(根据用户规则进行拒绝策略,或者拒绝多余请求):
        计数器:计数达到阈值时，直接拒绝请求。
        漏斗模式:桶内流量过多时出现积压，满了，则开始拒绝请求(eg:大池子洗澡,人流量多了,不让进澡堂)
        令牌桶: 漏斗类似,中间人token放桶内,请求过来拿桶内token请求,拿不到拒绝请求(eg:独立房间洗澡,一个人一把钥匙,没有钥匙了,不让进澡堂)
    熔断(自我诊断服务,决定是否拒绝,放行)
    降级(弃卒保帅,保证主要功能)、
    缓存
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
- [x] gorm的gen自动化生成
- [x] 服务注册和发现
- [x] jenkins流水线
- [x] 埋点Metric
- [x] trace(使用jaeger基于opentelemetry标准)
- [x] 日志(zap)
- [x] grafana看板
- [x] 业务异常捕获(panic和运行pkg包中error发送邮件告警,日志记录,jaeger上报)
- [x] pyroscope实时性能分析火焰图(pull的方式)
- [x] holmes现场异常自动采样到文件,并发送邮件告警(文件附件)
- [x] redis分布式锁
- [x] kafka生产者
- [x] kafka通用消费处理器(根据配置文件的topic对应的func,自动匹配处理器,链式执行)
- [x] es数据存取
- [ ] openAPI和swagger
- [x] 热点缓存中间件(singleflight)
- [x] 业务产生的painc和pkg包组件使用的error,painc发送邮件告警
- [ ] grpc的中间件和gin补齐
- [ ] grafana看板导入dashboard

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
- servive层只处理 解析入参,组装 biz层需要的入参,调取biz层获取结果, 返回结果.一般用数据库model+分页信息
- gorm 更新操作 要特别注意 默认值, 再不确定更新那个struct字段时,要求请求参数全部有值,gorm进行map[string]interface{}指定全部请求参数更新.
- gorm统计 加上 delete_at  is null