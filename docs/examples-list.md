# 亚马逊 SP-API PHP 示例文件列表

## 目录结构

```
php/examples/
├── getOrders.php                           # Orders API - 获取订单列表
├── getDefinitionsProductType.php           # Product Type Definitions API - 获取产品类型定义
├── searchDefinitionsProductTypes.php       # Product Type Definitions API - 搜索产品类型定义
├── searchDefinitionsProductTypesWithKeywords.php  # Product Type Definitions API - 带关键词搜索
├── searchCatalogItems.php                  # Catalog Items API - 搜索商品目录
├── searchCatalogItemsByASIN.php           # Catalog Items API - 按 ASIN 搜索商品
├── getCatalogItem.php                      # Catalog Items API - 获取单个商品详情
├── getCatalogItemSimple.php               # Catalog Items API - 简化版获取商品详情
├── getListingsItem.php                     # Listings Items API - 获取单个商品列表详情
├── getListingsItemSimple.php              # Listings Items API - 简化版获取商品列表详情
├── putListingsItem.php                     # Listings Items API - 创建/更新商品列表
├── putListingsItemCreate.php              # Listings Items API - 实际创建商品
├── patchListingsItem.php                   # Listings Items API - 部分更新商品列表
├── patchListingsItemSimple.php            # Listings Items API - 简化版部分更新
├── searchListingsItems.php                 # Listings Items API - 搜索商品列表
├── searchListingsItemsSimple.php          # Listings Items API - 简化版搜索
├── searchListingsItemsByASIN.php          # Listings Items API - 按 ASIN 搜索
├── getListingsRestrictions.php             # Listings Restrictions API - 获取商品限制
├── getListingsRestrictionsSimple.php      # Listings Restrictions API - 简化版获取限制
├── listCatalogCategories.php              # Catalog Categories API - 获取商品分类 (原生 HTTP)
└── listCatalogCategoriesBySKU.php        # Catalog Categories API - 按 SKU 获取分类 (原生 HTTP)
```

## 按 API 分类

### Orders API
- **getOrders.php** - 获取订单列表
  - 沙盒参数：`CreatedAfter: 'TEST_CASE_200'`
  - 支持分页：`'TEST_CASE_200_NEXT_TOKEN'`

### Product Type Definitions API
- **getDefinitionsProductType.php** - 获取产品类型定义
  - 沙盒参数：`productType: 'LUGGAGE'`
- **searchDefinitionsProductTypes.php** - 搜索产品类型定义
  - 基本搜索功能
- **searchDefinitionsProductTypesWithKeywords.php** - 带关键词搜索
  - 沙盒参数：`keywords: ['luggage', 'travel']`

### Catalog Items API
- **searchCatalogItems.php** - 搜索商品目录
  - 支持多种标识符类型
- **searchCatalogItemsByASIN.php** - 按 ASIN 搜索商品
  - 沙盒参数：`identifiers: ['B07N4M94X4', 'B08J7TQ9FL']`
- **getCatalogItem.php** - 获取单个商品详情
  - 完整数据集版本
- **getCatalogItemSimple.php** - 简化版获取商品详情
  - 只包含基本摘要信息

### Listings Items API
- **getListingsItem.php** - 获取单个商品列表详情
  - 完整数据集版本
- **getListingsItemSimple.php** - 简化版获取商品列表详情
  - 只包含基本摘要信息
- **putListingsItem.php** - 创建/更新商品列表
  - 验证预览模式
- **putListingsItemCreate.php** - 实际创建商品
  - 实际创建模式
- **patchListingsItem.php** - 部分更新商品列表
  - 多个补丁操作示例
- **patchListingsItemSimple.php** - 简化版部分更新
  - 单个补丁操作示例
- **searchListingsItems.php** - 搜索商品列表
  - 完整搜索功能
- **searchListingsItemsSimple.php** - 简化版搜索
  - 基本搜索功能
- **searchListingsItemsByASIN.php** - 按 ASIN 搜索
  - 使用 ASIN 标识符搜索

### Listings Restrictions API
- **getListingsRestrictions.php** - 获取商品限制
  - 包含条件类型过滤
- **getListingsRestrictionsSimple.php** - 简化版获取限制
  - 基本限制查询

### Catalog Categories API
- **listCatalogCategories.php** - 获取商品分类
  - 使用 ASIN 参数
- **listCatalogCategoriesBySKU.php** - 按 SKU 获取分类
  - 使用 SKU 参数

## 运行示例

### 1. 环境准备
```bash
# 安装依赖
cd php/sdk
composer install

# 创建环境变量文件
cd ../..
echo "SP_API_CLIENT_ID=你的clientId
SP_API_CLIENT_SECRET=你的clientSecret
SP_API_REFRESH_TOKEN=你的refreshToken
SP_API_ENDPOINT=https://api.amazon.com/auth/o2/token
SP_API_ENDPOINT_HOST=https://sandbox.sellingpartnerapi-na.amazon.com
SP_API_MARKETPLACE=ATVPDKIKX0DER" > .env
```

### 2. 运行示例
```bash
cd php/examples

# 运行基本示例
php getOrders.php
php getDefinitionsProductType.php
php searchCatalogItems.php

# 运行简化版示例
php getCatalogItemSimple.php
php getListingsItemSimple.php
php patchListingsItemSimple.php
php searchListingsItemsSimple.php
```

## 示例特点

### 简化版示例
- 使用最少的必需参数
- 适合快速测试和了解接口
- 包含基本错误处理

### 完整版示例
- 包含所有可选参数
- 展示高级功能
- 包含详细的错误处理和注释

### 原生 HTTP 示例
- 用于 SDK 中不存在的接口
- 直接使用 Guzzle HTTP 客户端
- 手动处理认证和签名

## 注意事项

1. **环境变量**：所有示例都需要正确的 `.env` 文件
2. **沙盒环境**：示例使用沙盒测试值，生产环境需要真实数据
3. **错误处理**：所有示例都包含基本的错误处理
4. **速率限制**：注意不同接口的速率限制
5. **参数验证**：确保参数组合正确，避免互斥参数冲突

## 相关文档

- [详细接口文档](./amazon-sp-api-php-examples.md)
- [快速参考指南](./quick-reference.md)
- [官方 SP-API 文档](https://developer-docs.amazon.com/sp-api/) 