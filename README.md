# ShiraChan

一个简单的接单系统

## 使用

### 环境要求

- 支持Go的平台

### 运行

- 配置环境变量
  ```
  secret=jwt密钥
  sql_host=数据库地址
  sql_port=数据库端口
  sql_name=数据库名
  sql_user=数据库用户
  sql_passwd=数据库密码
  ```
- 运行 
  `go run .\cmd\main\main.go`