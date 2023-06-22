
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

[English](./README-en.md) | 简体中文

						       
# 项目文档
[在线文档](https://www.platform-eff.com) : https://www.platform-eff.com

[初始化](https://www.platform-eff.com/guide/start-quickly/initialization.html)
						       
[从环境到部署教学视频](https://www.bilibili.com/video/BV1Rg411u7xH)

[开发教学](https://www.platform-eff.com/guide/start-quickly/env.html)

[交流社区](https://support.qq.com/products/371961)


# 重要提示

1.本项目从起步到开发到部署均有文档和详细视频教程

2.本项目需要您有一定的golang和vue基础

## 1. 基本介绍

### 1.1 项目介绍

> 研发效能平台：效能分析、提测平台、工单系统、测试管理、工具集合、接口测试、Web测试、性能测试等。平台是基于 [vue](https://vuejs.org) 和 [gin](https://gin-gonic.com) 开发的全栈前后端分离的开发基础平台，集成jwt鉴权，动态路由，动态菜单，casbin鉴权等。

[在线预览](http://demo.platform-eff.com): http://demo.platform-eff.com

测试用户名：admin

测试密码：123456

### 1.2 贡献指南
Hi! 首先感谢你使用 platform-eff。
platform-eff 的成长离不开大家的支持，如果你愿意为 platform-eff 贡献代码或提供建议，请阅读以下内容。

#### 1.2.1 Issue 规范
- issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。
									      
- 在提交 issue 之前，请搜索相关内容是否已被提出。

#### 1.2.2 Pull Request 规范
- 请先 fork 一份到自己的项目下，不要直接在仓库下建分支。

- commit 信息要以`[文件名]: 描述信息` 的形式填写，例如 `README.md: fix xxx bug`。

- 如果是修复 bug，请在 PR 中给出描述信息。

- 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并。

## 2. 使用说明

```
- node版本 > v16.8.3
- golang版本 >= v1.16
- IDE推荐：vscode
```

### 2.1 server项目

使用 `vscode` 等编辑工具，打开server目录，不可以打开 platform-eff 根目录

```bash

# 克隆项目
git clone https://github.com/txu2k8/platform-eff.git
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

### 2.2 web项目

```bash
# 进入web文件夹
cd web

# 安装依赖
npm install

# 启动web项目
npm run serve
```

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

##### （1）可以访问外国网站

````
go get -u github.com/swaggo/swag/cmd/swag
````

##### （2）无法访问外国网站

由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn) 或者 [goproxy.io](https://goproxy.io/zh/)

```bash
# 如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 如果嫌弃麻烦,可以使用go generate 编译前自动执行代码, 不过这个不能使用 `Goland` 或者 `Vscode` 的 命令行终端
cd server
go generate -run "go env -w .*?"

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.3.2 生成API文档

```` shell
cd server
swag init
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) 即可查看swagger文档

### 2.4 VSCode工作区

#### 2.4.1 开发

使用`VSCode`打开根目录下的工作区文件`platform-eff.code-workspace`，在边栏可以看到三个虚拟目录：`backend`、`frontend`、`root`。

#### 2.4.2 运行/调试

在运行和调试中也可以看到三个task：`Backend`、`Frontend`、`Both (Backend & Frontend)`。运行`Both (Backend & Frontend)`可以同时启动前后端项目。

#### 2.4.3 settings

在工作区配置文件中有`go.toolsEnvVars`字段，是用于`VSCode`自身的go工具环境变量。此外在多go版本的系统中，可以通过`gopath`、`go.goroot`指定运行版本。

```json
    "go.gopath": null,
    "go.goroot": null,
```

## 3. 技术选型

- 前端：用基于 [Vue](https://vuejs.org) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。
- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`MySql` > (5.7) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

## 4. 项目架构

### 4.1 系统架构图

![系统架构图](https://github.com/txu2k8/platform-eff/系统架构图.png)


## 5. 主要功能

- 首页：版本质量概览、测试进度概览、重要事项目标。
- 测试管理：自动化测试管理、执行、CICD
    - 功能测试：Web测试、接口测试、其他
    - 性能测试：性能测试工具化
- 提测管理：版本提测计划、提测变更内容
- 效能分析：测试、开发效能分析
- 工单系统：团队事项工单跟踪，例如环境搭建、XX项目答疑跟进等
- 实验室管理：
    - 物料管理：服务器、客户端、交换机、网卡等硬件设备管理
    - 环境管理：基于实验室物料，组合测试环境，供自动化测试选择
- 配置管理：
    - 系统配置：系统配置文件可前台修改
    - 常量配置：全局常量配置，字典
- 超级管理员：用户管理、用户组管理、角色管理、角色权限管理、菜单管理等

## 6 我的博客

> https://www.cnblogs.com/txu2k8/
>
>内有平台设计搭建教程（更新中...）。如果觉得项目对您有所帮助可以添加我的个人微信，欢迎您提出宝贵的需求。

