# 亚马逊 SP-API PHP 文档

欢迎使用亚马逊 Selling Partner API (SP-API) 的 PHP 示例和文档！

## 📚 文档概览

本目录包含了完整的亚马逊 SP-API PHP 开发指南，包括示例代码、参数说明、注意事项和最佳实践。

## 📁 文档结构

```
docs/
├── README.md                           # 本文档 - 文档概览
├── amazon-sp-api-php-examples.md      # 详细接口文档
├── quick-reference.md                  # 快速参考指南
└── examples-list.md                    # 示例文件列表
```

## 🚀 快速开始

### 1. 查看示例文件列表
- [examples-list.md](./examples-list.md) - 所有示例文件的详细列表和说明

### 2. 参考快速指南
- [quick-reference.md](./quick-reference.md) - 常用参数、错误处理和最佳实践

### 3. 阅读详细文档
- [amazon-sp-api-php-examples.md](./amazon-sp-api-php-examples.md) - 完整的接口文档和注意事项

## 📋 包含的 API 接口

### ✅ 已实现的接口

| API 类别 | 接口数量 | 状态 | 说明 |
|----------|----------|------|------|
| Orders API | 1 | ✅ 完成 | 获取订单列表 |
| Product Type Definitions API | 2 | ✅ 完成 | 产品类型定义和搜索 |
| Catalog Items API | 2 | ✅ 完成 | 商品目录搜索和详情 |
| Listings Items API | 4 | ✅ 完成 | 商品列表的获取、创建、更新和搜索 |
| Listings Restrictions API | 1 | ✅ 完成 | 商品限制查询 |
| Catalog Categories API | 1 | ✅ 完成 | 商品分类查询 (原生 HTTP) |

### 📊 统计信息

- **总接口数**: 11 个
- **示例文件数**: 22 个 (包含简化版和完整版)
- **API 类别**: 6 个
- **文档页数**: 4 个

## 🛠️ 示例文件特点

### 简化版示例
- 使用最少的必需参数
- 适合快速测试和学习
- 包含基本错误处理

### 完整版示例
- 包含所有可选参数
- 展示高级功能
- 详细的错误处理和注释

### 原生 HTTP 示例
- 用于 SDK 中不存在的接口
- 直接使用 Guzzle HTTP 客户端
- 手动处理认证和签名

## 🔧 环境配置

### 必需的环境变量
```env
SP_API_CLIENT_ID=你的clientId
SP_API_CLIENT_SECRET=你的clientSecret
SP_API_REFRESH_TOKEN=你的refreshToken
SP_API_ENDPOINT=https://api.amazon.com/auth/o2/token
SP_API_ENDPOINT_HOST=https://sandbox.sellingpartnerapi-na.amazon.com
SP_API_MARKETPLACE=ATVPDKIKX0DER
```

### 沙盒环境配置
- **LWA Token Endpoint**: `https://api.amazon.com/auth/o2/token`
- **SP-API Host**: `https://sandbox.sellingpartnerapi-na.amazon.com`
- **Marketplace ID**: `ATVPDKIKX0DER` (美国沙盒市场)

## 📝 重要注意事项

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

## 🎯 使用建议

### 新手用户
1. 先阅读 [quick-reference.md](./quick-reference.md)
2. 查看 [examples-list.md](./examples-list.md) 了解所有示例
3. 从简化版示例开始学习
4. 参考 [amazon-sp-api-php-examples.md](./amazon-sp-api-php-examples.md) 了解详细信息

### 有经验用户
1. 直接查看 [quick-reference.md](./quick-reference.md) 获取关键信息
2. 使用完整版示例作为开发参考
3. 参考详细文档了解高级功能和注意事项

### 生产环境
1. 确保使用正确的生产环境配置
2. 实现适当的错误处理和重试机制
3. 监控速率限制和 API 响应
4. 参考官方文档了解最新的 API 变更

## 🔗 相关资源

- [亚马逊 SP-API 官方文档](https://developer-docs.amazon.com/sp-api/)
- [SP-API 使用指南](https://developer-docs.amazon.com/sp-api/docs/sp-api-overview)
- [速率限制说明](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api)
- [沙盒环境说明](https://developer-docs.amazon.com/sp-api/docs/sp-api-sandbox)

## 🤝 贡献

如果您发现任何问题或有改进建议，请：

1. 检查示例代码是否正确运行
2. 验证文档信息是否准确
3. 测试沙盒参数是否有效
4. 提供详细的错误信息和解决方案

## 📄 许可证

本文档和示例代码遵循项目的许可证条款。

---

**最后更新**: 2024年12月

**版本**: 1.0.0

**状态**: ✅ 完成 