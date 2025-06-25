# 亚马逊 SP-API PHP 示例文档

本文档包含了亚马逊 Selling Partner API (SP-API) 的 PHP 示例代码和重要注意事项。

## 目录

1. [环境配置](#环境配置)
2. [Orders API](#orders-api)
3. [Product Type Definitions API](#product-type-definitions-api)
4. [Catalog Items API](#catalog-items-api)
5. [Listings Items API](#listings-items-api)
6. [Listings Restrictions API](#listings-restrictions-api)
7. [Catalog Categories API](#catalog-categories-api)

---

## 环境配置

### 必需的环境变量

在项目根目录创建 `.env` 文件：

```env
SP_API_CLIENT_ID=你的clientId
SP_API_CLIENT_SECRET=你的clientSecret
SP_API_REFRESH_TOKEN=你的refreshToken
SP_API_ENDPOINT=https://api.amazon.com/auth/o2/token
SP_API_ENDPOINT_HOST=https://sandbox.sellingpartnerapi-na.amazon.com
SP_API_MARKETPLACE=ATVPDKIKX0DER
```

### 沙盒环境配置

- **LWA Token Endpoint**: `https://api.amazon.com/auth/o2/token` (sandbox 和 production 相同)
- **SP-API Host**: `https://sandbox.sellingpartnerapi-na.amazon.com`
- **Marketplace ID**: `ATVPDKIKX0DER` (美国沙盒市场)

---

## Orders API

### 接口：getOrders

**文件**: `php/examples/getOrders.php`

**功能**: 获取订单列表

**沙盒参数**:
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `CreatedAfter`: `'TEST_CASE_200'` (沙盒专用测试值)

**重要注意事项**:
- 沙盒环境中 `CreatedAfter` 参数可以使用特殊测试值而不是真实日期
- 支持的分页令牌测试值：`'TEST_CASE_200_NEXT_TOKEN'`

**示例代码**:
```php
$response = $ordersApi->getOrders(
    ['ATVPDKIKX0DER'], // MarketplaceIds
    'TEST_CASE_200' // CreatedAfter - 沙盒测试值
);
```

---

## Product Type Definitions API

### 接口：getDefinitionsProductType

**文件**: `php/examples/getDefinitionsProductType.php`

**功能**: 获取产品类型定义

**沙盒参数**:
- `productType`: `'LUGGAGE'`
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `requirements`: `'LISTING'`
- `requirementsEnforced`: `'ENFORCED'`
- `locale`: `'en_US'`

**重要注意事项**:
- 只支持顶级属性的补丁操作，不支持嵌套属性
- 沙盒环境支持的产品类型：`'LUGGAGE'`, `'INVALID'` (用于测试错误)

### 接口：searchDefinitionsProductTypes

**文件**: 
- `php/examples/searchDefinitionsProductTypes.php`
- `php/examples/searchDefinitionsProductTypesWithKeywords.php`

**功能**: 搜索产品类型定义

**沙盒参数**:
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `keywords`: `['luggage', 'travel']` (可选)
- `itemName`: `null` (与 keywords 互斥)
- `locale`: `'en_US'`

**重要注意事项**:
- `keywords` 和 `itemName` 参数不能同时使用
- 沙盒环境支持的关键词：`['luggage', 'travel']`

---

## Catalog Items API

### 接口：searchCatalogItems

**文件**: 
- `php/examples/searchCatalogItems.php`
- `php/examples/searchCatalogItemsByASIN.php`

**功能**: 搜索商品目录

**沙盒参数**:
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `identifiers`: `['B07N4M94X4', 'B08J7TQ9FL']` (可选)
- `identifiersType`: `'ASIN'` 或 `'UPC'` 或 `'EAN'`
- `keywords`: `['luggage', 'travel']` (与 identifiers 互斥)
- `includedData`: `['summaries', 'attributes', 'images', 'classifications']`

**重要注意事项**:
- `identifiers` 和 `keywords` 参数不能同时使用
- 沙盒环境支持的 ASIN：`['B07N4M94X4', 'B08J7TQ9FL']`

### 接口：getCatalogItem

**文件**: 
- `php/examples/getCatalogItem.php`
- `php/examples/getCatalogItemSimple.php`

**功能**: 获取单个商品详情

**沙盒参数**:
- `asin`: `'B07N4M94X4'`
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `includedData`: `['summaries']` 或完整数据集

**重要注意事项**:
- 沙盒环境支持的测试 ASIN：`'B07N4M94X4'`
- 完整数据集包括：`['summaries', 'attributes', 'images', 'classifications', 'dimensions', 'identifiers', 'productTypes', 'relationships', 'salesRanks', 'vendorDetails']`

---

## Listings Items API

### 接口：getListingsItem

**文件**: 
- `php/examples/getListingsItem.php`
- `php/examples/getListingsItemSimple.php`

**功能**: 获取单个商品列表详情

**沙盒参数**:
- `sellerId`: `'A1B2C3D4E5F6G7'`
- `sku`: `'GM-ZDPI-9B4E'`
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `issueLocale`: `'en_US'` (可选)
- `includedData`: `['summaries', 'offers', 'fulfillmentAvailability', 'issues']` (可选)

**重要注意事项**:
- 沙盒环境支持的测试 SKU：`'GM-ZDPI-9B4E'`
- 默认包含 `summaries` 数据集
- 可以包含多个数据集：`summaries`, `offers`, `fulfillmentAvailability`, `issues`

### 接口：putListingsItem

**文件**: 
- `php/examples/putListingsItem.php`
- `php/examples/putListingsItemCreate.php`

**功能**: 创建或更新商品列表

**沙盒参数**:
- `sellerId`: `'A1B2C3D4E5F6G7'`
- `sku`: `'LUGGAGE-SKU-001'`
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `productType`: `'LUGGAGE'`
- `requirements`: `'LISTING'`

**重要注意事项**:
- 每个属性都需要指定 `marketplace_id` 和 `language_tag`
- 沙盒环境支持的产品类型：`'LUGGAGE'`
- 模式选项：`'VALIDATION_PREVIEW'` (验证预览), `'CREATE'` (实际创建)

### 接口：patchListingsItem

**文件**: 
- `php/examples/patchListingsItem.php`
- `php/examples/patchListingsItemSimple.php`

**功能**: 部分更新商品列表

**沙盒参数**:
- `sellerId`: `'A1B2C3D4E5F6G7'`
- `sku`: `'LUGGAGE-SKU-001'`
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `productType`: `'LUGGAGE'`

**PatchOperation 参数**:
- `op`: `'add'`, `'replace'`, `'merge'`, `'delete'`
- `path`: JSON 指针路径，如 `/attributes/item_name`
- `value`: 属性值数组

**重要注意事项**:
- 只支持顶级属性的补丁操作，不支持嵌套属性
- 路径格式使用 JSON Pointer 格式
- 每个值必须包含 `marketplace_id`、`language_tag`、`value`

### 接口：searchListingsItems

**文件**: 
- `php/examples/searchListingsItems.php`
- `php/examples/searchListingsItemsSimple.php`
- `php/examples/searchListingsItemsByASIN.php`

**功能**: 搜索商品列表

**沙盒参数**:
- `sellerId`: `'A1B2C3D4E5F6G7'`
- `marketplaceIds`: `['ATVPDKIKX0DER']` (只能包含一个)
- `identifiers`: `['GM-ZDPI-9B4E', 'HW-ZDPI-9B4E']` 或 `['B07N4M94X4', 'B08J7TQ9FL']`
- `identifiersType`: `'SKU'` 或 `'ASIN'`
- `includedData`: `['summaries', 'offers', 'fulfillmentAvailability', 'issues']`

**重要注意事项**:
- 每次请求只能搜索一个 marketplace
- 最多 20 个标识符
- `identifiers`、`variationParentSku`、`packageHierarchySku` 不能同时使用
- 每页最多 20 个结果

---

## Listings Restrictions API

### 接口：getListingsRestrictions

**文件**: 
- `php/examples/getListingsRestrictions.php`
- `php/examples/getListingsRestrictionsSimple.php`

**功能**: 获取商品限制信息

**沙盒参数**:
- `asin`: `'B07N4M94X4'`
- `sellerId`: `'A1B2C3D4E5F6G7'`
- `marketplaceIds`: `['ATVPDKIKX0DER']`
- `conditionType`: `'new_new'` (可选)

**条件类型选项**:
- `'new_new'` - 全新
- `'new_open_box'` - 全新开箱
- `'new_oem'` - 全新 OEM
- `'refurbished_refurbished'` - 翻新
- `'used_like_new'` - 二手如新
- `'used_very_good'` - 二手很好
- `'used_good'` - 二手良好
- `'used_acceptable'` - 二手可接受

**重要注意事项**:
- 沙盒环境支持的测试 ASIN：`'B07N4M94X4'`
- 可以按商品状况过滤限制信息

---

## Catalog Categories API

### 接口：listCatalogCategories

**文件**: 
- `php/examples/listCatalogCategories.php`
- `php/examples/listCatalogCategoriesBySKU.php`

**功能**: 获取商品分类信息

**注意**: 此接口在当前 SDK 中不存在，使用原生 HTTP 请求实现

**沙盒参数**:
- `marketplaceIds`: `'ATVPDKIKX0DER'`
- `ASIN`: `'B07N4M94X4'` (与 SellerSKU 二选一)
- `SellerSKU`: `'GM-ZDPI-9B4E'` (与 ASIN 二选一)
- `sellerId`: `'A1B2C3D4E5F6G7'` (使用 SKU 时必需)

**重要注意事项**:
- 使用 SKU 时必须同时提供 `sellerId`
- ASIN 和 SellerSKU 参数不能同时使用
- 此接口属于 Catalog Items v0 API，当前 SDK 只包含 v2022-04-01 版本

---

## 通用注意事项

### 1. 认证和配置
- 确保 `.env` 文件在项目根目录
- 检查所有环境变量是否正确设置
- 沙盒环境使用专门的 endpoint

### 2. 错误处理
- 所有示例都包含 try-catch 错误处理
- 检查 LWA token exchange 是否成功
- 验证 API 响应状态码

### 3. 速率限制
- 不同接口有不同的速率限制
- 建议实现适当的重试机制
- 监控 `x-amzn-RateLimit-Limit` 响应头

### 4. 沙盒环境限制
- 沙盒环境返回模拟数据
- 某些操作在沙盒中可能不可用
- 测试值可能与生产环境不同

### 5. 数据格式
- 时间参数使用 ISO 8601 格式
- 金额使用字符串格式
- 数组参数使用逗号分隔

### 6. 分页处理
- 使用 `pageToken` 进行分页
- 检查 `pagination.nextToken` 是否存在
- 合理设置 `pageSize` 参数

---

## 运行示例

1. 确保已安装依赖：
```bash
cd php/sdk
composer install
```

2. 设置环境变量（创建 `.env` 文件）

3. 运行示例：
```bash
cd php/examples
php getOrders.php
```

---

## 相关资源

- [亚马逊 SP-API 官方文档](https://developer-docs.amazon.com/sp-api/)
- [SP-API 使用指南](https://developer-docs.amazon.com/sp-api/docs/sp-api-overview)
- [速率限制说明](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api)
- [沙盒环境说明](https://developer-docs.amazon.com/sp-api/docs/sp-api-sandbox) 