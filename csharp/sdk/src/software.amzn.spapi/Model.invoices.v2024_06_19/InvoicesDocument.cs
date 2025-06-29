/*
 * The Selling Partner API for Invoices.
 *
 * Use the Selling Partner API for Invoices to retrieve and manage invoice-related operations, which can help selling partners manage their bookkeeping processes.
 *
 * The version of the OpenAPI document: 2024-06-19
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

namespace software.amzn.spapi.Model.invoices.v2024_06_19
{
    /// <summary>
    /// An object that contains the &#x60;documentId&#x60; and an S3 pre-signed URL that you can use to download the specified file.
    /// </summary>
    [DataContract(Name = "InvoicesDocument")]
    public partial class InvoicesDocument : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="InvoicesDocument" /> class.
        /// </summary>
        /// <param name="invoicesDocumentId">The identifier of the export document..</param>
        /// <param name="invoicesDocumentUrl">A pre-signed URL that you can use to download the invoices document in zip format. This URL expires after 30 seconds..</param>
        public InvoicesDocument(string invoicesDocumentId = default(string), string invoicesDocumentUrl = default(string))
        {
            this.InvoicesDocumentId = invoicesDocumentId;
            this.InvoicesDocumentUrl = invoicesDocumentUrl;
        }

        /// <summary>
        /// The identifier of the export document.
        /// </summary>
        /// <value>The identifier of the export document.</value>
        [DataMember(Name = "invoicesDocumentId", EmitDefaultValue = false)]
        public string InvoicesDocumentId { get; set; }

        /// <summary>
        /// A pre-signed URL that you can use to download the invoices document in zip format. This URL expires after 30 seconds.
        /// </summary>
        /// <value>A pre-signed URL that you can use to download the invoices document in zip format. This URL expires after 30 seconds.</value>
        [DataMember(Name = "invoicesDocumentUrl", EmitDefaultValue = false)]
        public string InvoicesDocumentUrl { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class InvoicesDocument {\n");
            sb.Append("  InvoicesDocumentId: ").Append(InvoicesDocumentId).Append("\n");
            sb.Append("  InvoicesDocumentUrl: ").Append(InvoicesDocumentUrl).Append("\n");
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
