<?php

require_once '../sdk/vendor/autoload.php';

use SpApi\AuthAndAuth\LWAAuthorizationCredentials;
use SpApi\Configuration;
use SpApi\Api\Listings\Items\v2021_08_01\ListingsApi;
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
$listingsApi = new ListingsApi($config);

try {
    // Call getListingsItem
    $response = $listingsApi->getListingsItem(
        'A1B2C3D4E5F6G7', // sellerId - 卖家标识符
        'GM-ZDPI-9B4E', // sku - 卖家提供的亚马逊商品标识符
        ['ATVPDKIKX0DER'], // marketplaceIds - 美国沙盒市场
        'en_US', // issueLocale - 问题本地化语言
        ['summaries', 'offers', 'fulfillmentAvailability', 'issues'] // includedData - 包含的数据集
    );

    // Process Listings Item API response
    echo "Listings Item API Response:\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling ListingsApi->getListingsItem: ', $e->getMessage(), PHP_EOL;
} 