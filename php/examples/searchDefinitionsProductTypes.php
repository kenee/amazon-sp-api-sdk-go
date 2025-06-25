<?php

require_once '../sdk/vendor/autoload.php';

use SpApi\AuthAndAuth\LWAAuthorizationCredentials;
use SpApi\Configuration;
use SpApi\Api\ProductTypeDefinitions\v2020_09_01\DefinitionsApi;
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

// Create an instance of the Definitions Api
$definitionsApi = new DefinitionsApi($config);

try {
    // Call searchDefinitionsProductTypes
    $response = $definitionsApi
    ->searchDefinitionsProductTypes(
        ['ATVPDKIKX0DER'], // marketplaceIds - 美国沙盒市场
        null, // keywords (optional) - 逗号分隔的关键词列表，不能与 itemName 同时使用
        null, // itemName (optional) - ASIN 标题，不能与 keywords 同时使用
        'en_US', // locale (optional) - 显示名称的语言环境
        'en_US' // searchLocale (optional) - 用于 keywords 和 itemName 参数的语言环境
    );

    // Process Product Type Search API response
    echo "Product Type Search API Response:\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling DefinitionsApi->searchDefinitionsProductTypes: ', $e->getMessage(), PHP_EOL;
} 