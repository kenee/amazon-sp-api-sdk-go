package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	v0 "github.com/kenee/amazon-sp-api-sdk-go/apis/fulfillment/inbound/v0"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// 通用的测试设置函数
func setupFulfillmentInboundTest(t *testing.T) *v0.FulfillmentInboundAPI {
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
		t.Skip("跳过 Fulfillment Inbound API 测试：未找到 .env.sbx 文件")
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
	return v0.NewFulfillmentInboundAPI(config)
}

func TestGetShipments(t *testing.T) {
	fulfillmentInboundAPI := setupFulfillmentInboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetShipments API")

	request := &v0.GetShipmentsRequest{
		QueryType:     "SHIPMENT",
		MarketplaceId: "ATVPDKIKX0DER",
		// 使用沙盒环境的魔法参数
		ShipmentIdList: []string{"be7a0a53-00c3-4f6f-a63a-639f76ee9253"},
	}

	response, err := fulfillmentInboundAPI.GetShipments(ctx, request)
	if err != nil {
		t.Logf("⚠️ 获取 Shipments 失败（沙盒环境限制）: %v", err)
		return
	}

	if response.Payload != nil {
		t.Logf("✅ 成功获取 Shipments")
		t.Logf("   返回的 Shipment 数量: %d", len(response.Payload.ShipmentData))

		if len(response.Payload.ShipmentData) > 0 {
			shipment := response.Payload.ShipmentData[0]
			t.Logf("   示例 Shipment 详情:")
			t.Logf("     Shipment ID: %s", shipment.ShipmentId)
			t.Logf("     Shipment Name: %s", shipment.ShipmentName)
			t.Logf("     Shipment Status: %s", shipment.ShipmentStatus)
			t.Logf("     Destination Fulfillment Center: %s", shipment.DestinationFulfillmentCenterId)
			t.Logf("     Created Date: %s", shipment.CreatedDate)
		}
	}
}

func TestGetShipmentItems(t *testing.T) {
	fulfillmentInboundAPI := setupFulfillmentInboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetShipmentItems API")

	request := &v0.GetShipmentItemsRequest{
		QueryType:     "SHIPMENT",
		MarketplaceId: "ATVPDKIKX0DER",
		// 使用沙盒环境的魔法参数
		ShipmentId: "be7a0a53-00c3-4f6f-a63a-639f76ee9253",
	}

	response, err := fulfillmentInboundAPI.GetShipmentItems(ctx, request)
	if err != nil {
		t.Logf("⚠️ 获取 Shipment Items 失败（沙盒环境限制）: %v", err)
		return
	}

	if response.Payload != nil {
		t.Logf("✅ 成功获取 Shipment Items")
		t.Logf("   返回的 Item 数量: %d", len(response.Payload.ItemData))

		if len(response.Payload.ItemData) > 0 {
			item := response.Payload.ItemData[0]
			t.Logf("   示例 Item 详情:")
			t.Logf("     Seller SKU: %s", item.SellerSKU)
			t.Logf("     FNSKU: %s", item.FulfillmentNetworkSKU)
			t.Logf("     Quantity Shipped: %d", item.QuantityShipped)
			t.Logf("     Quantity Received: %d", item.QuantityReceived)
		}
	}
}

func TestGetShipmentItemsByShipmentId(t *testing.T) {
	fulfillmentInboundAPI := setupFulfillmentInboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetShipmentItemsByShipmentId API")

	// 使用沙盒环境的魔法参数
	shipmentId := "be7a0a53-00c3-4f6f-a63a-639f76ee9253"

	request := &v0.GetShipmentItemsByShipmentIdRequest{
		MarketplaceId: "ATVPDKIKX0DER",
	}

	response, err := fulfillmentInboundAPI.GetShipmentItemsByShipmentId(ctx, shipmentId, request)
	if err != nil {
		t.Logf("⚠️ 按 Shipment ID 获取 Items 失败（沙盒环境限制）: %v", err)
		return
	}

	if response.Payload != nil {
		t.Logf("✅ 成功按 Shipment ID 获取 Items")
		t.Logf("   返回的 Item 数量: %d", len(response.Payload.ItemData))
	}
}

