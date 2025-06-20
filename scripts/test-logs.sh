#!/bin/bash

# Pelago Card Backend - Log Testing Script

echo "=== Testing Pelago Card Backend Logging ==="
echo ""

BASE_URL="http://localhost:8888"

echo "1. Testing successful merchant registration (INFO logs)..."
curl -s -X POST $BASE_URL/api/merchant/register \
  -H "Content-Type: application/json" \
  -d '{
    "entity_name": "Logging Test Corp",
    "brand_name": "LogTest",
    "website_url": "https://logtest.example.com",
    "merchant_country": "Singapore",
    "contact_name": "Log Tester",
    "contact_email": "logtester@example.com",
    "business_description": "Testing logging functionality"
  }' | jq '.'

echo ""
echo "2. Testing invalid merchant retrieval (ERROR logs)..."
curl -s -X GET $BASE_URL/api/merchant/INVALID_ID_FOR_TESTING 2>/dev/null || echo "Expected error response"

echo ""
echo "3. Testing merchant list with valid parameters (INFO logs)..."
curl -s -X GET "$BASE_URL/api/merchant/list?page=1&page_size=5" | jq '.total'

echo ""
echo "4. Testing merchant update with invalid ID (ERROR logs)..."
curl -s -X PUT $BASE_URL/api/merchant/update \
  -H "Content-Type: application/json" \
  -d '{
    "merchant_id": "INVALID_UPDATE_ID",
    "entity_name": "Updated Corp",
    "brand_name": "Updated",
    "merchant_country": "Singapore",
    "contact_name": "Updated Name",
    "contact_email": "updated@example.com"
  }' 2>/dev/null || echo "Expected error response"

echo ""
echo "=== Log testing completed. Check service logs for detailed output. ==="
echo "- RPC Service logs: Check terminal running card-rpc"
echo "- API Service logs: Check terminal running card-api"
echo "" 