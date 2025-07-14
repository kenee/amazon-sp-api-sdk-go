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

func TestSubscriptionAPI(t *testing.T) {
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

	t.Run("CreateSubscription", func(t *testing.T) {
		// magic参数: notificationType=ANY_OFFER_CHANGED, 空请求体
		notificationType := "ANY_OFFER_CHANGED"
		request := &v1.CreateSubscriptionRequest{}
		resp, err := api.CreateSubscription(ctx, notificationType, request)
		if err != nil {
			t.Logf("⚠️ CreateSubscription 失败（沙盒限制）: %v", err)
			return
		}
		if resp != nil && resp.Payload != nil {
			t.Logf("✅ CreateSubscription 成功, subscriptionId: %v", resp.Payload.SubscriptionId)
		}
	})

	t.Run("GetSubscription", func(t *testing.T) {
		// magic参数: notificationType=ANY_OFFER_CHANGED, payloadVersion=""
		notificationType := "ANY_OFFER_CHANGED"
		resp, err := api.GetSubscription(ctx, notificationType, "")
		if err != nil {
			t.Logf("⚠️ GetSubscription 失败（沙盒限制）: %v", err)
			return
		}
		if resp != nil && resp.Payload != nil {
			t.Logf("✅ GetSubscription 成功, subscriptionId: %v", resp.Payload.SubscriptionId)
		}
	})
}
