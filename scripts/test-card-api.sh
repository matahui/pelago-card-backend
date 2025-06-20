#!/bin/bash

echo "🧪 测试商户API接口"
echo "=================="

# 检查API服务是否运行
if ! nc -z localhost 8888 2>/dev/null; then
    echo "❌ 商户API服务未运行，请先启动服务"
    echo "   运行: ./scripts/start-card-api.sh"
    exit 1
fi

echo "✅ API服务运行正常，开始测试..."
echo ""

# 测试变量
API_BASE="http://localhost:8888/api"
MERCHANT_ID=""

# 1. 测试商户注册
echo "1️⃣  测试商户注册 API"
echo "POST ${API_BASE}/merchant/register"
echo "-----------------------------------"

REGISTER_RESPONSE=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X POST "${API_BASE}/merchant/register" \
  -H "Content-Type: application/json" \
  -d '{
    "entity_name": "Test Company Ltd",
    "brand_name": "TestBrand",
    "website_url": "https://test.example.com", 
    "merchant_country": "Singapore",
    "contact_name": "Test User",
    "contact_email": "test@example.com",
    "business_description": "API测试商户"
  }')

HTTP_CODE=$(echo "$REGISTER_RESPONSE" | grep "HTTP_CODE:" | cut -d: -f2)
RESPONSE_BODY=$(echo "$REGISTER_RESPONSE" | sed '/HTTP_CODE:/d')

echo "响应状态: $HTTP_CODE"
echo "响应内容: $RESPONSE_BODY"

if [ "$HTTP_CODE" = "200" ]; then
    echo "✅ 商户注册成功"
    # 提取merchant_id用于后续测试
    MERCHANT_ID=$(echo "$RESPONSE_BODY" | grep -o '"merchant_id":"[^"]*"' | cut -d'"' -f4)
    echo "🆔 商户ID: $MERCHANT_ID"
else
    echo "❌ 商户注册失败"
fi

echo ""

# 2. 测试获取商户信息
if [ ! -z "$MERCHANT_ID" ]; then
    echo "2️⃣  测试获取商户信息 API"
    echo "GET ${API_BASE}/merchant/${MERCHANT_ID}"
    echo "-----------------------------------"
    
    GET_RESPONSE=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X GET "${API_BASE}/merchant/${MERCHANT_ID}")
    
    HTTP_CODE=$(echo "$GET_RESPONSE" | grep "HTTP_CODE:" | cut -d: -f2)
    RESPONSE_BODY=$(echo "$GET_RESPONSE" | sed '/HTTP_CODE:/d')
    
    echo "响应状态: $HTTP_CODE"
    echo "响应内容: $RESPONSE_BODY"
    
    if [ "$HTTP_CODE" = "200" ]; then
        echo "✅ 获取商户信息成功"
    else
        echo "❌ 获取商户信息失败"
    fi
else
    echo "2️⃣  跳过获取商户信息测试（没有有效的merchant_id）"
fi

echo ""

# 3. 测试商户列表
echo "3️⃣  测试商户列表 API"
echo "GET ${API_BASE}/merchant/list?page=1&page_size=5"
echo "-----------------------------------"

LIST_RESPONSE=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X GET "${API_BASE}/merchant/list?page=1&page_size=5")

HTTP_CODE=$(echo "$LIST_RESPONSE" | grep "HTTP_CODE:" | cut -d: -f2)
RESPONSE_BODY=$(echo "$LIST_RESPONSE" | sed '/HTTP_CODE:/d')

echo "响应状态: $HTTP_CODE"
echo "响应内容: $RESPONSE_BODY"

if [ "$HTTP_CODE" = "200" ]; then
    echo "✅ 获取商户列表成功"
else
    echo "❌ 获取商户列表失败"
fi

echo ""

# 4. 测试更新商户信息
if [ ! -z "$MERCHANT_ID" ]; then
    echo "4️⃣  测试更新商户信息 API"
    echo "PUT ${API_BASE}/merchant/update"
    echo "-----------------------------------"
    
    UPDATE_RESPONSE=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X PUT "${API_BASE}/merchant/update" \
      -H "Content-Type: application/json" \
      -d "{
        \"merchant_id\": \"${MERCHANT_ID}\",
        \"entity_name\": \"Updated Test Company Ltd\",
        \"brand_name\": \"UpdatedBrand\",
        \"website_url\": \"https://updated.example.com\",
        \"merchant_country\": \"Malaysia\",
        \"contact_name\": \"Updated Test User\",
        \"contact_email\": \"updated@example.com\",
        \"business_description\": \"更新后的API测试商户\"
      }")
    
    HTTP_CODE=$(echo "$UPDATE_RESPONSE" | grep "HTTP_CODE:" | cut -d: -f2)
    RESPONSE_BODY=$(echo "$UPDATE_RESPONSE" | sed '/HTTP_CODE:/d')
    
    echo "响应状态: $HTTP_CODE"
    echo "响应内容: $RESPONSE_BODY"
    
    if [ "$HTTP_CODE" = "200" ]; then
        echo "✅ 更新商户信息成功"
    else
        echo "❌ 更新商户信息失败"
    fi
else
    echo "4️⃣  跳过更新商户信息测试（没有有效的merchant_id）"
fi

echo ""

# 5. 测试错误情况
echo "5️⃣  测试错误情况"
echo "GET ${API_BASE}/merchant/INVALID_ID"
echo "-----------------------------------"

ERROR_RESPONSE=$(curl -s -w "\nHTTP_CODE:%{http_code}" -X GET "${API_BASE}/merchant/INVALID_ID")

HTTP_CODE=$(echo "$ERROR_RESPONSE" | grep "HTTP_CODE:" | cut -d: -f2)
RESPONSE_BODY=$(echo "$ERROR_RESPONSE" | sed '/HTTP_CODE:/d')

echo "响应状态: $HTTP_CODE"
echo "响应内容: $RESPONSE_BODY"

if [ "$HTTP_CODE" = "500" ]; then
    echo "✅ 错误处理正常"
else
    echo "❌ 错误处理异常"
fi

echo ""
echo "🎉 API测试完成！"
echo ""
echo "📊 测试摘要："
echo "  - 商户注册: 已测试"
echo "  - 获取商户信息: 已测试"
echo "  - 商户列表: 已测试"
echo "  - 更新商户信息: 已测试"
echo "  - 错误处理: 已测试"
echo ""
echo "📝 注意事项："
echo "  - 测试数据已保存到数据库中"
echo "  - 可以通过Jaeger查看链路追踪: http://localhost:16686"
echo "  - 检查服务日志了解详细执行情况" 