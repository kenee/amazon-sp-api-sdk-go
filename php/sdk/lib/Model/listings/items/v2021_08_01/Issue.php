<?php

/**
 * Issue.
 *
 * PHP version 8.3
 *
 * @category Class
 *
 * @author   OpenAPI Generator team
 *
 * @see     https://openapi-generator.tech
 */

/**
 * Selling Partner API for Listings Items.
 *
 * The Selling Partner API for Listings Items (Listings Items API) provides programmatic access to selling partner listings on Amazon. Use this API in collaboration with the Selling Partner API for Product Type Definitions, which you use to retrieve the information about Amazon product types needed to use the Listings Items API.  For more information, see the [Listings Items API Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/listings-items-api-v2021-08-01-use-case-guide).
 *
 * The version of the OpenAPI document: 2021-08-01
 * Generated by: https://openapi-generator.tech
 * Generator version: 7.9.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace SpApi\Model\listings\items\v2021_08_01;

use SpApi\Model\ModelInterface;
use SpApi\ObjectSerializer;

/**
 * Issue Class Doc Comment.
 *
 * @category Class
 *
 * @description An issue with a listings item.
 *
 * @author   OpenAPI Generator team
 *
 * @see     https://openapi-generator.tech
 *
 * @implements \ArrayAccess<string, mixed>
 */
