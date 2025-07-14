package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/catalog"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestCatalogAPIDebug(t *testing.T) {
	// 1. 加载环境变量
	envPaths := []string{
		".env",
		"../../.env",
		"../../../.env",
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
		t.Skip("跳过 Catalog API Debug 测试：未找到 .env 文件")
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

	// 创建 Catalog API 客户端
	catalogAPI := catalog.NewCatalogAPI(config)

	// 创建上下文
	ctx := context.Background()

	t.Run("测试获取已知存在的产品信息", func(t *testing.T) {
		// 使用 PHP examples 中已知存在的 ASIN
		request := &catalog.GetCatalogItemRequest{
			ASIN:           "B071VG5N9D", // 来自 PHP examples
			MarketplaceIds: []string{"ATVPDKIKX0DER"},
		}

		response, err := catalogAPI.GetCatalogItem(ctx, request)
		if err != nil {
			t.Fatalf("获取产品信息失败: %v", err)
		}

		t.Logf("✅ 获取产品信息成功")
		t.Logf("   ASIN: %s", response.ASIN)
		t.Logf("   响应 ASIN: %s", response.ASIN)
	})

	t.Run("测试通过 ASIN 搜索产品", func(t *testing.T) {
		// 使用 PHP examples 中的 ASIN 列表
		request := &catalog.SearchCatalogItemsRequest{
			MarketplaceIds:  []string{"ATVPDKIKX0DER"},
			Identifiers:     []string{"B07N4M94X4", "B08J7TQ9FL"}, // 来自 PHP examples
			IdentifiersType: "ASIN",
			IncludedData:    []string{"summaries", "attributes", "images", "classifications"},
			Locale:          "en_US",
			PageSize:        20,
		}

		response, err := catalogAPI.SearchCatalogItems(ctx, request)
		if err != nil {
			t.Fatalf("通过 ASIN 搜索产品失败: %v", err)
		}

		t.Logf("✅ 通过 ASIN 搜索产品成功")
		t.Logf("   找到 %d 个产品", len(response.Items))
		for i, item := range response.Items {
			t.Logf("   %d. ASIN: %s", i+1, item.ASIN)
		}
	})

	t.Run("测试关键词搜索", func(t *testing.T) {
		// 使用 PHP examples 中的简单关键词
		request := &catalog.SearchCatalogItemsRequest{
			MarketplaceIds: []string{"ATVPDKIKX0DER"},
			Keywords:       []string{"book"}, // 来自 PHP examples
			PageSize:       1,
		}

		response, err := catalogAPI.SearchCatalogItems(ctx, request)
		if err != nil {
			t.Fatalf("关键词搜索失败: %v", err)
		}

		t.Logf("✅ 关键词搜索成功")
		t.Logf("   找到 %d 个产品", len(response.Items))
		for i, item := range response.Items {
			t.Logf("   %d. ASIN: %s", i+1, item.ASIN)
		}
	})

	t.Run("测试最简单的搜索参数", func(t *testing.T) {
		// 使用最简单的参数，只包含必需的 marketplaceIds
		request := &catalog.SearchCatalogItemsRequest{
			MarketplaceIds: []string{"ATVPDKIKX0DER"},
		}

		response, err := catalogAPI.SearchCatalogItems(ctx, request)
		if err != nil {
			t.Fatalf("搜索产品失败: %v", err)
		}

		t.Logf("✅ 简单搜索成功")
		t.Logf("   找到 %d 个产品", len(response.Items))
	})

	t.Run("测试 PHP SDK 相同的参数", func(t *testing.T) {
		// 使用与 PHP SDK 示例完全相同的参数
		request := &catalog.SearchCatalogItemsRequest{
			MarketplaceIds: []string{"ATVPDKIKX0DER"},
			IncludedData:   []string{"summaries"},
			Locale:         "en_US",
			PageSize:       10,
		}

		response, err := catalogAPI.SearchCatalogItems(ctx, request)
		if err != nil {
			t.Fatalf("搜索产品失败: %v", err)
		}

		t.Logf("✅ PHP SDK 参数搜索成功")
		t.Logf("   找到 %d 个产品", len(response.Items))
	})
}
