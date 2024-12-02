# Atom-go

> Atom-go，Golang构建的轻量级API脚手架，内置JWT认证、用户角色权限管理，并支持登录与操作日志，简化后端开发流程

## 项目介绍

### 技术栈

- web框架：[Gin](https://gin-gonic.com/zh-cn/)
- 数据库：[Gorm](https://gorm.io/zh_CN/)
- 认证机制：[JWT](https://github.com/golang-jwt/jwt)

### 核心功能

- 权限管理：角色与菜单权限管理，确保API访问安全
- 日志记录：
	- 登录日志：记录用户登录系统的时间、IP等信息
	- 操作日志：记录用户对系统进行的操作，如创建、修改、删除等
- API接口：最基础的登录、增删改查、上传文件等Api接口
- 安全性：使用JWT进行用户认证，确保接口安全

### 项目状态
- 前端：无前端页面，项目仅提供API接口服务
- 后端：完全独立的后端服务，可与其他前端框架或平台集成

## 项目启动

### 开发环境
> golang >= 1.18

### 快速开始
###### clone 项目
```
git clone https://github.com/mengxiangyu996/atom-go.git
```
###### 进入项目目录
```
cd atom-go
```
###### 更新项目依赖项
```
go mod tidy
```
###### 创建配置文件
```
cp application-example.yaml application.yaml
```
###### 修改 `application.yaml` 数据库配置
``` yaml
# 数据库配置
mysql:
  # 地址
  host: localhost
  # 端口，默认为3306
  port: 3306
  # 数据库名称
  database: isme-go
  # 用户名
  username: root
  # 密码
  password: root
  # 编码
  charset: utf8mb4
  # 连接池最大连接数
  maxIdleConns: 10
  # 连接池最大打开连接数
  maxOpenConns: 100
```
###### 启动项目
```
go run main.go
```

## 接口文档
[Api在线文档](https://apifox.com/apidoc/shared-d640a698-2888-4105-b1c0-d86ebe3cfd39)
![](https://oss-liuchengtu.hudunsoft.com/userimg/36/36b52b7451e67d999857794ef3b24d23.png)

## 目录结构
```
├───config                        项目配置
├───internal                      包含应用程序的核心逻辑和业务逻辑
│   ├───common                    通用的工具和功能
│   │   ├───captcha               验证码
│   │   ├───curl                  http请求
│   │   ├───ip-address            ip地址
│   │   ├───logger                登录、操作日志
│   │   ├───password              密码
│   │   ├───upload                上传
│   │   └───utils                 工具类
│   ├───constant                  存放常量
│   │   ├───redis
│   │   └───regexp
│   ├───controller                控制器
│   │   └───admin
│   ├───dto                       请求响应的结构体
│   ├───middleware                中间件
│   ├───model                     模型，与数据库表字段对应
│   ├───router                    路由
│   ├───service                   服务层
│   ├───token                     jwt工具
│   └───validator                 验证器
└───pkg                           基础依赖包
    ├───dal                       数据访问（Gorm&Redis）
    ├───datetime                  时间类型（对time.Time的封装类型）
    └───response                  返回前端的响应JSON
```
