# ShiraChan

一个简单的接单系统

## 使用

### 环境要求

- Go>=1.19(ent代码生成需要Go>=1.20)

### 运行

- 编辑配置文件`./config/config.yml`
  ```yaml
  secret: jwt密钥
  port: 端口
  ssl:
    enable: 是否启用https
    crt: 证书路径
    key: 私钥路径
  sql:
    user: 用户
    passwd: 密码
    host: 地址
    port: 端口
    name: 数据库名
  ```
- 运行 
  `go run ./cmd/main/main.go`