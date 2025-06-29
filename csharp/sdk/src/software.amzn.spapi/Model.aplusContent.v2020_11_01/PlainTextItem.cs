/*
 * Selling Partner API for A+ Content Management
 *
 * Use the A+ Content API to build applications that help selling partners add rich marketing content to their Amazon product detail pages. Selling partners can use A+ content to share their brand and product story, which helps buyers make informed purchasing decisions. Selling partners use content modules to add images and text.
 *
 * The version of the OpenAPI document: 2020-11-01
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

namespace software.amzn.spapi.Model.aplusContent.v2020_11_01
{
    /// <summary>
    /// Plain positional text that is used in collections of brief labels and descriptors.
    /// </summary>
    [DataContract(Name = "PlainTextItem")]
    public partial class PlainTextItem : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="PlainTextItem" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected PlainTextItem() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="PlainTextItem" /> class.
        /// </summary>
        /// <param name="position">The rank or index of this text item within the collection. Different items cannot occupy the same position within a single collection. (required).</param>
        /// <param name="value">The actual plain text. (required).</param>
        public PlainTextItem(int position = default(int), string value = default(string))
        {
            this.Position = position;
            // to ensure "value" is required (not null)
            if (value == null)
            {
                throw new ArgumentNullException("value is a required property for PlainTextItem and cannot be null");
            }
            this.Value = value;
        }

        /// <summary>
        /// The rank or index of this text item within the collection. Different items cannot occupy the same position within a single collection.
        /// </summary>
        /// <value>The rank or index of this text item within the collection. Different items cannot occupy the same position within a single collection.</value>
        [DataMember(Name = "position", IsRequired = true, EmitDefaultValue = true)]
        public int Position { get; set; }

        /// <summary>
        /// The actual plain text.
        /// </summary>
        /// <value>The actual plain text.</value>
        [DataMember(Name = "value", IsRequired = true, EmitDefaultValue = true)]
        public string Value { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class PlainTextItem {\n");
            sb.Append("  Position: ").Append(Position).Append("\n");
            sb.Append("  Value: ").Append(Value).Append("\n");
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
            // Position (int) maximum
            if (this.Position > (int)100)
            {
                yield return new ValidationResult("Invalid value for Position, must be a value less than or equal to 100.", new [] { "Position" });
            }

            // Position (int) minimum
            if (this.Position < (int)1)
            {
                yield return new ValidationResult("Invalid value for Position, must be a value greater than or equal to 1.", new [] { "Position" });
            }

            // Value (string) maxLength
            if (this.Value != null && this.Value.Length > 250)
            {
                yield return new ValidationResult("Invalid value for Value, length must be less than 250.", new [] { "Value" });
            }

            // Value (string) minLength
            if (this.Value != null && this.Value.Length < 1)
            {
                yield return new ValidationResult("Invalid value for Value, length must be greater than 1.", new [] { "Value" });
            }

            yield break;
        }
    }

}
