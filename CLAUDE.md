# Butterfly Admin

Go 语言后台管理系统，基于 Gin 框架和 GORM ORM，采用 DDD（领域驱动设计）架构风格。

## 项目概述

这是一个轻量级的后台管理 API 服务，提供用户认证、菜单管理、角色权限等功能。

## 技术栈

- **语言**: Go 1.26
- **Web框架**: Gin (via butterfly v1.1.6 库)
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT (dgrijalva/jwt-go)
- **配置管理**: Viper
- **ID生成**: Snowflake

## 项目结构 (标准 Go 布局 + DDD 分层)

```
butterfly-admin/
├── cmd/butterfly-admin/          # 程序入口
│   └── main.go
├── internal/                     # 私有代码（internal 强制封装，禁止外部模块引用）
│   ├── starter/                  # 启动器和路由注册
│   │   └── starter.go
│   ├── application/              # 应用层 (服务编排)
│   │   ├── all_app.go            # 应用聚合
│   │   ├── login_app.go          # 登录服务
│   │   ├── sys_user_app.go       # 用户服务
│   │   ├── sys_menu_app.go       # 菜单服务
│   │   ├── sys_role_app.go       # 角色服务
│   │   └── sys_permission_app.go
│   ├── interfaces/               # 接口层 (HTTP handlers)
│   │   ├── login_handler.go      # 登录相关接口
│   │   ├── sys_user_handler.go   # 用户管理接口
│   │   ├── sys_menu_handler.go   # 菜单管理接口
│   │   ├── sys_role_handler.go   # 角色管理接口
│   │   └── middleware/          # 中间件
│   │       └── jwt_handler.go   # JWT认证中间件
│   ├── domain/                   # 领域层 (核心业务逻辑)
│   │   ├── entity/               # 领域实体
│   │   │   ├── sys_user.go
│   │   │   ├── sys_menu.go
│   │   │   ├── sys_role.go
│   │   │   ├── sys_permission.go
│   │   │   ├── sys_token.go
│   │   │   └── sys_menu_option.go
│   │   ├── repository/           # 仓库接口定义
│   │   └── security/             # 安全服务接口
│   ├── infrastructure/           # 基础设施层
│   │   ├── persistence/          # 仓库实现
│   │   │   ├── all_repository.go
│   │   │   └── sys_*_repository.go
│   │   └── security/             # 安全服务实现 (JWT, 加密)
│   ├── config/                   # 配置层
│   │   ├── config.go             # 配置聚合
│   │   ├── database/db.go        # 数据库连接
│   │   ├── auth/auth.go          # 权限配置
│   │   └── sequence/             # Snowflake ID配置
│   ├── common/                   # 公共组件
│   │   ├── base_entity.go        # 基础实体
│   │   ├── constant/             # 常量
│   │   ├── local_time.go         # 时间处理
│   │   └── gorm_logger.go        # GORM日志
│   └── types/                    # 类型定义和请求表单
│       ├── sys_login_type.go
│       ├── sys_user_type.go
│       ├── sys_menu_type.go
│       └── sys_role_type.go
├── configs/                      # 运行时配置文件
│   └── config.yml
├── migrations/                   # 数据库脚本
│   └── butterfly_admin.sql
├── vendor/                       # 外部依赖
├── Makefile                      # 构建脚本
├── Dockerfile                    # 容器构建
├── go.mod / go.sum
└── .gitignore
```

## 核心实体

| 实体 | 表名 | 说明 |
|------|------|------|
| SysUser | t_sys_user | 用户信息 |
| SysMenu | t_sys_menu | 系统菜单 |
| SysRole | t_sys_role | 角色 |
| SysPermission | t_sys_permission | 权限 |
| SysToken | t_sys_token | 登录令牌 |
| SysMenuOption | t_sys_menu_option | 菜单选项 |

## API 路由

### 认证相关 (无需JWT)
- `POST /api/login` - 用户登录
- `POST /api/logout` - 用户登出
- `POST /api/refresh` - 刷新令牌
- `GET /api/currentUser` - 获取当前用户信息

### 菜单管理
- `GET /api/sys/menu/withOption` - 获取菜单及选项 (公共)
- `GET /api/sys/menu/refresh` - 刷新菜单 (忽略认证)

### 用户管理
- `GET /api/sys/user/all` - 获取所有用户 (公共)

### 角色管理
- `GET /api/sys/role/all` - 获取所有角色 (公共)

## 配置文件

配置文件位于 `configs/config.yml`：

```yaml
db:
  dsn: 'root:root@tcp(127.0.0.1:3306)/butterfly_admin?charset=utf8mb4&parseTime=True&loc=Local'

server:
  engineMode: 'release'    # gin模式: debug/release
  serverAddr: :8088        # 服务端口
  methodOverride: true

auth:
  ignorePath:              # 完全忽略的路径
  ignorePrefixPath:        # 忽略认证前缀路径
    - GET - /api/sys/menu/refresh
  commonPath:              # 公共路径 (需登录但无权限检查)
    - GET - /api/sys/menu/withOption
    - GET - /api/sys/role/all
    - GET - /api/sys/user/all
```

## 架构说明

项目采用经典 DDD 四层架构，并使用 Go 标准 `internal/` 布局封装私有代码：

1. **Interfaces层**: HTTP处理器，接收请求、参数验证、调用应用服务
2. **Application层**: 服务编排，协调领域对象完成业务流程
3. **Domain层**: 领域核心，实体定义、仓库接口、业务规则
4. **Infrastructure层**: 技术实现，数据库操作、外部服务集成

依赖方向: cmd → internal/starter → {Interfaces, Application, Infrastructure} → Domain ← {types, common}

## 开发指南

### 运行项目

> 配置文件 `configs/config.yml` 和日志目录 `logs/` 相对进程工作目录读取，**必须在项目根目录执行**。

```bash
go run ./cmd/butterfly-admin
```

### 构建项目

```bash
make build    # 产出 bin/butterfly-admin
```

### 添加新功能模块

1. 在 `internal/domain/entity/` 定义实体
2. 在 `internal/domain/repository/` 定义仓库接口
3. 在 `internal/infrastructure/persistence/` 实现仓库
4. 在 `internal/application/` 创建应用服务
5. 在 `internal/types/` 定义请求/响应类型
6. 在 `internal/interfaces/` 创建处理器并注册路由

### 数据库迁移

项目使用 GORM 自动迁移，实体表名遵循 `t_sys_*` 前缀规范。数据库初始化脚本见 `migrations/butterfly_admin.sql`。

## 基础实体字段

所有实体继承 `common.BaseEntity`:

```go
type BaseEntity struct {
    Id        int64        // 主键 (Snowflake ID)
    CreatedAt *LocalTime   // 创建时间
    UpdatedAt *LocalTime   // 更新时间
    Deleted   DeleteStatus // 软删除标记 (0:未删除, 1:已删除)
}
```

## 代码规范

- 使用中文注释
- 遵循 Go 标准命名规范
- Handler 以 `*_handler.go` 命名
- Repository 以 `*_repository.go` 命名
- Application 以 `*_app.go` 命名
- Entity 使用 `t_sys_*` 表名前缀
