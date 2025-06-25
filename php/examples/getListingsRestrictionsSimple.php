<?php

require_once '../sdk/vendor/autoload.php';

use SpApi\AuthAndAuth\LWAAuthorizationCredentials;
use SpApi\Configuration;
use SpApi\Api\Listings\Restrictions\v2021_08_01\ListingsApi;
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
    // Call getListingsRestrictions without conditionType
    $response = $listingsApi
        ->getListingsRestrictions(
            'B07N4M94X4', // asin - 商品的 ASIN
            'A1B2C3D4E5F6G7', // sellerId - 卖家标识符
            ['ATVPDKIKX0DER'] // marketplaceIds - 美国沙盒市场
        );

    // Process Listings Restrictions API response
    echo "Listings Restrictions API Response:\n";
    print_r($response);

    // Check if there are any restrictions
    $restrictions = $response->getRestrictions();
    if (!empty($restrictions)) {
        echo "\n⚠️ Restrictions found:\n";
        foreach ($restrictions as $restriction) {
            echo "Marketplace ID: " . $restriction->getMarketplaceId() . "\n";
            echo "Condition Type: " . $restriction->getConditionType() . "\n";
            
            $reasons = $restriction->getReasons();
            if (!empty($reasons)) {
                echo "Reasons:\n";
                foreach ($reasons as $reason) {
                    echo "- " . $reason->getMessage() . "\n";
                }
            }
            echo "\n";
        }
    } else {
        echo "\n✅ No restrictions found for this item.\n";
    }

} catch (Exception $e) {
    echo 'Exception when calling ListingsApi->getListingsRestrictions: ', $e->getMessage(), PHP_EOL;
} 