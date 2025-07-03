package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kenee/amazon-sp-api-sdk-go/apis/inventory"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func main() {
	// 配置认证信息
	credentials := &auth.LWAAuthorizationCredentials{
		ClientID:     os.Getenv("SP_API_CLIENT_ID"),
		ClientSecret: os.Getenv("SP_API_CLIENT_SECRET"),
		RefreshToken: os.Getenv("SP_API_REFRESH_TOKEN"),
		Endpoint:     os.Getenv("SP_API_ENDPOINT"),
	}

	// 创建配置
	config := client.NewConfigurationWithCredentials(credentials)
	config.SetHost(os.Getenv("SP_API_ENDPOINT_HOST"))

	// 创建 Inventory API 客户端
	inventoryAPI := inventory.NewInventoryAPI(config)

	// 创建上下文
	ctx := context.Background()

	// 示例 1: 获取库存摘要 (简化版本)
	fmt.Println("=== 获取库存摘要 (简化版本) ===")
	sellerId := os.Getenv("SP_API_SELLER_ID")
	if sellerId == "" {
		sellerId = "A13BCILWM2JF8S" // 默认值
	}

	marketplaceIds := []string{"ATVPDKIKX0DER"}

	response, err := inventoryAPI.GetInventorySummariesSimple(ctx, sellerId, marketplaceIds)
	if err != nil {
		log.Printf("获取库存摘要失败: %v", err)
	} else {
		fmt.Printf("✅ 库存摘要获取成功\n")
		fmt.Printf("   找到 %d 个库存项目\n", len(response.InventorySummaries))

		for i, summary := range response.InventorySummaries {
			fmt.Printf("\n   %d. ASIN: %s\n", i+1, summary.Asin)
			fmt.Printf("      SKU: %s\n", summary.SellerSku)
			fmt.Printf("      FnSKU: %s\n", summary.FnSku)
			fmt.Printf("      状态: %s\n", summary.Condition)
			fmt.Printf("      总数量: %d\n", summary.TotalQuantity)
			fmt.Printf("      可履行数量: %d\n", summary.InventoryDetails.FulfillableQuantity)
			fmt.Printf("      最后更新时间: %s\n", summary.LastUpdatedTime.Format("2006-01-02 15:04:05"))
		}
	}

	// 示例 2: 获取库存摘要 (完整版本)
	fmt.Println("\n=== 获取库存摘要 (完整版本) ===")
	request := &inventory.GetInventorySummariesRequest{
		SellerId:       sellerId,
		MarketplaceIds: marketplaceIds,
		Details:        true,
		StartDateTime:  "2024-01-01T00:00:00Z",
	}

	detailedResponse, err := inventoryAPI.GetInventorySummaries(ctx, request)
	if err != nil {
		log.Printf("获取详细库存摘要失败: %v", err)
	} else {
		fmt.Printf("✅ 详细库存摘要获取成功\n")
		fmt.Printf("   找到 %d 个库存项目\n", len(detailedResponse.InventorySummaries))

		for i, summary := range detailedResponse.InventorySummaries {
			fmt.Printf("\n   %d. ASIN: %s\n", i+1, summary.Asin)
			fmt.Printf("      产品名称: %s\n", summary.ProductName)
			fmt.Printf("      可履行数量: %d\n", summary.InventoryDetails.FulfillableQuantity)
			fmt.Printf("      入库中数量: %d\n", summary.InventoryDetails.InboundWorkingQuantity)
			fmt.Printf("      已发货数量: %d\n", summary.InventoryDetails.InboundShippedQuantity)
			fmt.Printf("      接收中数量: %d\n", summary.InventoryDetails.InboundReceivingQuantity)

			// 保留数量详情
			reserved := summary.InventoryDetails.ReservedQuantity
			fmt.Printf("      保留数量: %d\n", reserved.TotalReservedQuantity)
			fmt.Printf("        待客户订单: %d\n", reserved.PendingCustomerOrderQuantity)
			fmt.Printf("        待转运: %d\n", reserved.PendingTransshipmentQuantity)
			fmt.Printf("        处理中: %d\n", reserved.FcProcessingQuantity)

			// 不可履行数量详情
			unfulfillable := summary.InventoryDetails.UnfulfillableQuantity
			fmt.Printf("      不可履行数量: %d\n", unfulfillable.TotalUnfulfillableQuantity)
			fmt.Printf("        客户损坏: %d\n", unfulfillable.CustomerDamagedQuantity)
			fmt.Printf("        仓库损坏: %d\n", unfulfillable.WarehouseDamagedQuantity)
			fmt.Printf("        缺陷: %d\n", unfulfillable.DefectiveQuantity)
			fmt.Printf("        过期: %d\n", unfulfillable.ExpiredQuantity)
		}
	}
}
