# quora

[![Build Status](https://img.shields.io/badge/build-1.01-brightgreen)](https://travis-ci.org/pibigstar/go-todo)

> 蓝山考核
>


## 1.项目结构
<details>
<summary>展开查看</summary>
<pre><code>
├─.idea
├─app
│  ├─action
│  │  └─cmd
│  │      ├─api
│  │      │  ├─etc
│  │      │  └─internal
│  │      │      ├─config
│  │      │      ├─handler
│  │      │      │  ├─agree
│  │      │      │  ├─collection
│  │      │      │  ├─follow
│  │      │      │  └─like
│  │      │      ├─logic
│  │      │      │  ├─agree
│  │      │      │  ├─collection
│  │      │      │  ├─follow
│  │      │      │  └─like
│  │      │      ├─svc
│  │      │      └─types
│  │      └─rpc
│  │          ├─action
│  │          ├─etc
│  │          ├─internal
│  │          │  ├─config
│  │          │  ├─logic
│  │          │  ├─server
│  │          │  └─svc
│  │          └─pb
│  ├─article
│  ├─comment
│  │  ├─cmd
│  │  │  ├─api
│  │  │  │  ├─etc
│  │  │  │  └─internal
│  │  │  │      ├─config
│  │  │  │      ├─handler
│  │  │  │      │  └─comment
│  │  │  │      ├─logic
│  │  │  │      │  └─comment
│  │  │  │      ├─svc
│  │  │  │      └─types
│  │  │  └─rpc
│  │  │      ├─comment
│  │  │      ├─etc
│  │  │      ├─internal
│  │  │      │  ├─config
│  │  │      │  ├─logic
│  │  │      │  ├─server
│  │  │      │  └─svc
│  │  │      └─pb
│  │  └─model
│  ├─mq
│  │  ├─asynq
│  │  │  ├─job
│  │  │  │  ├─etc
│  │  │  │  ├─internal
│  │  │  │  │  ├─config
│  │  │  │  │  ├─logic
│  │  │  │  │  │  ├─defer
│  │  │  │  │  │  └─immed
│  │  │  │  │  └─svc
│  │  │  │  └─types
│  │  │  ├─model
│  │  │  └─scheduler
│  │  │      ├─etc
│  │  │      └─internal
│  │  │          ├─config
│  │  │          ├─logic
│  │  │          └─svc
│  │  ├─nsq
│  │  └─rabbitmq
│  │      ├─api
│  │      └─rpc
│  ├─qa
│  │  ├─cmd
│  │  │  ├─api
│  │  │  │  ├─etc
│  │  │  │  └─internal
│  │  │  │      ├─config
│  │  │  │      ├─handler
│  │  │  │      │  ├─answer
│  │  │  │      │  └─question
│  │  │  │      ├─logic
│  │  │  │      │  ├─answer
│  │  │  │      │  └─question
│  │  │  │      ├─middleware
│  │  │  │      ├─svc
│  │  │  │      └─types
│  │  │  └─rpc
│  │  │      ├─etc
│  │  │      ├─internal
│  │  │      │  ├─config
│  │  │      │  ├─logic
│  │  │      │  ├─server
│  │  │      │  └─svc
│  │  │      ├─pb
│  │  │      └─qa
│  │  └─model
│  └─user
│      ├─cmd
│      │  ├─api
│      │  │  ├─etc
│      │  │  └─internal
│      │  │      ├─config
│  ├─jwt
│  ├─mq
│  ├─mysql
│  ├─redis
│  └─tool
└─test

</pre></code>
</details>

## 1.1 grafana
![img.png](doc/img.png)
![img_1.png](doc/img_1.png)

## 1.2 gogs

## 1.3 asynqmon
![img_3.png](doc/img_3.png)
## 1.4 prometheus
![img_2.png](doc/img_2.png)


## 2. 使用技术
- [x] 微服务架构(go-zero)
- [x] mysql
- [x] 邮箱服务
- [x] redis 
- [x] etcd 注册发现服务
- [x] docker-compose 
- [x] asynq 异步任务
- [x] grafa+premetheus 
- [x] grpc 密码自定义设置和拦截器




### 2.1 后端框架

### 2.2 安全性

- [x] 密码加盐加密
- [x] 防止xxs注入
- [x] JWT
- [ ] grpc CA认证
- [x] grpc 密码自定义设置和拦截器


### 2.3 实现功能


- [x] 用户模块
- [x] 问答模块
- [x] 评论模块
- [x] 动作模块
- [ ] feed流
-




  