class Issue implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    public const SEVERITY_ERROR = 'ERROR';
    public const SEVERITY_WARNING = 'WARNING';
    public const SEVERITY_INFO = 'INFO';

    /**
     * The original name of the model.
     */
    protected static string $openAPIModelName = 'Issue';

    /**
     * Array of property to type mappings. Used for (de)serialization.
     *
     * @var string[]
     */
    protected static array $openAPITypes = [
        'code' => 'string',
        'message' => 'string',
        'severity' => 'string',
        'attribute_names' => 'string[]',
        'categories' => 'string[]',
        'enforcements' => '\SpApi\Model\listings\items\v2021_08_01\IssueEnforcements'];

    /**
     * Array of property to format mappings. Used for (de)serialization.
     *
     * @var string[]
     *
     * @phpstan-var array<string, string|null>
     *
     * @psalm-var array<string, string|null>
     */
    protected static array $openAPIFormats = [
        'code' => null,
        'message' => null,
        'severity' => null,
        'attribute_names' => null,
        'categories' => null,
        'enforcements' => null];

    /**
     * Array of nullable properties. Used for (de)serialization.
     *
     * @var bool[]
     */
    protected static array $openAPINullables = [
        'code' => false,
        'message' => false,
        'severity' => false,
        'attribute_names' => true,
        'categories' => false,
        'enforcements' => true,
    ];

    /**
     * If a nullable field gets set to null, insert it here.
     *
     * @var bool[]
     */
    protected array $openAPINullablesSetToNull = [];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     *
     * @var string[]
     */
    protected static array $attributeMap = [
        'code' => 'code',
        'message' => 'message',
        'severity' => 'severity',
        'attribute_names' => 'attributeNames',
        'categories' => 'categories',
        'enforcements' => 'enforcements',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     *
     * @var string[]
     */
    protected static array $setters = [
        'code' => 'setCode',
        'message' => 'setMessage',
        'severity' => 'setSeverity',
        'attribute_names' => 'setAttributeNames',
        'categories' => 'setCategories',
        'enforcements' => 'setEnforcements',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests).
     *
     * @var string[]
     */
    protected static array $getters = [
        'code' => 'getCode',
        'message' => 'getMessage',
        'severity' => 'getSeverity',
        'attribute_names' => 'getAttributeNames',
        'categories' => 'getCategories',
        'enforcements' => 'getEnforcements',
    ];

    /**
     * Associative array for storing property values.
     */
    protected array $container = [];

    /**
     * Constructor.
     *
     * @param null|array $data Associated array of property values
     *                         initializing the model
     */
    public function __construct(?array $data = null)
    {
        $this->setIfExists('code', $data ?? [], null);
        $this->setIfExists('message', $data ?? [], null);
        $this->setIfExists('severity', $data ?? [], null);
        $this->setIfExists('attribute_names', $data ?? [], null);
        $this->setIfExists('categories', $data ?? [], null);
        $this->setIfExists('enforcements', $data ?? [], null);
    }

    /**
     * Gets the string presentation of the object.
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Array of property to type mappings. Used for (de)serialization.
     */
    public static function openAPITypes(): array
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization.
     */
    public static function openAPIFormats(): array
    {
        return self::$openAPIFormats;
    }

    /**
     * Checks if a property is nullable.
     */
    public static function isNullable(string $property): bool
    {
        return self::openAPINullables()[$property] ?? false;
    }

    /**
     * Checks if a nullable property is set to null.
     */
    public function isNullableSetToNull(string $property): bool
    {
        return in_array($property, $this->getOpenAPINullablesSetToNull(), true);
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name.
     */
    public static function attributeMap(): array
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses).
     */
    public static function setters(): array
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests).
     */
    public static function getters(): array
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     */
    public function getModelName(): string
    {
        return self::$openAPIModelName;
    }

    /**
     * Gets allowable values of the enum.
     *
     * @return string[]
     */
    public function getSeverityAllowableValues(): array
    {
        return [
            self::SEVERITY_ERROR,
            self::SEVERITY_WARNING,
            self::SEVERITY_INFO,
        ];
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties(): array
    {
        $invalidProperties = [];

        if (null === $this->container['code']) {
            $invalidProperties[] = "'code' can't be null";
        }
        if (null === $this->container['message']) {
            $invalidProperties[] = "'message' can't be null";
        }
        if (null === $this->container['severity']) {
            $invalidProperties[] = "'severity' can't be null";
        }
        $allowedValues = $this->getSeverityAllowableValues();
        if (!is_null($this->container['severity']) && !in_array($this->container['severity'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'severity', must be one of '%s'",
                $this->container['severity'],
                implode("', '", $allowedValues)
            );
        }

        if (null === $this->container['categories']) {
            $invalidProperties[] = "'categories' can't be null";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed.
     *
     * @return bool True if all properties are valid
     */
    public function valid(): bool
    {
        return 0 === count($this->listInvalidProperties());
    }

    /**
     * Gets code.
     */
    public function getCode(): string
    {
        return $this->container['code'];
    }

    /**
     * Sets code.
     *
     * @param string $code an issue code that identifies the type of issue
     */
    public function setCode(string $code): self
    {
        if (is_null($code)) {
            throw new \InvalidArgumentException('non-nullable code cannot be null');
        }
        $this->container['code'] = $code;

        return $this;
    }

    /**
     * Gets message.
     */
    public function getMessage(): string
    {
        return $this->container['message'];
    }

    /**
     * Sets message.
     *
     * @param string $message a message that describes the issue
     */
    public function setMessage(string $message): self
    {
        if (is_null($message)) {
            throw new \InvalidArgumentException('non-nullable message cannot be null');
        }
        $this->container['message'] = $message;

        return $this;
    }

    /**
     * Gets severity.
     */
    public function getSeverity(): string
    {
        return $this->container['severity'];
    }

    /**
     * Sets severity.
     *
     * @param string $severity the severity of the issue
     */
    public function setSeverity(string $severity): self
    {
        if (is_null($severity)) {
            throw new \InvalidArgumentException('non-nullable severity cannot be null');
        }
        $allowedValues = $this->getSeverityAllowableValues();
        if (!in_array($severity, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'severity', must be one of '%s'",
                    $severity,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['severity'] = $severity;

        return $this;
    }

    /**
     * Gets attribute_names.
     */
    public function getAttributeNames(): ?array
    {
        return $this->container['attribute_names'];
    }

    /**
     * Sets attribute_names.
     *
     * @param null|array $attribute_names the names of the attributes associated with the issue, if applicable
     */
    public function setAttributeNames(?array $attribute_names): self
    {
        if (is_null($attribute_names)) {
            array_push($this->openAPINullablesSetToNull, 'attribute_names');
        } else {
            $nullablesSetToNull = $this->getOpenAPINullablesSetToNull();
            $index = array_search('attribute_names', $nullablesSetToNull);
            if (false !== $index) {
                unset($nullablesSetToNull[$index]);
                $this->setOpenAPINullablesSetToNull($nullablesSetToNull);
            }
        }
        $this->container['attribute_names'] = $attribute_names;

        return $this;
    }

    /**
     * Gets categories.
     */
    public function getCategories(): array
    {
        return $this->container['categories'];
    }

    /**
     * Sets categories.
     *
     * @param array $categories List of issue categories.   Possible values:   * 'INVALID_ATTRIBUTE' - Indicating an invalid attribute in the listing.   * 'MISSING_ATTRIBUTE' - Highlighting a missing attribute in the listing.   * 'INVALID_IMAGE' - Signifying an invalid image in the listing.   * 'MISSING_IMAGE' - Noting the absence of an image in the listing.   * 'INVALID_PRICE' - Pertaining to issues with the listing's price-related attributes.   * 'MISSING_PRICE' - Pointing out the absence of a price attribute in the listing.   * 'DUPLICATE' - Identifying listings with potential duplicate problems, such as this ASIN potentially being a duplicate of another ASIN.   * 'QUALIFICATION_REQUIRED' - Indicating that the listing requires qualification-related approval.
     */
    public function setCategories(array $categories): self
    {
        if (is_null($categories)) {
            throw new \InvalidArgumentException('non-nullable categories cannot be null');
        }
        $this->container['categories'] = $categories;

        return $this;
    }

    /**
     * Gets enforcements.
     */
    public function getEnforcements(): ?IssueEnforcements
    {
        return $this->container['enforcements'];
    }

    /**
     * Sets enforcements.
     *
     * @param null|IssueEnforcements $enforcements enforcements
     */
    public function setEnforcements(?IssueEnforcements $enforcements): self
    {
        if (is_null($enforcements)) {
            array_push($this->openAPINullablesSetToNull, 'enforcements');
        } else {
            $nullablesSetToNull = $this->getOpenAPINullablesSetToNull();
            $index = array_search('enforcements', $nullablesSetToNull);
            if (false !== $index) {
                unset($nullablesSetToNull[$index]);
                $this->setOpenAPINullablesSetToNull($nullablesSetToNull);
            }
        }
        $this->container['enforcements'] = $enforcements;

        return $this;
    }

    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     */
    public function offsetExists($offset): bool
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return null|mixed
     */
    #[\ReturnTypeWillChange]
    public function offsetGet($offset): mixed
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param null|int $offset Offset
     * @param mixed    $value  Value to be set
     */
    public function offsetSet($offset, mixed $value): void
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     */
    public function offsetUnset($offset): void
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     *
     * @see https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed returns data which can be serialized by json_encode(), which is a value
     *               of any type other than a resource
     */
    #[\ReturnTypeWillChange]
    public function jsonSerialize(): mixed
    {
        return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets a header-safe presentation of the object.
     */
    public function toHeaderValue(): string
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }

    /**
     * Array of nullable properties.
     */
    protected static function openAPINullables(): array
    {
        return self::$openAPINullables;
    }

    /**
     * Array of nullable field names deliberately set to null.
     *
     * @return bool[]
     */
    private function getOpenAPINullablesSetToNull(): array
    {
        return $this->openAPINullablesSetToNull;
    }

    /**
     * Setter - Array of nullable field names deliberately set to null.
     *
     * @param bool[] $openAPINullablesSetToNull
     */
    private function setOpenAPINullablesSetToNull(array $openAPINullablesSetToNull): void
    {
        $this->openAPINullablesSetToNull = $openAPINullablesSetToNull;
    }

    /**
     * Sets $this->container[$variableName] to the given data or to the given default Value; if $variableName
     * is nullable and its value is set to null in the $fields array, then mark it as "set to null" in the
     * $this->openAPINullablesSetToNull array.
     *
     * @param mixed $defaultValue
     */
    private function setIfExists(string $variableName, array $fields, $defaultValue): void
    {
        if (self::isNullable($variableName) && array_key_exists($variableName, $fields) && is_null($fields[$variableName])) {
            $this->openAPINullablesSetToNull[] = $variableName;
        }

        $this->container[$variableName] = $fields[$variableName] ?? $defaultValue;
    }
}
