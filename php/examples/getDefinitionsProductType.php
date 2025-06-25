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
    // Call getDefinitionsProductType
    $response = $definitionsApi
    ->getDefinitionsProductType(
        'LUGGAGE', // productType - 沙盒测试值
        ['ATVPDKIKX0DER'], // marketplaceIds - 美国沙盒市场
        null, // sellerId (optional)
        'LATEST', // productTypeVersion (optional, default: LATEST)
        'LISTING', // requirements (optional, default: LISTING)
        'ENFORCED', // requirementsEnforced (optional, default: ENFORCED)
        'en_US' // locale (optional, default: DEFAULT)
    );

    // Process Product Type Definition API response
    echo "Product Type Definition API Response:\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling DefinitionsApi->getDefinitionsProductType: ', $e->getMessage(), PHP_EOL;
} 