# eMysql

完成| 项目    |完成| 项目
---|---------|---|-------
 ✔ | http服务| &nbsp; | 缓存
 ✔ | grpc服务| &nbsp; | 测试
 ✔ | 数据库  | &nbsp; | API管理

## 概念

### CRUD

- CREATE
- READ
- UPDATE
- DELETE

### LIMIT

- UPDATE
- DELETE

## 成果

### 运行服务

```
cd goms/eMysql/cmd

# 使用默认配置文件
go run . &  

# 使用指定配置文件
go run . & -cfgpath=../configs  
```

### 测试(使用) API

http
```
# 使用 http 方法 GET /ping
curl  localhost:8080/ping

# 使用 http 方法 GET /ping, 参数 message=xxx
curl  localhost:8080/ping?message=xxx
```

grpc
```
# 获取 grpc 方法列表
grpcurl -plaintext localhost:50051 list

# 使用 grpc 方法 service.goms.User/Ping, 参数 {"message":"xxx"}
grpcurl -plaintext -d '{"message":"xxx"}'  localhost:50051 service.goms.User/Ping
```
