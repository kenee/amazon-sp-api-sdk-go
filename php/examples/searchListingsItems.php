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
    // Call searchListingsItems with various search parameters
    $response = $listingsApi->searchListingsItems(
        'A1B2C3D4E5F6G7', // sellerId
        ['ATVPDKIKX0DER'], // marketplaceIds
        'en_US', // issueLocale
        ['summaries', 'offers', 'fulfillmentAvailability', 'issues'], // includedData
        ['GM-ZDPI-9B4E', 'HW-ZDPI-9B4E', 'TC-ZDPI-9B4E'], // identifiers
        'SKU', // identifiersType
        null, // variationParentSku
        null, // packageHierarchySku
        new DateTime('2023-01-01T00:00:00Z'), // createdAfter
        null, // createdBefore
        new DateTime('2023-12-01T00:00:00Z'), // lastUpdatedAfter
        null, // lastUpdatedBefore
        ['ERROR', 'WARNING'], // withIssueSeverity
        ['BUYABLE', 'SEARCHABLE'], // withStatus
        null, // withoutStatus
        'lastUpdatedDate', // sortBy
        'DESC', // sortOrder
        10, // pageSize
        null // pageToken
    );

    // Process Listings API response
    echo "Search Listings Items API Response:\n";
    print_r($response);

} catch (Exception $e) {
    echo 'Exception when calling ListingsApi->searchListingsItems: ', $e->getMessage(), PHP_EOL;
} 