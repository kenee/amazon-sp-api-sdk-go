package sandbox

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/orders"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestOrdersRawAPI(t *testing.T) {
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

	t.Run("沙盒环境 - 获取原始响应", func(t *testing.T) {
		// 尝试不同的沙盒参数组合
		testCases := []struct {
			name    string
			request *orders.GetOrdersRequest
		}{
			{
				name: "TEST_CASE_200",
				request: &orders.GetOrdersRequest{
					MarketplaceIds: []string{"ATVPDKIKX0DER"},
					CreatedAfter:   "TEST_CASE_200",
				},
			},
			{
				name: "标准日期格式",
				request: &orders.GetOrdersRequest{
					MarketplaceIds: []string{"ATVPDKIKX0DER"},
					CreatedAfter:   "2023-01-01T00:00:00Z",
				},
			},
			{
				name: "只传marketplaceIds",
				request: &orders.GetOrdersRequest{
					MarketplaceIds: []string{"ATVPDKIKX0DER"},
				},
			},
			{
				name: "使用不同的marketplaceId",
				request: &orders.GetOrdersRequest{
					MarketplaceIds: []string{"A1PA6795UKMFR9"}, // 德国
					CreatedAfter:   "2023-01-01T00:00:00Z",
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// 构建查询参数
				params := make(map[string]interface{})
				if len(tc.request.MarketplaceIds) > 0 {
					params["marketplaceIds"] = tc.request.MarketplaceIds
				}
				if tc.request.CreatedAfter != "" {
					params["createdAfter"] = tc.request.CreatedAfter
				}

				// 构建查询字符串
				queryString := ordersAPI.GetAPIClient().BuildQueryString(params)

				// 构建完整路径
				path := "/orders/v0/orders"
				if queryString != "" {
					path += "?" + queryString
				}

				t.Logf("🔍 测试用例: %s", tc.name)
				t.Logf("🔍 请求路径: %s", path)

				// 直接调用API
				resp, err := ordersAPI.GetAPIClient().CallAPI(ctx, "GET", path, nil, nil)
				if err != nil {
					t.Logf("❌ API调用失败: %v", err)
					return
				}
				defer resp.Body.Close()

				// 读取响应体
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Logf("❌ 读取响应体失败: %v", err)
					return
				}

				t.Logf("🔍 响应状态码: %d", resp.StatusCode)
				t.Logf("🔍 原始响应体:")
				t.Logf("%s", string(body))

				// 尝试格式化JSON
				var prettyJSON interface{}
				if err := json.Unmarshal(body, &prettyJSON); err == nil {
					formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
					t.Logf("🔍 格式化后的JSON:")
					t.Logf("%s", string(formatted))
				}

				// 检查是否是成功响应
				if resp.StatusCode >= 200 && resp.StatusCode < 300 {
					t.Logf("✅ API调用成功，状态码: %d", resp.StatusCode)
				} else {
					t.Logf("⚠️  API调用返回错误状态码: %d", resp.StatusCode)
				}
				t.Logf("---")
			})
		}
	})
}
