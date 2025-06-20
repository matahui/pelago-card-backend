# Scripts 使用说明

这个目录包含了项目的各种启动和管理脚本。

## 📁 脚本列表

### 🏗️ 基础设施脚本
- `start-infra.sh` - 启动基础设施服务（MySQL, Redis, Etcd, Jaeger）

### 💳 商户服务脚本
- `start-card-dev.sh` - 一键启动商户服务完整开发环境
- `start-card-rpc.sh` - 单独启动商户RPC服务
- `start-card-api.sh` - 单独启动商户API服务
- `test-card-api.sh` - 测试商户API接口
- `test-logs.sh` - 测试日志系统功能

## 🚀 使用方法

### 方式一：一键启动（推荐）
```bash
# 给脚本执行权限
chmod +x scripts/*.sh

# 一键启动商户服务环境
./scripts/start-card-dev.sh
```

### 方式二：分步启动
```bash
# 1. 先启动基础设施
./scripts/start-infra.sh

# 2. 启动商户RPC服务（新终端窗口）
./scripts/start-card-rpc.sh

# 3. 启动商户API服务（新终端窗口）
./scripts/start-card-api.sh
```

### 测试服务
```bash
# 测试商户API
./scripts/test-card-api.sh

# 测试日志系统
./scripts/test-logs.sh
```

## 📋 服务依赖关系

```
基础设施服务 (MySQL, Redis, Etcd, Jaeger)
    ↓
商户RPC服务 (端口: 8080)
    ↓
商户API服务 (端口: 8888)
```

## 🛠️ 脚本特性

- ✅ 自动检查Docker运行状态
- ✅ 自动检查服务依赖关系
- ✅ 跨平台终端启动支持（Linux/macOS）
- ✅ 详细的状态反馈和错误提示
- ✅ 一键测试API功能
- ✅ 完整的日志系统测试

## 📝 注意事项

1. 确保Docker已安装并运行
2. 确保Go环境已配置
3. RPC服务需要在API服务之前启动
4. 所有脚本都应该从项目根目录执行

## 🔧 开发工具

### 代码生成
```bash
# 生成RPC代码
goctl rpc protoc rpc/proto/card.proto --go_out=rpc/pb --go-grpc_out=rpc/pb --zrpc_out=rpc/card

# 生成API代码
goctl api go -api api/apis/card.api -dir api

# 生成数据库模型
goctl model mysql ddl -src sql/v1.0.0/init.sql -dir rpc/card/internal/model
```

### 服务地址
- **商户API服务**: http://localhost:8888
- **商户RPC服务**: localhost:8080
- **MySQL数据库**: localhost:3306 (pelago_card)
- **Redis缓存**: localhost:6379
- **Etcd注册中心**: localhost:2379
- **Jaeger链路追踪**: http://localhost:16686 