package listing_flow

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/listings"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/producttypes"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// 注意：这个文件需要先安装 kenee/amazon-sp-api-sdk-go
// 实际使用时需要：
// 1. go get github.com/kenee/amazon-sp-api-sdk-go@latest
// 2. 配置 .env 文件
// 3. 确保有有效的 SP-API 凭据

// 测试配置
type TestConfig struct {
	SellerId       string
	SKU            string
	MarketplaceIds []string
	ProductType    string
	IssueLocale    string
}

// 全局测试配置 - 与 PHP sandbox 示例保持一致
var testConfig = TestConfig{
	SellerId:       "A1B2C3D4E5F6G7",          // 与 PHP sandbox 示例一致
	SKU:            "GM-ZDPI-9B4E",            // 与 PHP sandbox 示例一致
	MarketplaceIds: []string{"ATVPDKIKX0DER"}, // 美国沙盒市场
	ProductType:    "LUGGAGE",
	IssueLocale:    "en_US",
}

// loadTestConfig 统一加载测试配置
func loadTestConfig(t *testing.T) *client.Configuration {
	// 加载环境变量 - 尝试多个可能的路径
	envPaths := []string{
		".env.sbx",          // 当前目录
		"../.env.sbx",       // 上级目录
		"../../.env.sbx",    // 上上级目录
		"../../../.env.sbx", // 根目录
	}

	var envLoaded bool
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			t.Logf("✅ 成功加载环境变量文件: %s", path)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		t.Skip("跳过测试：未找到 .env.sbx 文件")
		return nil
	}

	// 检查必要的环境变量
	requiredEnvVars := []string{
		"SP_API_CLIENT_ID",
		"SP_API_CLIENT_SECRET",
		"SP_API_REFRESH_TOKEN",
		"SP_API_ENDPOINT",
		"SP_API_ENDPOINT_HOST",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			t.Skipf("跳过测试：缺少环境变量 %s", envVar)
			return nil
		}
	}

	// 设置认证凭据
	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}

	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

	return config
}

// TestListingFlow_Setup 设置刊登流程测试环境
func TestListingFlow_Setup(t *testing.T) {
	t.Log("=== 设置刊登流程测试环境 ===")

	// 加载环境变量 - 尝试多个可能的路径
	envPaths := []string{
		".env.sbx",
		"../.env.sbx",
		"../../.env.sbx",
		"../../../.env.sbx",
	}

	var envLoaded bool
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			t.Logf("✅ 成功加载环境变量文件: %s", path)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		t.Skip("跳过刊登流程测试：未找到 .env.sbx 文件")
		return
	}

	// 检查必要的环境变量
	requiredEnvVars := []string{
		"SP_API_CLIENT_ID",
		"SP_API_CLIENT_SECRET",
		"SP_API_REFRESH_TOKEN",
		"SP_API_ENDPOINT",
		"SP_API_ENDPOINT_HOST",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			t.Skipf("跳过刊登流程测试：缺少环境变量 %s", envVar)
			return
		}
	}

	t.Log("✅ 环境变量配置完成")
	t.Log("✅ 准备进行刊登流程测试")
}

// TestStep1_GetProductTypeDefinition 步骤1: 获取产品类型定义
func TestStep1_GetProductTypeDefinition(t *testing.T) {
	t.Log("=== 步骤1: 获取产品类型定义 ===")

	config := loadTestConfig(t)
	if config == nil {
		return // 测试已被跳过
	}

	productTypesAPI := producttypes.NewProductTypesAPI(config)
	request := &producttypes.GetDefinitionsProductTypeRequest{
		ProductType:    testConfig.ProductType,
		MarketplaceIds: testConfig.MarketplaceIds,
		Locale:         testConfig.IssueLocale,
	}
	ctx := context.Background()
	response, err := productTypesAPI.GetDefinitionsProductType(ctx, request)
	if err != nil {
		t.Fatalf("获取产品类型定义失败: %v", err)
	}
	t.Logf("✅ 产品类型: %s", response.ProductType)
	t.Logf("✅ 显示名称: %s", response.DisplayName)
	t.Logf("✅ 要求: %s", response.Requirements)
	t.Logf("✅ 要求强制执行: %s", response.RequirementsEnforced)
	if response.PropertyGroups != nil {
		for groupName, group := range response.PropertyGroups {
			t.Logf("📋 属性组: %s - %s", groupName, group.Title)
			for _, propName := range group.PropertyNames {
				t.Logf("   - %s", propName)
			}
		}
	}
}

