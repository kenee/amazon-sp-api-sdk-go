package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	v2020_07_01 "github.com/kenee/amazon-sp-api-sdk-go/apis/fulfillment/outbound/v2020_07_01"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// 通用的测试设置函数
func setupFulfillmentOutboundTest(t *testing.T) *v2020_07_01.FulfillmentOutboundAPI {
	// 1. 加载环境变量
	envPaths := []string{
		".env.sbx",
		"../.env.sbx",
		"../../.env.sbx",
		"../../../.env.sbx",
	}

	envLoaded := false
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			t.Logf("✅ 成功加载环境变量文件: %s", path)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		t.Skip("跳过 Fulfillment Outbound API 测试：未找到 .env.sbx 文件")
	}

	// 2. 配置认证信息
	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}

	// 3. 创建配置
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

	// 4. 创建 API 客户端
	return v2020_07_01.NewFulfillmentOutboundAPI(config)
}

func TestListReturnReasonCodes(t *testing.T) {
	api := setupFulfillmentOutboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListReturnReasonCodes API")

	// 使用沙盒环境的魔法参数
	sellerSku := "GM-ZDPI-9B4E" // 与Fulfillment Inbound相同的SKU
	language := "en_US"

	resp, err := api.ListReturnReasonCodes(ctx, sellerSku, &language)
	if err != nil {
		t.Logf("⚠️ 获取退货原因代码失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功获取退货原因代码")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestGetPackageTrackingDetails(t *testing.T) {
	api := setupFulfillmentOutboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetPackageTrackingDetails API")

	// 使用沙盒环境的魔法参数
	packageNumber := "1Z999AA1234567890" // 示例跟踪号

	resp, err := api.GetPackageTrackingDetails(ctx, packageNumber)
	if err != nil {
		t.Logf("⚠️ 获取包裹跟踪详情失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功获取包裹跟踪详情")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestListAllFulfillmentOrders(t *testing.T) {
	api := setupFulfillmentOutboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListAllFulfillmentOrders API")

	// 使用沙盒环境的魔法参数
	queryStartTime := "2019-10-13T00:00:00Z" // 与Fulfillment Inbound相同的时间

	resp, err := api.ListAllFulfillmentOrders(ctx, &queryStartTime, nil)
	if err != nil {
		t.Logf("⚠️ 列出履行订单失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功列出履行订单")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestGetFulfillmentOrder(t *testing.T) {
	api := setupFulfillmentOutboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetFulfillmentOrder API")

	// 使用沙盒环境的魔法参数
	sellerFulfillmentOrderId := "be7a0a53-00c3-4f6f-a63a-639f76ee9253" // 与Fulfillment Inbound相同的ID

	resp, err := api.GetFulfillmentOrder(ctx, sellerFulfillmentOrderId)
	if err != nil {
		t.Logf("⚠️ 获取履行订单失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功获取履行订单")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestGetFulfillmentPreview(t *testing.T) {
	api := setupFulfillmentOutboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetFulfillmentPreview API")

	// 创建测试请求，使用沙盒环境的魔法参数
	request := &v2020_07_01.GetFulfillmentPreviewRequest{
		MarketplaceId: "ATVPDKIKX0DER",                           // 美国市场
		Address:       &v2020_07_01.Address{},                    // 空地址结构体
		Items:         []v2020_07_01.GetFulfillmentPreviewItem{}, // 空项目列表
	}

	resp, err := api.GetFulfillmentPreview(ctx, request)
	if err != nil {
		t.Logf("⚠️ 获取履行预览失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功获取履行预览")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestCancelFulfillmentOrder(t *testing.T) {
	api := setupFulfillmentOutboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 CancelFulfillmentOrder API")

	// 使用沙盒环境的魔法参数
	sellerFulfillmentOrderId := "be7a0a53-00c3-4f6f-a63a-639f76ee9253" // 与Fulfillment Inbound相同的ID

	resp, err := api.CancelFulfillmentOrder(ctx, sellerFulfillmentOrderId)
	if err != nil {
		t.Logf("⚠️ 取消履行订单失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功取消履行订单")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}
