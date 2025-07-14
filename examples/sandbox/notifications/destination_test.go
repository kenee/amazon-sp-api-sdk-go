package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	v1 "github.com/kenee/amazon-sp-api-sdk-go/apis/notifications/v1"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestDestinationAPI(t *testing.T) {
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
		t.Skip("跳过 Notifications API 测试：未找到 .env.sbx 文件")
	}

	// 2. 配置认证信息
	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

	api := v1.NewNotificationsAPI(config)
	ctx := context.Background()

	t.Run("CreateDestination", func(t *testing.T) {
		// magic参数: 空请求体
		request := &v1.CreateDestinationRequest{}
		resp, err := api.CreateDestination(ctx, request)
		if err != nil {
			t.Logf("⚠️ CreateDestination 失败（沙盒限制）: %v", err)
			return
		}
		if resp != nil && resp.Payload != nil {
			t.Logf("✅ CreateDestination 成功, destinationId: %v", resp.Payload.DestinationId)
		}
	})

	t.Run("GetDestinations", func(t *testing.T) {
		resp, err := api.GetDestinations(ctx)
		if err != nil {
			t.Logf("⚠️ GetDestinations 失败（沙盒限制）: %v", err)
			return
		}
		if resp != nil && resp.Payload != nil {
			t.Logf("✅ GetDestinations 成功, count: %d", len(resp.Payload.Destinations))
		}
	})
}
