<?php

require_once '../sdk/vendor/autoload.php';

use SpApi\AuthAndAuth\LWAAuthorizationCredentials;
use SpApi\Configuration;
use SpApi\Api\Listings\Items\v2021_08_01\ListingsApi;
use SpApi\Model\Listings\Items\v2021_08_01\ListingsItemPutRequest;
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
    // Create request body
    $requestBody = new ListingsItemPutRequest([
        'productType' => 'LUGGAGE', // 沙盒环境支持的产品类型
        'requirements' => 'LISTING', // 需求集：LISTING, LISTING_PRODUCT_ONLY, LISTING_OFFER_ONLY
        'attributes' => [
            'item_name' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'Test Luggage Item',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'brand' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'TestBrand',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'bullet_point' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'Test bullet point 1',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ],
                [
                    'language_tag' => 'en_US',
                    'value' => 'Test bullet point 2',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'color' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'Black',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'item_dimensions' => [
                [
                    'width' => [
                        'unit' => 'inches',
                        'value' => 20.0
                    ],
                    'length' => [
                        'unit' => 'inches',
                        'value' => 14.0
                    ],
                    'height' => [
                        'unit' => 'inches',
                        'value' => 9.0
                    ],
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ]
        ]
    ]);

    // Call putListingsItem
    $response = $listingsApi
        ->putListingsItem(
            'A1B2C3D4E5F6G7', // sellerId - 卖家标识符
            'TEST-SKU-001', // sku - 卖家提供的商品标识符
            ['ATVPDKIKX0DER'], // marketplaceIds - 美国沙盒市场
            $requestBody, // 请求体
            ['issues'], // includedData - 包含的数据集
            'VALIDATION_PREVIEW' // mode - 操作模式：VALIDATION_PREVIEW 用于验证
        );

    // Process Listings API response
    echo "Listings API Response:\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling ListingsApi->putListingsItem: ', $e->getMessage(), PHP_EOL;
} 