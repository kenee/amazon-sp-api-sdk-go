# Amazon Selling Partner API SDK for Go

一个完整的 Amazon Selling Partner API (SP-API) Go SDK，提供类型安全、易于使用的 API 客户端。

## 特性

- ✅ **完整的 API 支持**: 涵盖所有主要 SP-API 端点
- ✅ **类型安全**: 强类型定义，编译时错误检查
- ✅ **错误处理**: 详细的错误信息和上下文
- ✅ **认证管理**: 自动处理 LWA 认证
- ✅ **易于使用**: 简洁的 API 设计
- ✅ **测试覆盖**: 完整的测试用例

## 支持的 API

### 核心 API
- **Listings API** - 产品刊登管理
- **Product Types API** - 产品类型定义
- **Catalog API** - 产品目录
- **Orders API** - 订单管理

### 扩展 API
- **Feeds API** - 批量数据处理
- **Inventory API** - 库存管理
- **Reports API** - 报告生成
- **Notifications API** - 通知管理

## 快速开始

### 安装

```bash
go get github.com/kenee/amazon-sp-api-sdk-go
```

### 环境配置

设置以下环境变量：

```bash
export SP_API_CLIENT_ID="your_client_id"
export SP_API_CLIENT_SECRET="your_client_secret"
export SP_API_REFRESH_TOKEN="your_refresh_token"
export SP_API_ENDPOINT="https://sellingpartnerapi-na.amazon.com"
export SP_API_ENDPOINT_HOST="sellingpartnerapi-na.amazon.com"
export SP_API_SELLER_ID="your_seller_id"
```

### 基本使用

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/kenee/amazon-sp-api-sdk-go/apis/listings"
    "github.com/kenee/amazon-sp-api-sdk-go/apis/producttypes"
    "github.com/kenee/amazon-sp-api-sdk-go/auth"
    "github.com/kenee/amazon-sp-api-sdk-go/client"
)

func main() {
    // 配置认证
    credentials := &auth.LWAAuthorizationCredentials{
        ClientID:     os.Getenv("SP_API_CLIENT_ID"),
        ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
        RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
        Endpoint:     os.Getenv("SP_API_ENDPOINT"),
    }

    config := client.NewConfigurationWithCredentials(credentials)
    config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

    ctx := context.Background()

    // 使用 Product Types API
    productTypesAPI := producttypes.NewProductTypesAPI(config)
    
    request := &producttypes.GetDefinitionsProductTypeRequest{
        ProductType:    "LUGGAGE",
        MarketplaceIds: []string{"ATVPDKIKX0DER"},
        SellerId:       os.Getenv("SP_API_SELLER_ID"),
    }

    response, err := productTypesAPI.GetDefinitionsProductType(ctx, request)
    if err != nil {
        log.Fatalf("获取产品类型定义失败: %v", err)
    }

    fmt.Printf("产品类型定义获取成功: %s\n", response.ProductType)
}
```

## API 使用示例

### Listings API

```go
// 创建刊登
listingsAPI := listings.NewListingsAPI(config)

request := &listings.PutListingsItemRequest{
    ProductType:    "LUGGAGE",
    Requirements:   "LISTING",
    Attributes:     map[string]interface{}{
        "brand": "Test Brand",
        "title": "Test Product",
    },
}

response, err := listingsAPI.PutListingsItem(ctx, "A13BCILWM2JF8S", "GM-ZDPI-9B4E", request)
```

### Catalog API

```go
// 获取产品信息
catalogAPI := catalog.NewCatalogAPI(config)

request := &catalog.GetCatalogItemRequest{
    ASIN:           "B08N5WRWNW",
    MarketplaceIds: []string{"ATVPDKIKX0DER"},
}

response, err := catalogAPI.GetCatalogItem(ctx, request)
```

### Feeds API

```go
// 创建 Feed
feedsAPI := feeds.NewFeedsAPI(config)

request := &feeds.CreateFeedRequest{
    FeedType:           "POST_PRODUCT_DATA",
    MarketplaceIds:     []string{"ATVPDKIKX0DER"},
    InputFeedDocumentId: "feed_document_id",
}

response, err := feedsAPI.CreateFeed(ctx, request)
```

### Inventory API

```go
// 获取库存摘要
inventoryAPI := inventory.NewInventoryAPI(config)

request := &inventory.GetInventorySummariesRequest{
    SellerId:       "A13BCILWM2JF8S",
    MarketplaceIds: []string{"ATVPDKIKX0DER"},
    Details:        true,
}

response, err := inventoryAPI.GetInventorySummaries(ctx, request)
```

### Reports API

```go
// 创建报告
reportsAPI := reports.NewReportsAPI(config)

request := &reports.CreateReportRequest{
    ReportType:     "GET_FLAT_FILE_OPEN_LISTINGS_DATA",
    MarketplaceIds: []string{"ATVPDKIKX0DER"},
}

response, err := reportsAPI.CreateReport(ctx, request)
```

## 错误处理

SDK 提供了详细的错误处理机制：

```go
if err != nil {
    switch e := err.(type) {
    case *client.AuthenticationError:
        log.Printf("认证错误: %s", e.Error())
    case *client.ValidationError:
        log.Printf("验证错误: %s", e.Error())
    case *client.RateLimitExceededError:
        log.Printf("频率限制: %s, 重试时间: %d秒", e.Error(), e.RetryAfter)
    case *client.APIError:
        log.Printf("API错误: %s (状态码: %d)", e.Error(), e.StatusCode)
    default:
        log.Printf("其他错误: %s", err.Error())
    }
}
```

## 测试

运行测试：

```bash
# 运行所有测试
go test -mod=mod ./tests/... -v

# 运行特定测试
go test -mod=mod ./tests/listing_flow/... -run TestStep1_GetProductTypeDefinition -v

# 运行错误处理测试
go test -mod=mod ./tests/error_handling_test.go -v
```

## 示例代码

查看 `examples/` 目录中的完整示例：

- `feeds_example.go` - Feeds API 使用示例
- `inventory_example.go` - Inventory API 使用示例
- `listings_example.go` - Listings API 使用示例

## 贡献

欢迎提交 Issue 和 Pull Request！

### 开发环境设置

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 运行测试
5. 提交 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 支持

- 查看 [Amazon SP-API 文档](https://developer-docs.amazon.com/sp-api/)
- 提交 [Issue](https://github.com/kenee/amazon-sp-api-sdk-go/issues)
- 查看 [示例代码](examples/)

## 更新日志

### v1.0.0
- 初始版本发布
- 支持核心 API (Listings, Product Types, Catalog, Orders)
- 完整的错误处理机制
- 类型安全的 API 设计

### v1.1.0
- 添加 Feeds API 支持
- 添加 Inventory API 支持
- 添加 Reports API 支持
- 改进错误处理和文档
