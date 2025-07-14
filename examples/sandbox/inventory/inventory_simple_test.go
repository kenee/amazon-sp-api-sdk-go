package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/inventory"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestInventorySimpleAPI(t *testing.T) {
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
		t.Skip("跳过库存 API 测试：未找到 .env.sbx 文件")
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

	// 4. 创建 Inventory API 客户端
	inventoryAPI := inventory.NewInventoryAPI(config)

	// 5. 创建上下文
	ctx := context.Background()

	// 6. 获取 Seller ID
	sellerId := os.Getenv("SP_API_SELLER_ID")
	if sellerId == "" {
		sellerId = "A13BCILWM2JF8S" // 默认值
	}

	marketplaceIds := []string{"ATVPDKIKX0DER"}

	t.Run("获取库存摘要 (简化版本)", func(t *testing.T) {
		response, err := inventoryAPI.GetInventorySummariesSimple(ctx, sellerId, marketplaceIds)
		if err != nil {
			t.Fatalf("获取库存摘要失败: %v", err)
		}

		t.Logf("✅ 库存摘要获取成功")
		t.Logf("   找到 %d 个库存项目", len(response.InventorySummaries))

		for i, summary := range response.InventorySummaries {
			t.Logf("\n   %d. ASIN: %s", i+1, summary.Asin)
			t.Logf("      SKU: %s", summary.SellerSku)
			t.Logf("      FnSKU: %s", summary.FnSku)
			t.Logf("      状态: %s", summary.Condition)
			t.Logf("      总数量: %d", summary.TotalQuantity)
			t.Logf("      可履行数量: %d", summary.InventoryDetails.FulfillableQuantity)
			t.Logf("      最后更新时间: %s", summary.LastUpdatedTime.Format("2006-01-02 15:04:05"))

			// 验证必要字段
			if summary.Asin == "" {
				t.Error("ASIN 为空")
			}
			if summary.TotalQuantity < 0 {
				t.Error("总数量不能为负数")
			}
			if summary.InventoryDetails.FulfillableQuantity < 0 {
				t.Error("可履行数量不能为负数")
			}
		}
	})
}
