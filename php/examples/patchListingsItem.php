<?php

require_once '../sdk/vendor/autoload.php';

use SpApi\AuthAndAuth\LWAAuthorizationCredentials;
use SpApi\Configuration;
use SpApi\Api\Listings\Items\v2021_08_01\ListingsApi;
use SpApi\Model\Listings\Items\v2021_08_01\ListingsItemPatchRequest;
use SpApi\Model\Listings\Items\v2021_08_01\PatchOperation;
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
    // Create patch operations
    $patches = [];
    
    // Example 1: Replace the item name
    $patch1 = new PatchOperation();
    $patch1->setOp('replace');
    $patch1->setPath('/attributes/item_name');
    $patch1->setValue([
        [
            'marketplace_id' => 'ATVPDKIKX0DER',
            'language_tag' => 'en_US',
            'value' => 'Updated Hardside Carry-On Spinner Suitcase Luggage'
        ]
    ]);
    $patches[] = $patch1;
    
    // Example 2: Add a new attribute
    $patch2 = new PatchOperation();
    $patch2->setOp('add');
    $patch2->setPath('/attributes/brand');
    $patch2->setValue([
        [
            'marketplace_id' => 'ATVPDKIKX0DER',
            'language_tag' => 'en_US',
            'value' => 'TravelPro'
        ]
    ]);
    $patches[] = $patch2;
    
    // Example 3: Update price
    $patch3 = new PatchOperation();
    $patch3->setOp('replace');
    $patch3->setPath('/offers/0/price/amount');
    $patch3->setValue([
        [
            'marketplace_id' => 'ATVPDKIKX0DER',
            'value' => '89.99'
        ]
    ]);
    $patches[] = $patch3;

    // Create patch request
    $patchRequest = new ListingsItemPatchRequest();
    $patchRequest->setProductType('LUGGAGE');
    $patchRequest->setPatches($patches);

    // Call patchListingsItem
    $response = $listingsApi->patchListingsItem(
        'A1B2C3D4E5F6G7', // sellerId
        'LUGGAGE-SKU-001', // sku
        ['ATVPDKIKX0DER'], // marketplaceIds
        $patchRequest, // body
        ['issues'], // includedData
        'VALIDATION_PREVIEW' // mode
    );

    // Process Listings API response
    echo "Patch Listings Item API Response:\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling ListingsApi->patchListingsItem: ', $e->getMessage(), PHP_EOL;
} 