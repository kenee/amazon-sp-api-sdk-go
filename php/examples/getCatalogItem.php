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
    // Call getCatalogItem
    $response = $catalogApi->getCatalogItem(
        'B07N4M94X4', // ASIN - 沙盒测试 ASIN
        ['ATVPDKIKX0DER'], // MarketplaceIds - 美国沙盒
        ['summaries', 'attributes', 'images', 'classifications', 'dimensions', 'identifiers', 'productTypes', 'relationships', 'salesRanks', 'vendorDetails'], // includedData - 包含的数据集
        'en_US' // locale - 语言环境
    );

    // Process Catalog Item API response
    echo "Catalog Item API Response:\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling CatalogApi->getCatalogItem: ', $e->getMessage(), PHP_EOL;
} 