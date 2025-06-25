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
    // Create request body for a simple luggage item
    $requestBody = new ListingsItemPutRequest([
        'productType' => 'LUGGAGE', // 沙盒环境支持的产品类型
        'requirements' => 'LISTING', // 需求集
        'attributes' => [
            'item_name' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'Hardside Carry-On Spinner Suitcase Luggage',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'brand' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'TestLuggageBrand',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'bullet_point' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'Hardside spinner luggage with TSA lock',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ],
                [
                    'language_tag' => 'en_US',
                    'value' => 'Expandable design for extra packing space',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ],
                [
                    'language_tag' => 'en_US',
                    'value' => 'Water-resistant polycarbonate shell',
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
                        'value' => 22.0
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
            ],
            'item_weight' => [
                [
                    'unit' => 'pounds',
                    'value' => 7.5,
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'material' => [
                [
                    'language_tag' => 'en_US',
                    'value' => 'Polycarbonate',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ],
            'warranty_description' => [
                [
                    'language_tag' => 'en_US',
                    'value' => '1 year limited warranty',
                    'marketplace_id' => 'ATVPDKIKX0DER'
                ]
            ]
        ]
    ]);

    // Call putListingsItem to create the item
    $response = $listingsApi
        ->putListingsItem(
            'A1B2C3D4E5F6G7', // sellerId - 卖家标识符
            'LUGGAGE-SKU-001', // sku - 卖家提供的商品标识符
            ['ATVPDKIKX0DER'], // marketplaceIds - 美国沙盒市场
            $requestBody, // 请求体
            ['issues', 'identifiers'], // includedData - 包含的数据集
            'CREATE' // mode - 操作模式：CREATE 用于实际创建商品
        );

    // Process Listings API response
    echo "Listings API Response:\n";
    print_r($response);

    // Check if the item was successfully created
    if ($response->getStatus() === 'ACCEPTED') {
        echo "\n✅ Item successfully created!\n";
        echo "SKU: " . $response->getSku() . "\n";
        echo "Submission ID: " . $response->getSubmissionId() . "\n";
        
        // Check for any issues
        $issues = $response->getIssues();
        if (!empty($issues)) {
            echo "\n⚠️ Issues found:\n";
            foreach ($issues as $issue) {
                echo "- " . $issue->getMessage() . "\n";
            }
        } else {
            echo "\n✅ No issues found.\n";
        }
    } else {
        echo "\n❌ Item creation failed with status: " . $response->getStatus() . "\n";
    }

} catch (Exception $e) {
    echo 'Exception when calling ListingsApi->putListingsItem: ', $e->getMessage(), PHP_EOL;
} 