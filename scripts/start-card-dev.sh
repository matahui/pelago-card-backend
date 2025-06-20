#!/bin/bash

echo "🚀 启动 Pelago Card 商户服务开发环境"
echo "======================================"

# 检查Docker是否在运行
if ! docker info >/dev/null 2>&1; then
    echo "❌ Docker未运行，请先启动Docker"
    exit 1
fi

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ Go环境未安装，请先安装Go"
    exit 1
fi

# 启动基础设施服务
echo "📦 启动基础设施服务..."
docker-compose up -d mysql redis etcd jaeger

# 等待服务启动
echo "⏳ 等待基础设施服务启动..."
sleep 15

# 检查基础设施服务状态
echo "📋 检查基础设施服务状态..."
docker-compose ps

# 构建服务
echo "🔨 构建商户服务..."
go build -o bin/card-rpc rpc/card/card.go
go build -o bin/card-api api/card.go

if [ ! -f "bin/card-rpc" ] || [ ! -f "bin/card-api" ]; then
    echo "❌ 服务构建失败"
    exit 1
fi

echo "✅ 服务构建完成"

# 检测操作系统
OS=$(uname -s)

# 启动RPC服务
echo "🚀 启动商户RPC服务..."
if [[ "$OS" == "Darwin" ]]; then
    # macOS
    osascript -e 'tell app "Terminal" to do script "cd \"'$(pwd)'\" && ./bin/card-rpc -f rpc/card/etc/card.yaml"'
elif [[ "$OS" == "Linux" ]]; then
    # Linux
    if command -v gnome-terminal &> /dev/null; then
        gnome-terminal -- bash -c "cd $(pwd) && ./bin/card-rpc -f rpc/card/etc/card.yaml; exec bash"
    elif command -v xterm &> /dev/null; then
        xterm -e "cd $(pwd) && ./bin/card-rpc -f rpc/card/etc/card.yaml; bash" &
    else
        echo "⚠️  请在新终端窗口中运行: ./bin/card-rpc -f rpc/card/etc/card.yaml"
    fi
else
    echo "⚠️  请在新终端窗口中运行: ./bin/card-rpc -f rpc/card/etc/card.yaml"
fi

# 等待RPC服务启动
echo "⏳ 等待RPC服务启动..."
sleep 8

# 启动API服务
echo "🚀 启动商户API服务..."
if [[ "$OS" == "Darwin" ]]; then
    # macOS
    osascript -e 'tell app "Terminal" to do script "cd \"'$(pwd)'\" && ./bin/card-api -f api/etc/card-api.yaml"'
elif [[ "$OS" == "Linux" ]]; then
    # Linux
    if command -v gnome-terminal &> /dev/null; then
        gnome-terminal -- bash -c "cd $(pwd) && ./bin/card-api -f api/etc/card-api.yaml; exec bash"
    elif command -v xterm &> /dev/null; then
        xterm -e "cd $(pwd) && ./bin/card-api -f api/etc/card-api.yaml; bash" &
    else
        echo "⚠️  请在新终端窗口中运行: ./bin/card-api -f api/etc/card-api.yaml"
    fi
else
    echo "⚠️  请在新终端窗口中运行: ./bin/card-api -f api/etc/card-api.yaml"
fi

echo ""
echo "🎉 商户服务开发环境启动完成！"
echo ""
echo "📊 服务地址："
echo "  商户API服务:    http://localhost:8888"
echo "  商户RPC服务:    localhost:8080"
echo "  MySQL数据库:    localhost:3306 (pelago_card)"
echo "  Redis缓存:      localhost:6379"
echo "  Etcd注册中心:   localhost:2379"
echo "  Jaeger链路追踪: http://localhost:16686"
echo ""
echo "🔧 测试命令："
echo "  ./scripts/test-card-api.sh    # 测试商户API"
echo "  ./scripts/test-logs.sh        # 测试日志系统"
echo ""
echo "📝 注意事项："
echo "  - 确保所有服务在新终端窗口中正常启动"
echo "  - 如果服务启动失败，请检查端口是否被占用"
echo "  - 使用 Ctrl+C 停止各个服务" 