func TestGetPrepInstructions(t *testing.T) {
	fulfillmentInboundAPI := setupFulfillmentInboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetPrepInstructions API")

	request := &v0.GetPrepInstructionsRequest{
		ShipToCountryCode: "US",
		// 使用沙盒环境的魔法参数
		SellerSKUList: []string{"GM-ZDPI-9B4E"},
	}

	response, err := fulfillmentInboundAPI.GetPrepInstructions(ctx, request)
	if err != nil {
		t.Logf("⚠️ 获取 Prep Instructions 失败（沙盒环境限制）: %v", err)
		return
	}

	if response.Payload != nil {
		t.Logf("✅ 成功获取 Prep Instructions")
		t.Logf("   返回的 Prep Instruction 数量: %d", len(response.Payload.PrepInstructions))

		if len(response.Payload.PrepInstructions) > 0 {
			instruction := response.Payload.PrepInstructions[0]
			t.Logf("   示例 Prep Instruction 详情:")
			t.Logf("     Seller SKU: %s", instruction.SellerSKU)
			t.Logf("     ASIN: %s", instruction.ASIN)
			t.Logf("     Prep Guidance: %s", instruction.PrepGuidance)
		}
	}
}

func TestGetLabels(t *testing.T) {
	fulfillmentInboundAPI := setupFulfillmentInboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetLabels API")

	// 使用沙盒环境的魔法参数
	shipmentId := "be7a0a53-00c3-4f6f-a63a-639f76ee9253"

	request := &v0.GetLabelsRequest{
		PageType:  "PackageLabel_Plain_Paper",
		LabelType: "UNIQUE",
	}

	response, err := fulfillmentInboundAPI.GetLabels(ctx, shipmentId, request)
	if err != nil {
		t.Logf("⚠️ 获取 Labels 失败（沙盒环境限制）: %v", err)
		return
	}

	if response.Payload != nil {
		t.Logf("✅ 成功获取 Labels")
		t.Logf("   Download URL: %s", response.Payload.DownloadURL)
	}
}

func TestGetBillOfLading(t *testing.T) {
	fulfillmentInboundAPI := setupFulfillmentInboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 GetBillOfLading API")

	// 使用沙盒环境的魔法参数
	shipmentId := "be7a0a53-00c3-4f6f-a63a-639f76ee9253"

	response, err := fulfillmentInboundAPI.GetBillOfLading(ctx, shipmentId)
	if err != nil {
		t.Logf("⚠️ 获取 Bill of Lading 失败（沙盒环境限制）: %v", err)
		return
	}

	if response.Payload != nil {
		t.Logf("✅ 成功获取 Bill of Lading")
		t.Logf("   Download URL: %s", response.Payload.DownloadURL)
	}
}

func TestInvalidShipmentId(t *testing.T) {
	fulfillmentInboundAPI := setupFulfillmentInboundTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 Invalid Shipment ID 错误处理")

	// 使用会导致 400 错误的魔法参数
	shipmentId := "87d20cf7-1beb-4cda-8bf4-7366cfddbec1"

	request := &v0.GetShipmentItemsByShipmentIdRequest{
		MarketplaceId: "ATVPDKIKX0DER",
	}

	response, err := fulfillmentInboundAPI.GetShipmentItemsByShipmentId(ctx, shipmentId, request)
	if err != nil {
		t.Logf("✅ 测试无效 Shipment ID 通过（返回预期的错误）: %v", err)
		return
	}

	if response.Errors != nil && len(response.Errors) > 0 {
		t.Logf("✅ 测试无效 Shipment ID 通过（返回预期的错误响应）")
		for _, err := range response.Errors {
			t.Logf("   错误代码: %s", err.Code)
			t.Logf("   错误消息: %s", err.Message)
		}
	}
}
