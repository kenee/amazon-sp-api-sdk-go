package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kenee/amazon-sp-api-sdk-go/apis/feeds"
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

	// 创建 Feeds API 客户端
	feedsAPI := feeds.NewFeedsAPI(config)

	// 创建上下文
	ctx := context.Background()

	// 示例 1: 创建 Feed Document
	fmt.Println("=== 创建 Feed Document ===")
	createDocRequest := &feeds.CreateFeedDocumentRequest{
		ContentType: "text/tab-separated-values; charset=UTF-8",
	}

	docResponse, err := feedsAPI.CreateFeedDocument(ctx, createDocRequest)
	if err != nil {
		log.Printf("创建 Feed Document 失败: %v", err)
	} else {
		fmt.Printf("✅ Feed Document 创建成功\n")
		fmt.Printf("   Document ID: %s\n", docResponse.FeedDocumentId)
		fmt.Printf("   URL: %s\n", docResponse.Url)
	}

	// 示例 2: 创建 Feed
	fmt.Println("\n=== 创建 Feed ===")
	createFeedRequest := &feeds.CreateFeedRequest{
		FeedType:            "POST_PRODUCT_DATA",
		MarketplaceIds:      []string{"ATVPDKIKX0DER"},
		InputFeedDocumentId: docResponse.FeedDocumentId,
	}

	feedResponse, err := feedsAPI.CreateFeed(ctx, createFeedRequest)
	if err != nil {
		log.Printf("创建 Feed 失败: %v", err)
	} else {
		fmt.Printf("✅ Feed 创建成功\n")
		fmt.Printf("   Feed ID: %s\n", feedResponse.FeedId)
	}

	// 示例 3: 获取 Feed 信息
	fmt.Println("\n=== 获取 Feed 信息 ===")
	if feedResponse != nil {
		feedInfo, err := feedsAPI.GetFeed(ctx, feedResponse.FeedId)
		if err != nil {
			log.Printf("获取 Feed 信息失败: %v", err)
		} else {
			fmt.Printf("✅ Feed 信息获取成功\n")
			fmt.Printf("   Feed ID: %s\n", feedInfo.FeedId)
			fmt.Printf("   Feed Type: %s\n", feedInfo.FeedType)
			fmt.Printf("   Processing Status: %s\n", feedInfo.ProcessingStatus)
			fmt.Printf("   Created Time: %s\n", feedInfo.CreatedTime.Format(time.RFC3339))
		}
	}

	// 示例 4: 获取 Feed 列表
	fmt.Println("\n=== 获取 Feed 列表 ===")
	getFeedsRequest := &feeds.GetFeedsRequest{
		FeedTypes:          []string{"POST_PRODUCT_DATA"},
		ProcessingStatuses: []string{"IN_QUEUE", "IN_PROGRESS", "DONE"},
		PageSize:           10,
	}

	feedsList, err := feedsAPI.GetFeeds(ctx, getFeedsRequest)
	if err != nil {
		log.Printf("获取 Feed 列表失败: %v", err)
	} else {
		fmt.Printf("✅ Feed 列表获取成功\n")
		fmt.Printf("   找到 %d 个 Feed\n", len(feedsList.Feeds))
		for i, feed := range feedsList.Feeds {
			fmt.Printf("   %d. Feed ID: %s, Status: %s\n", i+1, feed.FeedId, feed.ProcessingStatus)
		}
	}
}
