/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * The version of the OpenAPI document: 2020-07-01
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

namespace software.amzn.spapi.Model.fulfillment.outbound.v2020_07_01
{
    /// <summary>
    /// Information for tracking package deliveries.
    /// </summary>
    [DataContract(Name = "TrackingEvent")]
    public partial class TrackingEvent : IValidatableObject
    {

        /// <summary>
        /// Gets or Sets EventCode
        /// </summary>
        [DataMember(Name = "eventCode", IsRequired = true, EmitDefaultValue = true)]
        public EventCode EventCode { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="TrackingEvent" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected TrackingEvent() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="TrackingEvent" /> class.
        /// </summary>
        /// <param name="eventDate">Date timestamp (required).</param>
        /// <param name="eventAddress">eventAddress (required).</param>
        /// <param name="eventCode">eventCode (required).</param>
        /// <param name="eventDescription">A description for the corresponding event code. (required).</param>
        public TrackingEvent(DateTime eventDate = default(DateTime), TrackingAddress eventAddress = default(TrackingAddress), EventCode eventCode = default(EventCode), string eventDescription = default(string))
        {
            this.EventDate = eventDate;
            // to ensure "eventAddress" is required (not null)
            if (eventAddress == null)
            {
                throw new ArgumentNullException("eventAddress is a required property for TrackingEvent and cannot be null");
            }
            this.EventAddress = eventAddress;
            this.EventCode = eventCode;
            // to ensure "eventDescription" is required (not null)
            if (eventDescription == null)
            {
                throw new ArgumentNullException("eventDescription is a required property for TrackingEvent and cannot be null");
            }
            this.EventDescription = eventDescription;
        }

        /// <summary>
        /// Date timestamp
        /// </summary>
        /// <value>Date timestamp</value>
        [DataMember(Name = "eventDate", IsRequired = true, EmitDefaultValue = true)]
        public DateTime EventDate { get; set; }

        /// <summary>
        /// Gets or Sets EventAddress
        /// </summary>
        [DataMember(Name = "eventAddress", IsRequired = true, EmitDefaultValue = true)]
        public TrackingAddress EventAddress { get; set; }

        /// <summary>
        /// A description for the corresponding event code.
        /// </summary>
        /// <value>A description for the corresponding event code.</value>
        [DataMember(Name = "eventDescription", IsRequired = true, EmitDefaultValue = true)]
        public string EventDescription { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class TrackingEvent {\n");
            sb.Append("  EventDate: ").Append(EventDate).Append("\n");
            sb.Append("  EventAddress: ").Append(EventAddress).Append("\n");
            sb.Append("  EventCode: ").Append(EventCode).Append("\n");
            sb.Append("  EventDescription: ").Append(EventDescription).Append("\n");
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
