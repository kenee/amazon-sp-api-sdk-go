package production

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kenee/amazon-sp-api-sdk-go/apis/feeds"
	"github.com/kenee/amazon-sp-api-sdk-go/auth"
	"github.com/kenee/amazon-sp-api-sdk-go/client"
)

func TestFeedsAPIProduction(t *testing.T) {
	// 1. 加载环境变量
	envPaths := []string{
		".env.prod",          // 当前目录
		"../.env.prod",       // 上级目录
		"../../.env.prod",    // 上上级目录
		"../../../.env.prod", // 根目录
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
		t.Skip("跳过 Feeds API 生产环境测试：未找到 .env 或 .env.prod1 文件")
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

	// 创建 Feeds API 客户端
	feedsAPI := feeds.NewFeedsAPI(config)

	// 创建上下文
	ctx := context.Background()

	t.Run("创建 Feed Document", func(t *testing.T) {
		request := &feeds.CreateFeedDocumentRequest{
			ContentType: "text/tab-separated-values; charset=UTF-8",
		}

		response, err := feedsAPI.CreateFeedDocument(ctx, request)
		if err != nil {
			t.Fatalf("创建 Feed Document 失败: %v", err)
		}

		if response.FeedDocumentId == "" {
			t.Error("Feed Document ID 为空")
		}

		if response.Url == "" {
			t.Error("Feed Document URL 为空")
		}

		t.Logf("✅ Feed Document 创建成功")
		t.Logf("   Document ID: %s", response.FeedDocumentId)
		t.Logf("   URL: %s", response.Url)
	})

	t.Run("获取 Feed 列表", func(t *testing.T) {
		// 使用与PHP生产环境示例一致的参数
		request := &feeds.GetFeedsRequest{
			FeedTypes:      []string{"POST_PRODUCT_DATA"}, // Feed类型
			MarketplaceIds: []string{"ATVPDKIKX0DER"},     // 美国生产环境
			// 其他参数保持null，与PHP示例一致
		}

		response, err := feedsAPI.GetFeeds(ctx, request)
		if err != nil {
			t.Fatalf("获取 Feed 列表失败: %v", err)
		}

		t.Logf("✅ Feed 列表获取成功")
		t.Logf("   找到 %d 个 Feed", len(response.Feeds))

		for i, feed := range response.Feeds {
			t.Logf("   %d. Feed ID: %s, Status: %s", i+1, feed.FeedId, feed.ProcessingStatus)
		}
	})

	t.Run("创建Feed", func(t *testing.T) {
		// 使用与PHP生产环境示例一致的参数
		request := &feeds.CreateFeedRequest{
			FeedType:            "POST_PRODUCT_DATA",       // Feed类型
			MarketplaceIds:      []string{"ATVPDKIKX0DER"}, // 美国生产环境
			InputFeedDocumentId: "test-document-id",        // 输入文档ID（测试用）
			// FeedOptions保持null，与PHP示例一致
		}

		response, err := feedsAPI.CreateFeed(ctx, request)
		if err != nil {
			t.Fatalf("创建Feed失败: %v", err)
		}

		t.Logf("✅ Feed创建成功")
		t.Logf("   Feed ID: %s", response.FeedId)
	})
}

func TestFeedEndToEndProduction(t *testing.T) {
	// 1. 加载环境变量
	envPaths := []string{
		".env.prod1",
		"../.env.prod1",
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
		t.Skip("跳过 Feeds API 生产环境测试：未找到 .env 或 .env.prod1 文件")
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
	feedsAPI := feeds.NewFeedsAPI(config)
	ctx := context.Background()

	t.Log("\n🚀 开始 Feed End-to-End 测试 (生产环境)")
	t.Log("====================================\n")

	// 步骤1: 创建 Feed Document
	t.Log("📝 步骤1: 创建 Feed Document")
	docReq := &feeds.CreateFeedDocumentRequest{
		ContentType: "text/tab-separated-values; charset=UTF-8",
	}
	docResp, err := feedsAPI.CreateFeedDocument(ctx, docReq)
	if err != nil {
		t.Fatalf("创建 Feed Document 失败: %v", err)
	}
	if docResp.FeedDocumentId == "" || docResp.Url == "" {
		t.Fatalf("Feed Document ID 或 URL 为空")
	}
	t.Logf("✅ Feed Document 创建成功: %s", docResp.FeedDocumentId)
	t.Logf("   Upload URL: %s...", docResp.Url[:80])
	t.Log("   说明: 预签名URL有效期5分钟，Amazon自动清理临时文件，无需担心费用。\n")

	// 步骤2: 上传 Feed 文件内容到S3
	t.Log("📤 步骤2: 上传 Feed 文件内容到S3 (PUT)")
	feedContent := "sku\tproduct-id\tproduct-id-type\tprice\tminimum-seller-allowed-price\tmaximum-seller-allowed-price\titem-condition\tquantity\tshipping-template-name\n"
	feedContent += "TEST-SKU-001\tB07FZ8S74R\tASIN\t29.99\t25.00\t35.00\tnew\t10\tStandard\n"
	req, err := http.NewRequestWithContext(ctx, "PUT", docResp.Url, bytes.NewReader([]byte(feedContent)))
	if err != nil {
		t.Fatalf("构造上传请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "text/tab-separated-values; charset=UTF-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("上传Feed文件失败: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("Feed文件上传失败，HTTP状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}
	t.Logf("✅ Feed 文件上传成功 (HTTP %d)", resp.StatusCode)
	t.Log("   说明: 必须在5分钟内上传，Content-Type需与创建时一致。\n")

	// 步骤3: 创建 Feed
	t.Log("📋 步骤3: 创建 Feed")
	createReq := &feeds.CreateFeedRequest{
		FeedType:            "POST_PRODUCT_DATA",
		MarketplaceIds:      []string{"ATVPDKIKX0DER"},
		InputFeedDocumentId: docResp.FeedDocumentId,
	}
	createResp, err := feedsAPI.CreateFeed(ctx, createReq)
	if err != nil {
		t.Fatalf("创建Feed失败: %v", err)
	}
	if createResp.FeedId == "" {
		t.Fatalf("Feed ID 为空")
	}
	t.Logf("✅ Feed 创建成功: %s", createResp.FeedId)

	// 步骤4: 获取 Feed 详情
	t.Log("📊 步骤4: 获取 Feed 详情")
	feedInfo, err := feedsAPI.GetFeed(ctx, createResp.FeedId)
	if err != nil {
		t.Fatalf("获取Feed详情失败: %v", err)
	}
	if feedInfo.FeedId == "" {
		t.Fatalf("Feed详情返回的FeedId为空")
	}
	t.Logf("✅ Feed 详情获取成功: %s, 状态: %s", feedInfo.FeedId, feedInfo.ProcessingStatus)

	// 步骤5: 获取 Feed 列表
	t.Log("📋 步骤5: 获取 Feed 列表")
	feedsList, err := feedsAPI.GetFeeds(ctx, &feeds.GetFeedsRequest{
		FeedTypes:      []string{"POST_PRODUCT_DATA"},
		MarketplaceIds: []string{"ATVPDKIKX0DER"},
		PageSize:       10,
	})
	if err != nil {
		t.Fatalf("获取Feed列表失败: %v", err)
	}
	t.Logf("✅ Feed 列表获取成功, 共%d个", len(feedsList.Feeds))
	for i, f := range feedsList.Feeds {
		t.Logf("   %d. Feed ID: %s, 状态: %s, 创建时间: %s", i+1, f.FeedId, f.ProcessingStatus, f.CreatedTime.Format("2006-01-02 15:04:05"))
	}

	t.Log("\n🎉 Feed End-to-End 测试完成！\n")
	// 费用说明
	t.Log("💰 费用说明：Amazon自动处理S3临时存储和清理，不会产生额外费用。\n")
}
