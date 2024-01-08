/*
 * Ais-Stream WebsocketObjects
 *
 * OpenAPI 3.0 definitions for the data models used by aisstream.io.
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;
using System.Text;

namespace aisStream.Model
{
    /// <summary>
    /// AisStreamMessage
    /// </summary>
    [DataContract(Name = "AisStreamMessage")]
    public partial class AisStreamMessage : IValidatableObject
    {

        /// <summary>
        /// Gets or Sets MessageType
        /// </summary>
        [DataMember(Name = "MessageType", IsRequired = true, EmitDefaultValue = true)]
        public AisMessageTypes MessageType { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="AisStreamMessage" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected AisStreamMessage() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="AisStreamMessage" /> class.
        /// </summary>
        /// <param name="metaData">metaData (required).</param>
        /// <param name="messageType">messageType (required).</param>
        /// <param name="message">message (required).</param>
        public AisStreamMessage(Object metaData = default(Object), AisMessageTypes messageType = default(AisMessageTypes), AisStreamMessageMessage message = default(AisStreamMessageMessage))
        {
            // to ensure "metaData" is required (not null)
            if (metaData == null)
            {
                throw new ArgumentNullException("metaData is a required property for AisStreamMessage and cannot be null");
            }
            this.MetaData = metaData;
            this.MessageType = messageType;
            // to ensure "message" is required (not null)
            if (message == null)
            {
                throw new ArgumentNullException("message is a required property for AisStreamMessage and cannot be null");
            }
            this.Message = message;
        }

        /// <summary>
        /// Gets or Sets MetaData
        /// </summary>
        [DataMember(Name = "MetaData", IsRequired = true, EmitDefaultValue = true)]
        public Object MetaData { get; set; }

        /// <summary>
        /// Gets or Sets Message
        /// </summary>
        [DataMember(Name = "Message", IsRequired = true, EmitDefaultValue = true)]
        public AisStreamMessageMessage Message { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class AisStreamMessage {\n");
            sb.Append("  MetaData: ").Append(MetaData).Append("\n");
            sb.Append("  MessageType: ").Append(MessageType).Append("\n");
            sb.Append("  Message: ").Append(Message).Append("\n");
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
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}