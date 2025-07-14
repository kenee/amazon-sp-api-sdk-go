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

func TestInventoryAPI(t *testing.T) {
	// 1. 加载沙盒环境配置
	envPaths := []string{
		".env.sbx",
		"../.env.sbx",
		"../../.env.sbx",
		"../../../.env.sbx",
	}

	envLoaded := false
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			t.Logf("✅ 成功加载沙盒环境环境变量文件: %s", path)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		t.Skip("跳过 Inventory API 测试：未找到 .env.sbx 文件")
		return
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

	t.Run("获取库存摘要 (完整版本)", func(t *testing.T) {
		request := &inventory.GetInventorySummariesRequest{
			GranularityType: "Marketplace",
			GranularityId:   marketplaceIds[0],
			SellerId:        sellerId,
			MarketplaceIds:  marketplaceIds,
			Details:         true,
			StartDateTime:   "2024-01-01T00:00:00Z",
		}

		response, err := inventoryAPI.GetInventorySummaries(ctx, request)
		if err != nil {
			t.Fatalf("获取详细库存摘要失败: %v", err)
		}

		t.Logf("✅ 详细库存摘要获取成功")
		t.Logf("   找到 %d 个库存项目", len(response.InventorySummaries))

		for i, summary := range response.InventorySummaries {
			t.Logf("\n   %d. ASIN: %s", i+1, summary.Asin)
			t.Logf("      产品名称: %s", summary.ProductName)
			t.Logf("      可履行数量: %d", summary.InventoryDetails.FulfillableQuantity)
			t.Logf("      入库中数量: %d", summary.InventoryDetails.InboundWorkingQuantity)
			t.Logf("      已发货数量: %d", summary.InventoryDetails.InboundShippedQuantity)
			t.Logf("      接收中数量: %d", summary.InventoryDetails.InboundReceivingQuantity)

			// 保留数量详情
			reserved := summary.InventoryDetails.ReservedQuantity
			t.Logf("      保留数量: %d", reserved.TotalReservedQuantity)
			t.Logf("        待客户订单: %d", reserved.PendingCustomerOrderQuantity)
			t.Logf("        待转运: %d", reserved.PendingTransshipmentQuantity)
			t.Logf("        处理中: %d", reserved.FcProcessingQuantity)

			// 不可履行数量详情
			unfulfillable := summary.InventoryDetails.UnfulfillableQuantity
			t.Logf("      不可履行数量: %d", unfulfillable.TotalUnfulfillableQuantity)
			t.Logf("        客户损坏: %d", unfulfillable.CustomerDamagedQuantity)
			t.Logf("        仓库损坏: %d", unfulfillable.WarehouseDamagedQuantity)
			t.Logf("        缺陷: %d", unfulfillable.DefectiveQuantity)
			t.Logf("        过期: %d", unfulfillable.ExpiredQuantity)

			// 验证数据完整性
			if summary.Asin == "" {
				t.Error("ASIN 为空")
			}
			if summary.TotalQuantity < 0 {
				t.Error("总数量不能为负数")
			}
			if summary.InventoryDetails.FulfillableQuantity < 0 {
				t.Error("可履行数量不能为负数")
			}
			if reserved.TotalReservedQuantity < 0 {
				t.Error("保留数量不能为负数")
			}
			if unfulfillable.TotalUnfulfillableQuantity < 0 {
				t.Error("不可履行数量不能为负数")
			}
		}
	})
}
