## GoMicSrvActualProject（GOMI）

这是一个基于 Go 的微服务项目，当前包含：

- **user_srv**: 用户服务（后端）
- **user_web**: 用户相关的 Web / 网关服务（如果有）

本项目已经将敏感配置（例如数据库连接串）从代码中分离出来，使用配置文件和 `.gitignore` 来保护隐私信息，适合直接提交到 Git 仓库。

---

### 环境要求

- Go 1.20+（或与你本地一致的版本）
- MySQL 5.7+ / 8.0+
- Git（可选）

---

### 仓库结构（简要）

```text
GoMicSrvActualProject/
  ├── user_srv/              # 用户服务（后端）
  │   ├── global/            # 全局对象，如 DB 连接
  │   ├── handler/           # 业务处理（用户相关接口）
  │   ├── tests/             # 测试代码
  │   ├── main.go            # 用户服务入口
  │   ├── go.mod
  │   ├── go.sum
  │   ├── config.example.json# 示例配置（可提交）
  │   └── config.json        # 实际配置（本地，不提交）
  ├── user_web/              # 用户 Web/网关服务（结构依据你的实现）
  │   ├── main.go
  │   └── go.mod
  └── README.md              # 当前说明文件
```

---

### user_srv：用户服务

**功能概述：**

- 提供用户相关的接口（增删改查等）
- 使用 `gorm` 连接 MySQL
- 通过配置文件读取数据库连接信息，避免明文写在代码中

**配置方式：**

1. 进入 `user_srv` 目录：

   ```bash
   cd user_srv
   ```

2. 基于示例配置创建真实配置文件：

   ```bash
   cp config.example.json config.json
   ```

3. 编辑 `config.json`，把 `dsn` 改成你自己的数据库连接串，例如：

   ```json
   {
     "dsn": "root:your_password@tcp(127.0.0.1:3306)/your_db?charset=utf8mb4&parseTime=True&loc=Local"
   }
   ```

4. 建议在根目录或 `user_srv` 下的 `.gitignore` 中加入：

   ```gitignore
   user_srv/config.json
   ```

**运行 user_srv：**

```bash
cd user_srv
go mod tidy
go run main.go
```

如果配置正确，服务会启动并成功连接到 MySQL。

---

### user_web：Web / 网关服务（如果使用）

根据你的实现，通常用来：

- 对外提供 HTTP 接口
- 作为前端或其他服务访问 `user_srv` 的入口

**运行 user_web：**

```bash
cd user_web
go mod tidy
go run main.go
```

如有额外的配置（例如端口号、上游地址等），建议同样使用配置文件/环境变量管理，并在后续补充到本 README。

---

### 开发建议

- **不要在代码中写明文密码**：统一通过配置文件（如 `config.json`）或环境变量管理。
- **示例配置文件可提交**：例如 `config.example.json`，方便其他开发者参考字段格式。
- **真实配置文件忽略提交**：确保 `.gitignore` 中包含实际配置文件路径。

