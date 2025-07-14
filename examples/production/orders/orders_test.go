package production

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/orders"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestOrdersAPIProduction(t *testing.T) {
	// 1. 加载环境变量
	envPaths := []string{
		".env.prod",          // 当前目录
		"../.env.prod",       // 上级目录
		"../../.env.prod",    // 上上级目录
		"../../../.env.prod", // 根目录
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
		t.Skip("跳过 Orders API 生产环境测试：未找到 .env 或 .env.prod1 文件")
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

	t.Run("生产环境 - 获取订单 (简化版本)", func(t *testing.T) {
		marketplaceIds := []string{"ATVPDKIKX0DER"}
		// 使用与PHP生产环境示例一致的时间参数：30天前
		createdAfter := time.Now().AddDate(0, 0, -30).Format("2006-01-02T15:04:05Z")

		response, err := ordersAPI.GetOrdersSimple(ctx, marketplaceIds, createdAfter)
		if err != nil {
			t.Fatalf("生产环境获取订单失败: %v", err)
		}

		t.Logf("✅ 生产环境订单获取成功")
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

	t.Run("生产环境 - 获取特定状态订单", func(t *testing.T) {
		marketplaceIds := []string{"ATVPDKIKX0DER"}
		orderStatuses := []string{"Shipped", "Unshipped"}

		response, err := ordersAPI.GetOrdersWithStatus(ctx, marketplaceIds, orderStatuses)
		if err != nil {
			t.Fatalf("生产环境获取特定状态订单失败: %v", err)
		}

		t.Logf("✅ 生产环境特定状态订单获取成功")
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

	t.Run("生产环境 - 获取订单 (完整版本)", func(t *testing.T) {
		// 使用与PHP生产环境示例一致的参数
		request := &orders.GetOrdersRequest{
			MarketplaceIds: []string{"ATVPDKIKX0DER"},                                    // 美国生产环境
			CreatedAfter:   time.Now().AddDate(0, 0, -30).Format("2006-01-02T15:04:05Z"), // 30天前
			// 其他参数保持null，与PHP示例一致
		}

		response, err := ordersAPI.GetOrders(ctx, request)
		if err != nil {
			t.Fatalf("生产环境获取完整订单信息失败: %v", err)
		}

		t.Logf("✅ 生产环境完整订单信息获取成功")
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
