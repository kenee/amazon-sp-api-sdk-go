/*
 * Amazon Shipping API
 *
 * The Amazon Shipping API is designed to support outbound shipping use cases both for orders originating on Amazon-owned marketplaces as well as external channels/marketplaces. With these APIs, you can request shipping rates, create shipments, cancel shipments, and track shipments.
 *
 * The version of the OpenAPI document: v2
 * Contact: swa-api-core@amazon.com
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

namespace software.amzn.spapi.Model.shipping.v2
{
    /// <summary>
    /// The format options available for a label.
    /// </summary>
    [DataContract(Name = "PrintOption")]
    public partial class PrintOption : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="PrintOption" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected PrintOption() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="PrintOption" /> class.
        /// </summary>
        /// <param name="supportedDPIs">A list of the supported DPI options for a document..</param>
        /// <param name="supportedPageLayouts">A list of the supported page layout options for a document. (required).</param>
        /// <param name="supportedFileJoiningOptions">A list of the supported needFileJoining boolean values for a document. (required).</param>
        /// <param name="supportedDocumentDetails">A list of the supported documented details. (required).</param>
        public PrintOption(List<int> supportedDPIs = default(List<int>), List<string> supportedPageLayouts = default(List<string>), List<bool> supportedFileJoiningOptions = default(List<bool>), List<SupportedDocumentDetail> supportedDocumentDetails = default(List<SupportedDocumentDetail>))
        {
            // to ensure "supportedPageLayouts" is required (not null)
            if (supportedPageLayouts == null)
            {
                throw new ArgumentNullException("supportedPageLayouts is a required property for PrintOption and cannot be null");
            }
            this.SupportedPageLayouts = supportedPageLayouts;
            // to ensure "supportedFileJoiningOptions" is required (not null)
            if (supportedFileJoiningOptions == null)
            {
                throw new ArgumentNullException("supportedFileJoiningOptions is a required property for PrintOption and cannot be null");
            }
            this.SupportedFileJoiningOptions = supportedFileJoiningOptions;
            // to ensure "supportedDocumentDetails" is required (not null)
            if (supportedDocumentDetails == null)
            {
                throw new ArgumentNullException("supportedDocumentDetails is a required property for PrintOption and cannot be null");
            }
            this.SupportedDocumentDetails = supportedDocumentDetails;
            this.SupportedDPIs = supportedDPIs;
        }

        /// <summary>
        /// A list of the supported DPI options for a document.
        /// </summary>
        /// <value>A list of the supported DPI options for a document.</value>
        [DataMember(Name = "supportedDPIs", EmitDefaultValue = false)]
        public List<int> SupportedDPIs { get; set; }

        /// <summary>
        /// A list of the supported page layout options for a document.
        /// </summary>
        /// <value>A list of the supported page layout options for a document.</value>
        [DataMember(Name = "supportedPageLayouts", IsRequired = true, EmitDefaultValue = true)]
        public List<string> SupportedPageLayouts { get; set; }

        /// <summary>
        /// A list of the supported needFileJoining boolean values for a document.
        /// </summary>
        /// <value>A list of the supported needFileJoining boolean values for a document.</value>
        [DataMember(Name = "supportedFileJoiningOptions", IsRequired = true, EmitDefaultValue = true)]
        public List<bool> SupportedFileJoiningOptions { get; set; }

        /// <summary>
        /// A list of the supported documented details.
        /// </summary>
        /// <value>A list of the supported documented details.</value>
        [DataMember(Name = "supportedDocumentDetails", IsRequired = true, EmitDefaultValue = true)]
        public List<SupportedDocumentDetail> SupportedDocumentDetails { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class PrintOption {\n");
            sb.Append("  SupportedDPIs: ").Append(SupportedDPIs).Append("\n");
            sb.Append("  SupportedPageLayouts: ").Append(SupportedPageLayouts).Append("\n");
            sb.Append("  SupportedFileJoiningOptions: ").Append(SupportedFileJoiningOptions).Append("\n");
            sb.Append("  SupportedDocumentDetails: ").Append(SupportedDocumentDetails).Append("\n");
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
