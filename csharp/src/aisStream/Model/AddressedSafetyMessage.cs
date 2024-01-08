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
    /// AddressedSafetyMessage
    /// </summary>
    [DataContract(Name = "AddressedSafetyMessage")]
    public partial class AddressedSafetyMessage : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="AddressedSafetyMessage" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected AddressedSafetyMessage() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="AddressedSafetyMessage" /> class.
        /// </summary>
        /// <param name="messageID">messageID (required).</param>
        /// <param name="repeatIndicator">repeatIndicator (required).</param>
        /// <param name="userID">userID (required).</param>
        /// <param name="valid">valid (required).</param>
        /// <param name="sequenceinteger">sequenceinteger (required).</param>
        /// <param name="destinationID">destinationID (required).</param>
        /// <param name="retransmission">retransmission (required).</param>
        /// <param name="spare">spare (required).</param>
        /// <param name="text">text (required).</param>
        public AddressedSafetyMessage(int messageID = default(int), int repeatIndicator = default(int), int userID = default(int), bool valid = default(bool), int sequenceinteger = default(int), int destinationID = default(int), bool retransmission = default(bool), bool spare = default(bool), string text = default(string))
        {
            this.MessageID = messageID;
            this.RepeatIndicator = repeatIndicator;
            this.UserID = userID;
            this.Valid = valid;
            this.Sequenceinteger = sequenceinteger;
            this.DestinationID = destinationID;
            this.Retransmission = retransmission;
            this.Spare = spare;
            // to ensure "text" is required (not null)
            if (text == null)
            {
                throw new ArgumentNullException("text is a required property for AddressedSafetyMessage and cannot be null");
            }
            this.Text = text;
        }

        /// <summary>
        /// Gets or Sets MessageID
        /// </summary>
        [DataMember(Name = "MessageID", IsRequired = true, EmitDefaultValue = true)]
        public int MessageID { get; set; }

        /// <summary>
        /// Gets or Sets RepeatIndicator
        /// </summary>
        [DataMember(Name = "RepeatIndicator", IsRequired = true, EmitDefaultValue = true)]
        public int RepeatIndicator { get; set; }

        /// <summary>
        /// Gets or Sets UserID
        /// </summary>
        [DataMember(Name = "UserID", IsRequired = true, EmitDefaultValue = true)]
        public int UserID { get; set; }

        /// <summary>
        /// Gets or Sets Valid
        /// </summary>
        [DataMember(Name = "Valid", IsRequired = true, EmitDefaultValue = true)]
        public bool Valid { get; set; }

        /// <summary>
        /// Gets or Sets Sequenceinteger
        /// </summary>
        [DataMember(Name = "Sequenceinteger", IsRequired = true, EmitDefaultValue = true)]
        public int Sequenceinteger { get; set; }

        /// <summary>
        /// Gets or Sets DestinationID
        /// </summary>
        [DataMember(Name = "DestinationID", IsRequired = true, EmitDefaultValue = true)]
        public int DestinationID { get; set; }

        /// <summary>
        /// Gets or Sets Retransmission
        /// </summary>
        [DataMember(Name = "Retransmission", IsRequired = true, EmitDefaultValue = true)]
        public bool Retransmission { get; set; }

        /// <summary>
        /// Gets or Sets Spare
        /// </summary>
        [DataMember(Name = "Spare", IsRequired = true, EmitDefaultValue = true)]
        public bool Spare { get; set; }

        /// <summary>
        /// Gets or Sets Text
        /// </summary>
        [DataMember(Name = "Text", IsRequired = true, EmitDefaultValue = true)]
        public string Text { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class AddressedSafetyMessage {\n");
            sb.Append("  MessageID: ").Append(MessageID).Append("\n");
            sb.Append("  RepeatIndicator: ").Append(RepeatIndicator).Append("\n");
            sb.Append("  UserID: ").Append(UserID).Append("\n");
            sb.Append("  Valid: ").Append(Valid).Append("\n");
            sb.Append("  Sequenceinteger: ").Append(Sequenceinteger).Append("\n");
            sb.Append("  DestinationID: ").Append(DestinationID).Append("\n");
            sb.Append("  Retransmission: ").Append(Retransmission).Append("\n");
            sb.Append("  Spare: ").Append(Spare).Append("\n");
            sb.Append("  Text: ").Append(Text).Append("\n");
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