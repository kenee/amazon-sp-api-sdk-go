package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/reports"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestReportsListAPI(t *testing.T) {
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
		t.Skip("跳过 Reports API 测试：未找到 .env.sbx 文件")
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

	// 创建 Reports API 客户端
	reportsAPI := reports.NewReportsAPI(config)

	// 创建上下文
	ctx := context.Background()

	t.Run("获取报告列表", func(t *testing.T) {
		request := &reports.GetReportsRequest{
			ReportTypes:        []string{"GET_FLAT_FILE_OPEN_LISTINGS_DATA"},
			MarketplaceIds:     []string{"ATVPDKIKX0DER"},
			ProcessingStatuses: []string{"IN_QUEUE", "IN_PROGRESS", "DONE"},
			PageSize:           10,
		}

		response, err := reportsAPI.GetReports(ctx, request)
		if err != nil {
			t.Fatalf("获取报告列表失败: %v", err)
		}

		t.Logf("✅ 报告列表获取成功")
		t.Logf("   找到 %d 个报告", len(response.Reports))

		for i, report := range response.Reports {
			t.Logf("   %d. Report ID: %s, Type: %s, Status: %s",
				i+1, report.ReportId, report.ReportType, report.ProcessingStatus)
		}
	})
}
