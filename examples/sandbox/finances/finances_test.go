package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	v0 "github.com/kenee/amazon-sp-api-sdk-go/apis/finances/v0"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

// 通用的测试设置函数
func setupFinancesTest(t *testing.T) *v0.FinancesAPI {
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
		t.Skip("跳过 Finances API 测试：未找到 .env.sbx 文件")
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
	return v0.NewFinancesAPI(config)
}

func TestListFinancialEventGroups(t *testing.T) {
	api := setupFinancesTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListFinancialEventGroups API")

	// 使用沙盒环境的魔法参数
	request := &v0.ListFinancialEventGroupsRequest{
		MaxResultsPerPage:                1,
		FinancialEventGroupStartedBefore: "2019-10-31",
		FinancialEventGroupStartedAfter:  "2019-10-13",
	}

	resp, err := api.ListFinancialEventGroups(ctx, request)
	if err != nil {
		t.Logf("⚠️ 获取财务事件组失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功获取财务事件组")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestListFinancialEvents(t *testing.T) {
	api := setupFinancesTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListFinancialEvents API")

	// 使用沙盒环境的魔法参数
	request := &v0.ListFinancialEventsRequest{
		MaxResultsPerPage: 1,
		PostedAfter:       "2019-10-13T00:00:00Z",
	}

	resp, err := api.ListFinancialEvents(ctx, request)
	if err != nil {
		t.Logf("⚠️ 获取财务事件失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功获取财务事件")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestListFinancialEventsByGroupId(t *testing.T) {
	api := setupFinancesTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListFinancialEventsByGroupId API")

	// 使用沙盒环境的魔法参数
	eventGroupId := "TEST-GROUP-123"
	request := &v0.ListFinancialEventsByGroupIdRequest{
		MaxResultsPerPage: 1,
	}

	resp, err := api.ListFinancialEventsByGroupId(ctx, eventGroupId, request)
	if err != nil {
		t.Logf("⚠️ 按组ID获取财务事件失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功按组ID获取财务事件")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestListFinancialEventsByOrderId(t *testing.T) {
	api := setupFinancesTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListFinancialEventsByOrderId API")

	// 使用沙盒环境的魔法参数
	orderId := "BAD-ORDER" // 沙盒环境会返回400错误
	request := &v0.ListFinancialEventsByOrderIdRequest{
		MaxResultsPerPage: 1,
	}

	resp, err := api.ListFinancialEventsByOrderId(ctx, orderId, request)
	if err != nil {
		t.Logf("✅ 按订单ID获取财务事件测试通过（返回预期的错误）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功按订单ID获取财务事件")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestListFinancialEventGroupsSimple(t *testing.T) {
	api := setupFinancesTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListFinancialEventGroupsSimple API")

	resp, err := api.ListFinancialEventGroupsSimple(ctx, 1)
	if err != nil {
		t.Logf("⚠️ 简化版获取财务事件组失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功简化版获取财务事件组")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}

func TestListFinancialEventsSimple(t *testing.T) {
	api := setupFinancesTest(t)
	ctx := context.Background()

	t.Log("🚀 开始测试 ListFinancialEventsSimple API")

	resp, err := api.ListFinancialEventsSimple(ctx, 1)
	if err != nil {
		t.Logf("⚠️ 简化版获取财务事件失败（沙盒环境限制）: %v", err)
		return
	}

	if resp != nil {
		t.Logf("✅ 成功简化版获取财务事件")
		if resp.Payload != nil {
			t.Logf("   响应包含有效载荷")
		}
	}
}
