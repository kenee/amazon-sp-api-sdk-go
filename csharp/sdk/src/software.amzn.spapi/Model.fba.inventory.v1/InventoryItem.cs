/*
 * Selling Partner API for FBA Inventory
 *
 * The Selling Partner API for FBA Inventory lets you programmatically retrieve information about inventory in Amazon's fulfillment network.
 *
 * The version of the OpenAPI document: v1
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = software.amzn.spapi.Client.OpenAPIDateConverter;

namespace software.amzn.spapi.Model.fba.inventory.v1
{
    /// <summary>
    /// An item in the list of inventory to be added.
    /// </summary>
    [DataContract(Name = "InventoryItem")]
    public partial class InventoryItem : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="InventoryItem" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected InventoryItem() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="InventoryItem" /> class.
        /// </summary>
        /// <param name="sellerSku">The seller SKU of the item. (required).</param>
        /// <param name="marketplaceId">The marketplaceId. (required).</param>
        /// <param name="quantity">The quantity of item to add. (required).</param>
        public InventoryItem(string sellerSku = default(string), string marketplaceId = default(string), int quantity = default(int))
        {
            // to ensure "sellerSku" is required (not null)
            if (sellerSku == null)
            {
                throw new ArgumentNullException("sellerSku is a required property for InventoryItem and cannot be null");
            }
            this.SellerSku = sellerSku;
            // to ensure "marketplaceId" is required (not null)
            if (marketplaceId == null)
            {
                throw new ArgumentNullException("marketplaceId is a required property for InventoryItem and cannot be null");
            }
            this.MarketplaceId = marketplaceId;
            this.Quantity = quantity;
        }

        /// <summary>
        /// The seller SKU of the item.
        /// </summary>
        /// <value>The seller SKU of the item.</value>
        [DataMember(Name = "sellerSku", IsRequired = true, EmitDefaultValue = true)]
        public string SellerSku { get; set; }

        /// <summary>
        /// The marketplaceId.
        /// </summary>
        /// <value>The marketplaceId.</value>
        [DataMember(Name = "marketplaceId", IsRequired = true, EmitDefaultValue = true)]
        public string MarketplaceId { get; set; }

        /// <summary>
        /// The quantity of item to add.
        /// </summary>
        /// <value>The quantity of item to add.</value>
        [DataMember(Name = "quantity", IsRequired = true, EmitDefaultValue = true)]
        public int Quantity { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class InventoryItem {\n");
            sb.Append("  SellerSku: ").Append(SellerSku).Append("\n");
            sb.Append("  MarketplaceId: ").Append(MarketplaceId).Append("\n");
            sb.Append("  Quantity: ").Append(Quantity).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
