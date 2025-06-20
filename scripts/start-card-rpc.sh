#!/bin/bash

echo "🚀 启动商户RPC服务"
echo "=================="

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ Go环境未安装，请先安装Go"
    exit 1
fi

# 检查基础设施服务是否运行
if ! docker-compose ps mysql | grep -q "Up"; then
    echo "❌ MySQL服务未运行，请先运行: ./scripts/start-infra.sh"
    exit 1
fi

# 构建RPC服务
echo "🔨 构建商户RPC服务..."
go build -o bin/card-rpc rpc/card/card.go

if [ ! -f "bin/card-rpc" ]; then
    echo "❌ RPC服务构建失败"
    exit 1
fi

echo "✅ RPC服务构建完成"

# 启动RPC服务
echo "🚀 启动商户RPC服务..."
echo "📊 服务地址: localhost:8080"
echo "📝 使用 Ctrl+C 停止服务"
echo ""

./bin/card-rpc -f rpc/card/etc/card.yaml 