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

func TestCatalogAPI(t *testing.T) {
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
		t.Skip("跳过 Catalog API 测试：未找到 .env.sbx 文件")
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

	t.Run("获取产品信息 (简化版本)", func(t *testing.T) {
		asin := "B08N5WRWNW" // 测试 ASIN
		marketplaceIds := []string{"ATVPDKIKX0DER"}

		response, err := catalogAPI.GetCatalogItemSimple(ctx, asin, marketplaceIds)
		if err != nil {
			t.Fatalf("获取产品信息失败: %v", err)
		}

		t.Logf("✅ 产品信息获取成功")
		t.Logf("   ASIN: %s", asin)
		t.Logf("   响应 ASIN: %s", response.ASIN)

		// 验证响应
		if response.ASIN == "" {
			t.Error("响应 ASIN 为空")
		}
	})

	t.Run("搜索产品", func(t *testing.T) {
		request := &catalog.SearchCatalogItemsRequest{
			MarketplaceIds: []string{"ATVPDKIKX0DER"},
			Keywords:       []string{"laptop"},
			IncludedData:   []string{"summaries", "attributes"},
			Locale:         "en_US",
			PageSize:       5,
			KeywordsLocale: "en_US",
		}

		response, err := catalogAPI.SearchCatalogItems(ctx, request)
		if err != nil {
			t.Fatalf("搜索产品失败: %v", err)
		}

		t.Logf("✅ 产品搜索成功")
		t.Logf("   找到 %d 个产品", len(response.Items))

		for i, item := range response.Items {
			t.Logf("   %d. ASIN: %s", i+1, item.ASIN)
			if len(item.Summaries) > 0 {
				t.Logf("      标题: %s", item.Summaries[0].ItemName)
			}
		}

		// 验证响应
		if len(response.Items) == 0 {
			t.Log("未找到产品，这可能是正常的")
		}
	})

	t.Run("获取产品信息 (完整版本)", func(t *testing.T) {
		request := &catalog.GetCatalogItemRequest{
			ASIN:           "B08N5WRWNW",
			MarketplaceIds: []string{"ATVPDKIKX0DER"},
			IncludedData:   []string{"summaries", "attributes", "images", "classifications"},
			Locale:         "en_US",
		}

		response, err := catalogAPI.GetCatalogItem(ctx, request)
		if err != nil {
			t.Fatalf("获取完整产品信息失败: %v", err)
		}

		t.Logf("✅ 完整产品信息获取成功")
		t.Logf("   ASIN: %s", request.ASIN)
		t.Logf("   响应 ASIN: %s", response.ASIN)

		// 验证响应
		if response.ASIN == "" {
			t.Error("响应 ASIN 为空")
		}
	})
}