// TestStep2_GetListingsRestrictions 步骤2: 检查刊登限制
func TestStep2_GetListingsRestrictions(t *testing.T) {
	t.Log("=== 步骤2: 检查刊登限制 ===")

	if os.Getenv("SP_API_CLIENT_ID") == "" {
		t.Skip("跳过：缺少 SP-API 凭据")
		return
	}

	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))
	listingsAPI := listings.NewListingsAPI(config)
	ctx := context.Background()

	// 使用与 PHP sandbox 示例相同的 ASIN
	asin := "B07N4M94X4" // 与 PHP sandbox 示例一致
	t.Logf("用于限制检查的 ASIN: %s", asin)

	req := &listings.GetListingsRestrictionsRequest{
		SellerId:       testConfig.SellerId,
		MarketplaceIds: testConfig.MarketplaceIds,
		ASIN:           asin,
		IssueLocale:    testConfig.IssueLocale,
	}
	response, err := listingsAPI.GetListingsRestrictions(ctx, req)
	if err != nil {
		t.Fatalf("检查刊登限制失败: %v", err)
	}
	if len(response.Restrictions) > 0 {
		t.Logf("⚠️ 发现 %d 个刊登限制", len(response.Restrictions))
		for _, restriction := range response.Restrictions {
			t.Logf("📋 市场ID: %s", restriction.MarketplaceId)
			t.Logf("📋 条件类型: %s", restriction.ConditionType)
			for _, reason := range restriction.Reasons {
				t.Logf("   - 原因: %s", reason.Message)
				if reason.ReasonCode != "" {
					t.Logf("   - 错误代码: %s", reason.ReasonCode)
				}
			}
		}
	} else {
		t.Log("✅ 没有发现刊登限制")
	}
}

