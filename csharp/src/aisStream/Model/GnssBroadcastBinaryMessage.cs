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
    /// GnssBroadcastBinaryMessage
    /// </summary>
    [DataContract(Name = "GnssBroadcastBinaryMessage")]
    public partial class GnssBroadcastBinaryMessage : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="GnssBroadcastBinaryMessage" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected GnssBroadcastBinaryMessage() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="GnssBroadcastBinaryMessage" /> class.
        /// </summary>
        /// <param name="messageID">messageID (required).</param>
        /// <param name="repeatIndicator">repeatIndicator (required).</param>
        /// <param name="userID">userID (required).</param>
        /// <param name="valid">valid (required).</param>
        /// <param name="spare1">spare1 (required).</param>
        /// <param name="longitude">longitude (required).</param>
        /// <param name="latitude">latitude (required).</param>
        /// <param name="spare2">spare2 (required).</param>
        /// <param name="data">data (required).</param>
        public GnssBroadcastBinaryMessage(int messageID = default(int), int repeatIndicator = default(int), int userID = default(int), bool valid = default(bool), int spare1 = default(int), double longitude = default(double), double latitude = default(double), int spare2 = default(int), string data = default(string))
        {
            this.MessageID = messageID;
            this.RepeatIndicator = repeatIndicator;
            this.UserID = userID;
            this.Valid = valid;
            this.Spare1 = spare1;
            this.Longitude = longitude;
            this.Latitude = latitude;
            this.Spare2 = spare2;
            // to ensure "data" is required (not null)
            if (data == null)
            {
                throw new ArgumentNullException("data is a required property for GnssBroadcastBinaryMessage and cannot be null");
            }
            this.Data = data;
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
        /// Gets or Sets Spare1
        /// </summary>
        [DataMember(Name = "Spare1", IsRequired = true, EmitDefaultValue = true)]
        public int Spare1 { get; set; }

        /// <summary>
        /// Gets or Sets Longitude
        /// </summary>
        [DataMember(Name = "Longitude", IsRequired = true, EmitDefaultValue = true)]
        public double Longitude { get; set; }

        /// <summary>
        /// Gets or Sets Latitude
        /// </summary>
        [DataMember(Name = "Latitude", IsRequired = true, EmitDefaultValue = true)]
        public double Latitude { get; set; }

        /// <summary>
        /// Gets or Sets Spare2
        /// </summary>
        [DataMember(Name = "Spare2", IsRequired = true, EmitDefaultValue = true)]
        public int Spare2 { get; set; }

        /// <summary>
        /// Gets or Sets Data
        /// </summary>
        [DataMember(Name = "Data", IsRequired = true, EmitDefaultValue = true)]
        public string Data { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class GnssBroadcastBinaryMessage {\n");
            sb.Append("  MessageID: ").Append(MessageID).Append("\n");
            sb.Append("  RepeatIndicator: ").Append(RepeatIndicator).Append("\n");
            sb.Append("  UserID: ").Append(UserID).Append("\n");
            sb.Append("  Valid: ").Append(Valid).Append("\n");
            sb.Append("  Spare1: ").Append(Spare1).Append("\n");
            sb.Append("  Longitude: ").Append(Longitude).Append("\n");
            sb.Append("  Latitude: ").Append(Latitude).Append("\n");
            sb.Append("  Spare2: ").Append(Spare2).Append("\n");
            sb.Append("  Data: ").Append(Data).Append("\n");
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