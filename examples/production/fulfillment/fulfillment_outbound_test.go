package production

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/fulfillment/outbound/v2020_07_01"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestFulfillmentOutboundAPIProduction(t *testing.T) {
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
		t.Skip("跳过 Fulfillment Outbound API 生产环境测试：未找到 .env 或 .env.prod1 文件")
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
	config.SetUserAgent("Go-SDK-Test/1.0")

	// 创建API客户端
	api := v2020_07_01.NewFulfillmentOutboundAPI(config)
	ctx := context.Background()

	t.Run("生产环境 - 获取退货原因代码", func(t *testing.T) {
		// 使用生产环境的真实参数
		sellerSku := "REAL-SKU-123"
		language := "en_US"

		resp, err := api.ListReturnReasonCodes(ctx, sellerSku, &language)
		if err != nil {
			t.Fatalf("获取退货原因代码失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功获取退货原因代码响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 获取包裹跟踪详情", func(t *testing.T) {
		// 使用生产环境的真实参数
		packageNumber := "REAL-PACKAGE-123"

		resp, err := api.GetPackageTrackingDetails(ctx, packageNumber)
		if err != nil {
			t.Fatalf("获取包裹跟踪详情失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功获取包裹跟踪详情响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 列出所有履行订单", func(t *testing.T) {
		// 使用生产环境的真实参数
		queryStartTime := "2024-01-01T00:00:00Z"

		resp, err := api.ListAllFulfillmentOrders(ctx, &queryStartTime, nil)
		if err != nil {
			t.Fatalf("列出履行订单失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功列出履行订单响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 获取履行订单", func(t *testing.T) {
		// 使用生产环境的真实参数
		sellerFulfillmentOrderId := "REAL-ORDER-123"

		resp, err := api.GetFulfillmentOrder(ctx, sellerFulfillmentOrderId)
		if err != nil {
			t.Fatalf("获取履行订单失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功获取履行订单响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 获取履行预览", func(t *testing.T) {
		// 创建真实的测试请求
		request := &v2020_07_01.GetFulfillmentPreviewRequest{
			MarketplaceId: "ATVPDKIKX0DER",                           // 美国市场
			Address:       &v2020_07_01.Address{},                    // 真实地址结构体
			Items:         []v2020_07_01.GetFulfillmentPreviewItem{}, // 真实项目列表
		}

		resp, err := api.GetFulfillmentPreview(ctx, request)
		if err != nil {
			t.Fatalf("获取履行预览失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功获取履行预览响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 取消履行订单", func(t *testing.T) {
		// 使用生产环境的真实参数
		sellerFulfillmentOrderId := "REAL-ORDER-123"

		resp, err := api.CancelFulfillmentOrder(ctx, sellerFulfillmentOrderId)
		if err != nil {
			t.Fatalf("取消履行订单失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功取消履行订单响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	// 添加延迟避免频率限制
	time.Sleep(1 * time.Second)
}
