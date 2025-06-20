-- Pelago Card 数据库初始化脚本
-- 版本: v1.0.0
-- 创建时间: 2025-06-20
-- 功能范围: 商户入驻

-- 设置字符集
SET NAMES utf8mb4;

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS pelago_card DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE pelago_card;

-- 商户表 (merchants)
-- 存储商户入驻信息
CREATE TABLE merchants (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    merchant_id VARCHAR(50) UNIQUE NOT NULL COMMENT '商户唯一ID，格式：M + yyyyMMddHHmm + 3位随机数',
    entity_name VARCHAR(200) NOT NULL COMMENT '商户主体名称，公司名称',
    brand_name VARCHAR(100) NOT NULL COMMENT '商户品牌名称',
    website_url VARCHAR(500) COMMENT '商户官方网站',
    merchant_logo VARCHAR(500) COMMENT '商户官方logo URL',
    merchant_country VARCHAR(50) NOT NULL COMMENT '商户经营国家',
    contact_name VARCHAR(100) NOT NULL COMMENT '商户主要联系人姓名',
    contact_email VARCHAR(100) NOT NULL COMMENT '商户主要联系人邮箱',
    api_key VARCHAR(100) NOT NULL COMMENT '商户API密钥，系统自动生成',
    verified_status VARCHAR(50) DEFAULT 'Pending for approval' COMMENT '审核状态：Pending for approval, Approved, Rejected',
    merchant_status VARCHAR(50) DEFAULT 'Inactive' COMMENT '商户状态：Active, Inactive',
    business_description TEXT COMMENT '商户业务描述',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    INDEX idx_merchant_id (merchant_id),
    INDEX idx_entity_name (entity_name),
    INDEX idx_brand_name (brand_name),
    INDEX idx_verified_status (verified_status),
    INDEX idx_merchant_status (merchant_status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商户表';

-- 插入测试数据（可选）
INSERT INTO merchants (
    merchant_id, entity_name, brand_name, website_url, merchant_logo, 
    merchant_country, contact_name, contact_email, api_key, 
    verified_status, merchant_status, business_description
) VALUES 
(
    'M202506200001001', 
    'Demo Company Limited', 
    'Demo Brand', 
    'https://demo.com', 
    'https://demo.com/logo.png',
    'Singapore', 
    'John Demo', 
    'john@demo.com', 
    'demo_api_key_12345678901234567890',
    'Approved',
    'Active',
    'Demo merchant for testing purposes'
);

-- 验证表创建
SELECT 'Merchants table created successfully' as status;
SELECT COUNT(*) as merchant_count FROM merchants; 