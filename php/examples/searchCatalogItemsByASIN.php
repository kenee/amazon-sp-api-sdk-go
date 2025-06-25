<?php

require_once '../sdk/vendor/autoload.php';

use SpApi\AuthAndAuth\LWAAuthorizationCredentials;
use SpApi\Configuration;
use SpApi\Api\CatalogItems\v2022_04_01\CatalogApi;
use Dotenv\Dotenv;

// Set the credentials, region and marketplace in .env file
$dotenv = Dotenv::createImmutable('../');
$dotenv->load();

// Set up LWA credentials
$lwaAuthorizationCredentials = new LWAAuthorizationCredentials([
    "clientId" => $_ENV['SP_API_CLIENT_ID'],
    "clientSecret" => $_ENV['SP_API_CLIENT_SECRET'],
    "refreshToken" => $_ENV['SP_API_REFRESH_TOKEN'],
    "endpoint" => $_ENV['SP_API_ENDPOINT']
]);

// Initialize config and set SP-API endpoint region
$config = new Configuration([], $lwaAuthorizationCredentials);
$config->setHost($_ENV['SP_API_ENDPOINT_HOST']);

// Create API instance
$catalogApi = new CatalogApi($config);

try {
    // Call searchCatalogItems with ASIN identifiers
    $response = $catalogApi->searchCatalogItems(
        ['ATVPDKIKX0DER'], // MarketplaceIds - 美国沙盒
        ['B07N4M94X4', 'B08J7TQ9FL'], // identifiers - ASIN 列表
        'ASIN', // identifiersType - 标识符类型为 ASIN
        ['summaries', 'attributes', 'images', 'classifications'], // includedData - 包含的数据集
        'en_US', // locale - 语言环境
        null, // sellerId - 卖家ID
        null, // keywords - 不使用关键词搜索
        null, // brandNames - 品牌名称
        null, // classificationIds - 分类ID
        20, // pageSize - 每页结果数
        null, // pageToken - 分页令牌
        null // keywordsLocale - 关键词语言环境
    );

    // Process Catalog Items API response
    echo "Catalog Items API Response (ASIN Search):\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling CatalogApi->searchCatalogItems: ', $e->getMessage(), PHP_EOL;
} 