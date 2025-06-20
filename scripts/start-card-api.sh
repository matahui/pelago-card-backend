#!/bin/bash

echo "🚀 启动商户API服务"
echo "=================="

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ Go环境未安装，请先安装Go"
    exit 1
fi

# 检查RPC服务是否运行
if ! nc -z localhost 8080 2>/dev/null; then
    echo "❌ 商户RPC服务未运行，请先运行: ./scripts/start-card-rpc.sh"
    exit 1
fi

# 构建API服务
echo "🔨 构建商户API服务..."
go build -o bin/card-api api/card.go

if [ ! -f "bin/card-api" ]; then
    echo "❌ API服务构建失败"
    exit 1
fi

echo "✅ API服务构建完成"

# 启动API服务
echo "🚀 启动商户API服务..."
echo "📊 服务地址: http://localhost:8888"
echo "📝 使用 Ctrl+C 停止服务"
echo ""

./bin/card-api -f api/etc/card-api.yaml 