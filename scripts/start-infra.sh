#!/bin/bash

echo "🏗️  启动基础设施服务..."

# 检查Docker是否在运行
if ! docker info >/dev/null 2>&1; then
    echo "❌ Docker未运行，请先启动Docker"
    exit 1
fi

# 启动基础设施服务
echo "📦 启动基础设施服务..."
docker-compose up -d mysql redis etcd jaeger

# 等待服务启动
echo "⏳ 等待服务启动..."
sleep 10

# 显示服务状态
echo "📋 检查服务状态..."
docker-compose ps

echo ""
echo "🎉 基础设施服务启动完成！"
echo ""
echo "📊 服务地址："
echo "  MySQL:      127.0.0.1:3306"
echo "  Redis:      127.0.0.1:6379"
echo "  Etcd:       127.0.0.1:2379"
echo "  Jaeger UI:  http://127.0.0.1:16686"
echo ""
echo "📝 数据库连接信息："
echo "  Host: localhost:3306"
echo "  Database: pelago_card"
echo "  Username: root"
echo "  Password: password"
echo ""
echo "🚀 现在可以启动业务服务了！" 