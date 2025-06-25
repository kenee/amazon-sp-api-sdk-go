<?php

require_once '../sdk/vendor/autoload.php';

use SpApi\AuthAndAuth\LWAAuthorizationCredentials;
use SpApi\Configuration;
use Dotenv\Dotenv;
use GuzzleHttp\Client;
use GuzzleHttp\Exception\RequestException;

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

try {
    // Create HTTP client
    $client = new Client();
    
    // Get access token
    $accessToken = $config->getAccessToken();
    
    // Prepare request parameters for SKU
    $params = [
        'marketplaceIds' => 'ATVPDKIKX0DER', // 美国沙盒市场
        'SellerSKU' => 'GM-ZDPI-9B4E', // 沙盒测试 SKU
        'sellerId' => 'A1B2C3D4E5F6G7' // 卖家ID（使用 SKU 时必需）
    ];
    
    // Build query string
    $queryString = http_build_query($params);
    
    // Make request to listCatalogCategories endpoint
    $response = $client->get($_ENV['SP_API_ENDPOINT_HOST'] . '/catalog/v0/categories?' . $queryString, [
        'headers' => [
            'Authorization' => 'Bearer ' . $accessToken,
            'Content-Type' => 'application/json',
            'Accept' => 'application/json'
        ]
    ]);
    
    // Get response body
    $body = $response->getBody()->getContents();
    $data = json_decode($body, true);
    
    // Process Catalog Categories API response
    echo "Catalog Categories API Response (SKU):\n";
    print_r($data);
    
} catch (RequestException $e) {
    echo 'Exception when calling listCatalogCategories: ', $e->getMessage(), PHP_EOL;
    if ($e->hasResponse()) {
        echo 'Response: ', $e->getResponse()->getBody()->getContents(), PHP_EOL;
    }
} catch (Exception $e) {
    echo 'Exception: ', $e->getMessage(), PHP_EOL;
} 