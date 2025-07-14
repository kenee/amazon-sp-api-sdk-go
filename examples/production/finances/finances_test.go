package production

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	v0 "github.com/kenee/amazon-sp-api-sdk-go/apis/finances/v0"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestFinancesAPIProduction(t *testing.T) {
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
		t.Skip("跳过 Feeds API 生产环境测试：未找到 .env 或 .env.prod1 文件")
	}

	// 设置认证凭据
	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}

	// 创建配置
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

	// 创建API客户端
	api := v0.NewFinancesAPI(config)
	ctx := context.Background()

	t.Run("生产环境 - 列出财务事件组", func(t *testing.T) {
		// 使用生产环境的真实参数
		request := &v0.ListFinancialEventGroupsRequest{
			MaxResultsPerPage:               10,
			FinancialEventGroupStartedAfter: "2024-01-01T00:00:00Z",
		}

		resp, err := api.ListFinancialEventGroups(ctx, request)
		if err != nil {
			t.Fatalf("列出财务事件组失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功列出财务事件组响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 列出财务事件", func(t *testing.T) {
		// 使用生产环境的真实参数
		request := &v0.ListFinancialEventsRequest{
			MaxResultsPerPage: 10,
			PostedAfter:       "2024-01-01T00:00:00Z",
		}

		resp, err := api.ListFinancialEvents(ctx, request)
		if err != nil {
			t.Fatalf("列出财务事件失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功列出财务事件响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 按组ID列出财务事件", func(t *testing.T) {
		// 使用生产环境的真实参数
		eventGroupId := "REAL-GROUP-123"
		request := &v0.ListFinancialEventsByGroupIdRequest{
			MaxResultsPerPage: 10,
		}

		resp, err := api.ListFinancialEventsByGroupId(ctx, eventGroupId, request)
		if err != nil {
			t.Fatalf("按组ID列出财务事件失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功按组ID列出财务事件响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 按订单ID列出财务事件", func(t *testing.T) {
		// 使用生产环境的真实参数
		orderId := "REAL-ORDER-123"
		request := &v0.ListFinancialEventsByOrderIdRequest{
			MaxResultsPerPage: 10,
		}

		resp, err := api.ListFinancialEventsByOrderId(ctx, orderId, request)
		if err != nil {
			t.Fatalf("按订单ID列出财务事件失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功按订单ID列出财务事件响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 简化版列出财务事件组", func(t *testing.T) {
		resp, err := api.ListFinancialEventGroupsSimple(ctx, 5)
		if err != nil {
			t.Fatalf("简化版列出财务事件组失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功简化版列出财务事件组响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	t.Run("生产环境 - 简化版列出财务事件", func(t *testing.T) {
		resp, err := api.ListFinancialEventsSimple(ctx, 5)
		if err != nil {
			t.Fatalf("简化版列出财务事件失败: %v", err)
		}

		if resp != nil {
			t.Logf("成功简化版列出财务事件响应")
			if resp.Payload != nil {
				t.Logf("响应包含有效载荷")
			}
		}
	})

	// 添加延迟避免频率限制
	time.Sleep(1 * time.Second)
}
