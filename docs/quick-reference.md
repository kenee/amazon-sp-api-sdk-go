# 亚马逊 SP-API PHP 快速参考指南

## 环境变量

```env
SP_API_CLIENT_ID=你的clientId
SP_API_CLIENT_SECRET=你的clientSecret
SP_API_REFRESH_TOKEN=你的refreshToken
SP_API_ENDPOINT=https://api.amazon.com/auth/o2/token
SP_API_ENDPOINT_HOST=https://sandbox.sellingpartnerapi-na.amazon.com
SP_API_MARKETPLACE=ATVPDKIKX0DER
```

## 接口快速参考

### Orders API
| 接口 | 文件 | 关键参数 | 注意事项 |
|------|------|----------|----------|
| getOrders | `getOrders.php` | `CreatedAfter: 'TEST_CASE_200'` | 沙盒使用测试值 |

### Product Type Definitions API
| 接口 | 文件 | 关键参数 | 注意事项 |
|------|------|----------|----------|
| getDefinitionsProductType | `getDefinitionsProductType.php` | `productType: 'LUGGAGE'` | 沙盒支持的产品类型 |
| searchDefinitionsProductTypes | `searchDefinitionsProductTypes.php` | `keywords: ['luggage']` | 与 itemName 互斥 |

### Catalog Items API
| 接口 | 文件 | 关键参数 | 注意事项 |
|------|------|----------|----------|
| searchCatalogItems | `searchCatalogItems.php` | `identifiers: ['B07N4M94X4']` | 与 keywords 互斥 |
| getCatalogItem | `getCatalogItem.php` | `asin: 'B07N4M94X4'` | 沙盒测试 ASIN |

### Listings Items API
| 接口 | 文件 | 关键参数 | 注意事项 |
|------|------|----------|----------|
| getListingsItem | `getListingsItem.php` | `sku: 'GM-ZDPI-9B4E'` | 获取单个商品详情 |
| putListingsItem | `putListingsItem.php` | `productType: 'LUGGAGE'` | 需要 marketplace_id 和 language_tag |
| patchListingsItem | `patchListingsItem.php` | `op: 'replace'` | 只支持顶级属性 |
| searchListingsItems | `searchListingsItems.php` | `identifiers: ['GM-ZDPI-9B4E']` | 最多 20 个标识符 |

### Listings Restrictions API
| 接口 | 文件 | 关键参数 | 注意事项 |
|------|------|----------|----------|
| getListingsRestrictions | `getListingsRestrictions.php` | `asin: 'B07N4M94X4'` | 可过滤条件类型 |

### Catalog Categories API
| 接口 | 文件 | 关键参数 | 注意事项 |
|------|------|----------|----------|
| listCatalogCategories | `listCatalogCategories.php` | `ASIN: 'B07N4M94X4'` | SDK 中不存在，使用原生请求 |

## 沙盒测试值

### Marketplace IDs
- 美国沙盒：`ATVPDKIKX0DER`

### 测试 ASINs
- `B07N4M94X4`
- `B08J7TQ9FL`

### 测试 SKUs
- `GM-ZDPI-9B4E`
- `HW-ZDPI-9B4E`
- `TC-ZDPI-9B4E`
- `LUGGAGE-SKU-001`

### 测试 Seller ID
- `A1B2C3D4E5F6G7`

### 产品类型
- `LUGGAGE`
- `INVALID` (用于测试错误)

## 常见错误和解决方案

### 1. LWA Token Exchange 失败
```
Exception when calling OrderApi->getOrders: Unsuccessful LWA token exchange
```
**解决方案**：
- 检查 `.env` 文件是否存在
- 验证环境变量是否正确
- 确认 refresh token 是否有效

### 2. 参数验证错误
```
Invalid value for "marketplaceIds" when calling ListingsApi.searchListingsItems
```
**解决方案**：
- 确保 marketplaceIds 数组只包含一个元素
- 检查参数类型和格式

### 3. 互斥参数错误
```
Cannot use 'identifiers' if you specify 'variationParentSku'
```
**解决方案**：
- 检查参数组合是否冲突
- 参考接口文档的参数互斥规则

## 速率限制

| 接口 | 速率 (req/s) | 突发 | 说明 |
|------|-------------|------|------|
| getOrders | 0.0167 | 15 | 每分钟 1 次 |
| getDefinitionsProductType | 5 | 10 | 产品类型定义 |
| searchCatalogItems | 5 | 10 | 商品搜索 |
| getListingsItem | 5 | 10 | 获取商品详情 |
| putListingsItem | 5 | 5 | 创建商品 |
| patchListingsItem | 5 | 5 | 更新商品 |
| searchListingsItems | 5 | 5 | 搜索商品列表 |

## 数据格式要求

### 时间格式
```php
new DateTime('2023-01-01T00:00:00Z') // ISO 8601
```

### 金额格式
```php
'100.00' // 字符串格式
```

### 数组参数
```php
['ATVPDKIKX0DER'] // 逗号分隔
```

## 分页处理

```php
// 检查是否有下一页
if ($response->getPagination() && $response->getPagination()->getNextToken()) {
    $nextToken = $response->getPagination()->getNextToken();
    // 使用 nextToken 获取下一页
}
```

## 错误处理最佳实践

```php
try {
    $response = $api->someMethod($params);
    // 处理成功响应
} catch (ApiException $e) {
    // 处理 API 错误
    echo 'API Error: ' . $e->getMessage();
} catch (Exception $e) {
    // 处理其他错误
    echo 'General Error: ' . $e->getMessage();
}
```

## 常用工具函数

### 检查环境变量
```php
function checkEnvVars() {
    $required = ['SP_API_CLIENT_ID', 'SP_API_CLIENT_SECRET', 'SP_API_REFRESH_TOKEN'];
    foreach ($required as $var) {
        if (empty($_ENV[$var])) {
            throw new Exception("Missing required environment variable: $var");
        }
    }
}
```

### 格式化响应
```php
function formatResponse($response) {
    return json_encode($response, JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);
}
``` 