// TestStep3_PutListingsItem 步骤3: 创建产品刊登
func TestStep3_PutListingsItem(t *testing.T) {
	t.Log("=== 步骤3: 创建产品刊登 ===")

	if os.Getenv("SP_API_CLIENT_ID") == "" {
		t.Skip("跳过：缺少 SP-API 凭据")
		return
	}

	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))
	listingsAPI := listings.NewListingsAPI(config)
	request := &listings.PutListingsItemRequest{
		SellerId:       testConfig.SellerId,
		SKU:            testConfig.SKU,
		MarketplaceIds: testConfig.MarketplaceIds,
		ProductType:    testConfig.ProductType,
		Requirements:   "LISTING",
		Attributes: map[string]interface{}{
			"item_name": []map[string]interface{}{
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "Hardside Carry-On Spinner Suitcase Luggage",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
			"brand": []map[string]interface{}{
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "TravelPro",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
			"bullet_point": []map[string]interface{}{
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "Hardside spinner luggage with TSA lock",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "Expandable design for extra packing space",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "Water-resistant polycarbonate shell",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
			"color": []map[string]interface{}{
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "Black",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
			"item_dimensions": []map[string]interface{}{
				{
					"width": map[string]interface{}{
						"unit":  "inches",
						"value": 22.0,
					},
					"length": map[string]interface{}{
						"unit":  "inches",
						"value": 14.0,
					},
					"height": map[string]interface{}{
						"unit":  "inches",
						"value": 9.0,
					},
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
			"item_weight": []map[string]interface{}{
				{
					"unit":           "pounds",
					"value":          7.5,
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
			"material": []map[string]interface{}{
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "Polycarbonate",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
			"warranty_description": []map[string]interface{}{
				{
					"language_tag":   testConfig.IssueLocale,
					"value":          "1 year limited warranty",
					"marketplace_id": testConfig.MarketplaceIds[0],
				},
			},
		},
		IssueLocale: testConfig.IssueLocale,
	}
	ctx := context.Background()
	response, err := listingsAPI.PutListingsItem(ctx, request)
	if err != nil {
		t.Fatalf("创建产品刊登失败: %v", err)
	}
	t.Logf("✅ SKU: %s", response.SKU)
	t.Logf("✅ 状态: %s", response.Status)
	t.Logf("✅ 提交ID: %s", response.SubmissionId)
	if len(response.Issues) > 0 {
		t.Logf("⚠️ 发现 %d 个问题", len(response.Issues))
		for _, issue := range response.Issues {
			t.Logf("   - 代码: %s", issue.Code)
			t.Logf("   - 消息: %s", issue.Message)
			t.Logf("   - 严重性: %s", issue.Severity)
		}
	} else {
		t.Log("✅ 刊登创建成功，没有发现问题")
	}
}

// TestStep4_GetListingsItem 步骤4: 检查刊登状态
func TestStep4_GetListingsItem(t *testing.T) {
	t.Log("=== 步骤4: 检查刊登状态 ===")

	if os.Getenv("SP_API_CLIENT_ID") == "" {
		t.Skip("跳过：缺少 SP-API 凭据")
		return
	}

	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))
	listingsAPI := listings.NewListingsAPI(config)
	request := &listings.GetListingsItemRequest{
		SellerId:       testConfig.SellerId,
		SKU:            testConfig.SKU,
		MarketplaceIds: testConfig.MarketplaceIds,
		IssueLocale:    testConfig.IssueLocale,
		IncludedData:   []string{"summaries", "issues", "offers", "fulfillmentAvailability"},
	}
	ctx := context.Background()
	response, err := listingsAPI.GetListingsItem(ctx, request)
	if err != nil {
		t.Fatalf("检查刊登状态失败: %v", err)
	}
	t.Logf("✅ SKU: %s", response.SKU)
	if len(response.Summaries) > 0 {
		summary := response.Summaries[0]
		t.Logf("✅ 市场ID: %s", summary.MarketplaceId)
		t.Logf("✅ ASIN: %s", summary.ASIN)
		t.Logf("✅ 产品类型: %s", summary.ProductType)
		t.Logf("✅ 商品名称: %s", summary.ItemName)
	}
	if len(response.Issues) > 0 {
		t.Logf("⚠️ 发现 %d 个问题", len(response.Issues))
		for _, issue := range response.Issues {
			t.Logf("   - 代码: %s", issue.Code)
			t.Logf("   - 消息: %s", issue.Message)
			t.Logf("   - 严重性: %s", issue.Severity)
		}
	} else {
		t.Log("✅ 没有发现问题")
	}
	if len(response.Offers) > 0 {
		offer := response.Offers[0]
		if len(offer.Offers) > 0 && offer.Offers[0].Price != nil {
			t.Logf("💰 价格: %s %s", offer.Offers[0].Price.Amount, offer.Offers[0].Price.CurrencyCode)
		}
	}
	if len(response.FulfillmentAvailability) > 0 {
		fa := response.FulfillmentAvailability[0]
		if len(fa.FulfillmentAvailability) > 0 {
			t.Logf("📦 库存: %d", fa.FulfillmentAvailability[0].Quantity)
			t.Logf("📦 履行渠道: %s", fa.FulfillmentAvailability[0].FulfillmentChannel)
		}
	}
}

// TestStep5_PatchListingsItem 步骤5: 更新产品信息
func TestStep5_PatchListingsItem(t *testing.T) {
	t.Log("=== 步骤5: 更新产品信息 ===")

	if os.Getenv("SP_API_CLIENT_ID") == "" {
		t.Skip("跳过：缺少 SP-API 凭据")
		return
	}

	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))
	listingsAPI := listings.NewListingsAPI(config)
	patches := []listings.PatchOperation{
		{
			Op:   "add",
			Path: "/attributes/size",
			Value: []map[string]interface{}{
				{
					"marketplace_id": testConfig.MarketplaceIds[0],
					"language_tag":   testConfig.IssueLocale,
					"value":          "Medium",
				},
			},
		},
		{
			Op:   "replace",
			Path: "/offers/0/price/amount",
			Value: []map[string]interface{}{
				{
					"marketplace_id": testConfig.MarketplaceIds[0],
					"value":          "89.99",
				},
			},
		},
		{
			Op:   "add",
			Path: "/attributes/color",
			Value: []map[string]interface{}{
				{
					"marketplace_id": testConfig.MarketplaceIds[0],
					"language_tag":   testConfig.IssueLocale,
					"value":          "Black",
				},
			},
		},
		{
			Op:   "add",
			Path: "/attributes/material",
			Value: []map[string]interface{}{
				{
					"marketplace_id": testConfig.MarketplaceIds[0],
					"language_tag":   testConfig.IssueLocale,
					"value":          "Polycarbonate",
				},
			},
		},
		{
			Op:   "add",
			Path: "/attributes/closure_type",
			Value: []map[string]interface{}{
				{
					"marketplace_id": testConfig.MarketplaceIds[0],
					"language_tag":   testConfig.IssueLocale,
					"value":          "Zipper",
				},
			},
		},
	}
	request := &listings.PatchListingsItemRequest{
		SellerId:       testConfig.SellerId,
		SKU:            testConfig.SKU,
		MarketplaceIds: testConfig.MarketplaceIds,
		ProductType:    testConfig.ProductType,
		Patches:        patches,
		IssueLocale:    testConfig.IssueLocale,
	}
	ctx := context.Background()
	response, err := listingsAPI.PatchListingsItem(ctx, request)
	if err != nil {
		t.Fatalf("更新产品信息失败: %v", err)
	}
	t.Logf("✅ SKU: %s", response.SKU)
	t.Logf("✅ 状态: %s", response.Status)
	t.Logf("✅ 提交ID: %s", response.SubmissionId)
	if len(response.Issues) > 0 {
		t.Logf("⚠️ 发现 %d 个问题", len(response.Issues))
		for _, issue := range response.Issues {
			t.Logf("   - 代码: %s", issue.Code)
			t.Logf("   - 消息: %s", issue.Message)
			t.Logf("   - 严重性: %s", issue.Severity)
		}
	} else {
		t.Log("✅ 产品信息更新成功，没有发现问题")
	}
}

// TestStep6_SearchListingsItems 步骤6: 搜索刊登
func TestStep6_SearchListingsItems(t *testing.T) {
	t.Log("=== 步骤6: 搜索刊登 ===")

	if os.Getenv("SP_API_CLIENT_ID") == "" {
		t.Skip("跳过：缺少 SP-API 凭据")
		return
	}

	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))
	listingsAPI := listings.NewListingsAPI(config)
	ctx := context.Background()

	t.Logf("search params: sellerId=%s, marketplaceIds=%v, includedData=%v", testConfig.SellerId, testConfig.MarketplaceIds, []string{"summaries", "issues"})

	request := &listings.SearchListingsItemsRequest{
		SellerId:        testConfig.SellerId,
		MarketplaceIds:  testConfig.MarketplaceIds,
		IssueLocale:     testConfig.IssueLocale,
		IncludedData:    []string{"summaries", "issues"},
		Identifiers:     []string{testConfig.SKU}, // 添加SKU作为标识符
		IdentifiersType: "SKU",                    // 指定标识符类型
		PageSize:        10,
	}
	response, err := listingsAPI.SearchListingsItems(ctx, request)
	if err != nil {
		t.Fatalf("搜索刊登失败: %v", err)
	}
	t.Logf("✅ 找到 %d 个刊登", len(response.Listings))
	for i, listing := range response.Listings {
		t.Logf("📋 刊登 %d:", i+1)
		t.Logf("   - SKU: %s", listing.SKU)
		t.Logf("   - 状态: %s", listing.Status)
		if len(listing.Summaries) > 0 {
			summary := listing.Summaries[0]
			t.Logf("   - 商品名称: %s", summary.ItemName)
			t.Logf("   - ASIN: %s", summary.ASIN)
			t.Logf("   - 产品类型: %s", summary.ProductType)
		}
		if len(listing.Issues) > 0 {
			t.Logf("   - 问题数量: %d", len(listing.Issues))
		}
	}
	if response.Pagination != nil {
		t.Logf("📄 分页信息:")
		t.Logf("   - 下一页令牌: %s", response.Pagination.NextToken)
		t.Logf("   - 上一页令牌: %s", response.Pagination.PreviousToken)
	}
}

// TestCompleteListingFlow 完整刊登流程测试
func TestCompleteListingFlow(t *testing.T) {
	t.Log("🚀 开始完整的 Amazon 刊登发品流程测试")

	// 加载配置
	config := loadTestConfig(t)
	if config == nil {
		return // 测试已被跳过
	}

	// 按顺序执行所有步骤
	t.Run("Step1_GetProductTypeDefinition", TestStep1_GetProductTypeDefinition)
	time.Sleep(1 * time.Second) // 避免频率限制

	t.Run("Step2_GetListingsRestrictions", TestStep2_GetListingsRestrictions)
	time.Sleep(1 * time.Second)

	t.Run("Step3_PutListingsItem", TestStep3_PutListingsItem)
	time.Sleep(2 * time.Second) // 给刊登处理一些时间

	t.Run("Step4_GetListingsItem", TestStep4_GetListingsItem)
	time.Sleep(1 * time.Second)

	t.Run("Step5_PatchListingsItem", TestStep5_PatchListingsItem)
	time.Sleep(1 * time.Second)

	t.Run("Step6_SearchListingsItems", TestStep6_SearchListingsItems)
	time.Sleep(1 * time.Second)

	t.Log("✅ 完整刊登流程测试完成")
}

// 辅助函数：打印分隔线
func printSeparator() {
	fmt.Println(strings.Repeat("=", 50))
}
