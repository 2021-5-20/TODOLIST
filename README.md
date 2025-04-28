# Golang 备忘录管理系统 (To-Do List App)

✨ 一个基于 **Golang + Gin + GORM + Bootstrap5** 构建的小型备忘录管理系统，支持用户注册、登录鉴权（JWT）、备忘录的创建、查看、更新和删除（CRUD）。

---

## 项目功能 Features

- 用户注册、登录、退出
- JWT鉴权，Token保护接口
- 创建备忘录（添加内容）
- 查询个人备忘录（按时间倒序）
- 更新备忘录内容
- 删除备忘录（带二次确认）
- Bootstrap美化前端界面
- Loading动画提示
- 完整的前后端分离架构
- 错误处理与友好提示

---

## 技术栈 Tech Stack

- **后端**：Golang 1.21+、Gin框架、GORM ORM、JWT
- **数据库**：MySQL
- **前端**：Bootstrap 5、Vanilla JavaScript
- **项目管理**：模块分层（config / handlers / middleware / models / routes）

---

## 项目结构 Project Structure

```plaintext
t_d_list/
├── main.go                  # 程序入口，初始化数据库、路由
├── config/
│   └── config.go             # 数据库连接、JWT配置
├── models/
│   ├── user.go               # 用户表定义
│   └── memo.go               # 备忘录表定义
├── handlers/
│   ├── user_handlers.go      # 登录、注册处理器
│   └── memo_handlers.go      # 备忘录管理处理器
├── middleware/
│   └── jwt.go                # JWT鉴权中间件
├── routes/
│   └── routes.go             # 路由注册
├── frontend/
│   ├── login.html            # 登录页面
│   ├── register.html         # 注册页面
│   ├── memo.html             # 备忘录管理页面
│   └── js/
│       ├── main.js           # 登录逻辑
│       ├── register.js       # 注册逻辑
│       └── memo.js           # 备忘录管理逻辑
└── README.md                 # 项目说明文档
```

---

## 未来计划 Future Plans
- 引入Docker部署
- 添加图片上传功能
- 支持Memo标签分类
- 前端使用Vue3重构
- 后端增加Swagger API文档

---

Any advice can contact with me 
my email: gongzichen77@gmail.com
