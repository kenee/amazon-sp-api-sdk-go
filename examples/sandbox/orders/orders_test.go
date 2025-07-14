package sandbox

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/orders"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestOrdersAPISandbox(t *testing.T) {
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
		t.Skip("跳过 Orders API 沙盒测试：未找到 .env.sbx 文件")
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

	// 创建 Orders API 客户端
	ordersAPI := orders.NewOrdersAPI(config)

	// 创建上下文
	ctx := context.Background()

	t.Run("沙盒环境 - 获取订单 (简化版本)", func(t *testing.T) {
		marketplaceIds := []string{"ATVPDKIKX0DER"}
		// 沙盒环境使用特殊参数
		createdAfter := "TEST_CASE_200"

		response, err := ordersAPI.GetOrdersSimple(ctx, marketplaceIds, createdAfter)
		if err != nil {
			t.Fatalf("沙盒环境获取订单失败: %v", err)
		}

		t.Logf("✅ 沙盒环境订单获取成功")
		if response.Payload != nil {
			t.Logf("   找到 %d 个订单", len(response.Payload.Orders))
			for i, order := range response.Payload.Orders {
				t.Logf("   %d. Order ID: %s, Status: %s",
					i+1, order.AmazonOrderId, order.OrderStatus)
			}
		} else {
			t.Logf("   未找到订单")
		}
	})

	t.Run("沙盒环境 - 获取特定状态订单", func(t *testing.T) {
		marketplaceIds := []string{"ATVPDKIKX0DER"}
		orderStatuses := []string{"Shipped", "Unshipped"}

		response, err := ordersAPI.GetOrdersWithStatus(ctx, marketplaceIds, orderStatuses)
		if err != nil {
			t.Fatalf("沙盒环境获取特定状态订单失败: %v", err)
		}

		t.Logf("✅ 沙盒环境特定状态订单获取成功")
		if response.Payload != nil {
			t.Logf("   找到 %d 个订单", len(response.Payload.Orders))
			for i, order := range response.Payload.Orders {
				t.Logf("   %d. Order ID: %s, Status: %s",
					i+1, order.AmazonOrderId, order.OrderStatus)
			}
		} else {
			t.Logf("   未找到订单")
		}
	})

	t.Run("沙盒环境 - 获取订单 (完整版本)", func(t *testing.T) {
		request := &orders.GetOrdersRequest{
			MarketplaceIds:    []string{"ATVPDKIKX0DER"},
			CreatedAfter:      "TEST_CASE_200", // 沙盒环境特殊参数
			OrderStatuses:     []string{"Shipped", "Unshipped"},
			MaxResultsPerPage: 10,
		}

		response, err := ordersAPI.GetOrders(ctx, request)
		if err != nil {
			t.Fatalf("沙盒环境获取完整订单信息失败: %v", err)
		}

		t.Logf("✅ 沙盒环境完整订单信息获取成功")
		if response.Payload != nil {
			t.Logf("   找到 %d 个订单", len(response.Payload.Orders))
			for i, order := range response.Payload.Orders {
				t.Logf("   %d. Order ID: %s", i+1, order.AmazonOrderId)
				t.Logf("      状态: %s", order.OrderStatus)
				t.Logf("      创建时间: %s", order.PurchaseDate)
				t.Logf("      履行渠道: %s", order.FulfillmentChannel)
				if order.BuyerInfo != nil {
					t.Logf("      买家邮箱: %s", order.BuyerInfo.BuyerEmail)
				}
			}
		} else {
			t.Logf("   未找到订单")
		}
	})
}
