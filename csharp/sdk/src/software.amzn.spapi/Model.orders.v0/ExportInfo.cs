/*
 * Selling Partner API for Orders
 *
 * Use the Orders Selling Partner API to programmatically retrieve order information. With this API, you can develop fast, flexible, and custom applications to manage order synchronization, perform order research, and create demand-based decision support tools.   _Note:_ For the JP, AU, and SG marketplaces, the Orders API supports orders from 2016 onward. For all other marketplaces, the Orders API supports orders for the last two years (orders older than this don't show up in the response).
 *
 * The version of the OpenAPI document: v0
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

namespace software.amzn.spapi.Model.orders.v0
{
    /// <summary>
    /// Contains information that is related to the export of an order item.
    /// </summary>
    [DataContract(Name = "ExportInfo")]
    public partial class ExportInfo : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ExportInfo" /> class.
        /// </summary>
        /// <param name="exportCharge">exportCharge.</param>
        /// <param name="exportChargeModel">Holds the &#x60;ExportCharge&#x60; collection model that is associated with the specified order item.\\n\\n**Possible values**: &#x60;AMAZON_FACILITATED&#x60;: Import/export charge is withheld by Amazon and remitted to the customs authority by the carrier on behalf of the buyer/seller..</param>
        public ExportInfo(Money exportCharge = default(Money), string exportChargeModel = default(string))
        {
            this.ExportCharge = exportCharge;
            this.ExportChargeModel = exportChargeModel;
        }

        /// <summary>
        /// Gets or Sets ExportCharge
        /// </summary>
        [DataMember(Name = "ExportCharge", EmitDefaultValue = false)]
        public Money ExportCharge { get; set; }

        /// <summary>
        /// Holds the &#x60;ExportCharge&#x60; collection model that is associated with the specified order item.\\n\\n**Possible values**: &#x60;AMAZON_FACILITATED&#x60;: Import/export charge is withheld by Amazon and remitted to the customs authority by the carrier on behalf of the buyer/seller.
        /// </summary>
        /// <value>Holds the &#x60;ExportCharge&#x60; collection model that is associated with the specified order item.\\n\\n**Possible values**: &#x60;AMAZON_FACILITATED&#x60;: Import/export charge is withheld by Amazon and remitted to the customs authority by the carrier on behalf of the buyer/seller.</value>
        [DataMember(Name = "ExportChargeModel", EmitDefaultValue = false)]
        public string ExportChargeModel { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ExportInfo {\n");
            sb.Append("  ExportCharge: ").Append(ExportCharge).Append("\n");
            sb.Append("  ExportChargeModel: ").Append(ExportChargeModel).Append("\n");
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
