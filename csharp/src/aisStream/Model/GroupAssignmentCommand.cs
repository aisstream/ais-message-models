/*
 * Ais-Stream WebsocketObjects
 *
 * OpenAPI 3.0 definitions for the data models used by aisstream.io.
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Newtonsoft.Json;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Runtime.Serialization;
using System.Text;

namespace aisStream.Model
{
    /// <summary>
    /// GroupAssignmentCommand
    /// </summary>
    [DataContract(Name = "GroupAssignmentCommand")]
    public partial class GroupAssignmentCommand : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="GroupAssignmentCommand" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected GroupAssignmentCommand() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="GroupAssignmentCommand" /> class.
        /// </summary>
        /// <param name="messageID">messageID (required).</param>
        /// <param name="repeatIndicator">repeatIndicator (required).</param>
        /// <param name="userID">userID (required).</param>
        /// <param name="valid">valid (required).</param>
        /// <param name="spare1">spare1 (required).</param>
        /// <param name="longitude1">longitude1 (required).</param>
        /// <param name="latitude1">latitude1 (required).</param>
        /// <param name="longitude2">longitude2 (required).</param>
        /// <param name="latitude2">latitude2 (required).</param>
        /// <param name="stationType">stationType (required).</param>
        /// <param name="shipType">shipType (required).</param>
        /// <param name="spare2">spare2 (required).</param>
        /// <param name="txRxMode">txRxMode (required).</param>
        /// <param name="reportingInterval">reportingInterval (required).</param>
        /// <param name="quietTime">quietTime (required).</param>
        /// <param name="spare3">spare3 (required).</param>
        public GroupAssignmentCommand(int messageID = default(int), int repeatIndicator = default(int), int userID = default(int), bool valid = default(bool), int spare1 = default(int), double longitude1 = default(double), double latitude1 = default(double), double longitude2 = default(double), double latitude2 = default(double), int stationType = default(int), int shipType = default(int), int spare2 = default(int), int txRxMode = default(int), int reportingInterval = default(int), int quietTime = default(int), int spare3 = default(int))
        {
            this.MessageID = messageID;
            this.RepeatIndicator = repeatIndicator;
            this.UserID = userID;
            this.Valid = valid;
            this.Spare1 = spare1;
            this.Longitude1 = longitude1;
            this.Latitude1 = latitude1;
            this.Longitude2 = longitude2;
            this.Latitude2 = latitude2;
            this.StationType = stationType;
            this.ShipType = shipType;
            this.Spare2 = spare2;
            this.TxRxMode = txRxMode;
            this.ReportingInterval = reportingInterval;
            this.QuietTime = quietTime;
            this.Spare3 = spare3;
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
        /// Gets or Sets Longitude1
        /// </summary>
        [DataMember(Name = "Longitude1", IsRequired = true, EmitDefaultValue = true)]
        public double Longitude1 { get; set; }

        /// <summary>
        /// Gets or Sets Latitude1
        /// </summary>
        [DataMember(Name = "Latitude1", IsRequired = true, EmitDefaultValue = true)]
        public double Latitude1 { get; set; }

        /// <summary>
        /// Gets or Sets Longitude2
        /// </summary>
        [DataMember(Name = "Longitude2", IsRequired = true, EmitDefaultValue = true)]
        public double Longitude2 { get; set; }

        /// <summary>
        /// Gets or Sets Latitude2
        /// </summary>
        [DataMember(Name = "Latitude2", IsRequired = true, EmitDefaultValue = true)]
        public double Latitude2 { get; set; }

        /// <summary>
        /// Gets or Sets StationType
        /// </summary>
        [DataMember(Name = "StationType", IsRequired = true, EmitDefaultValue = true)]
        public int StationType { get; set; }

        /// <summary>
        /// Gets or Sets ShipType
        /// </summary>
        [DataMember(Name = "ShipType", IsRequired = true, EmitDefaultValue = true)]
        public int ShipType { get; set; }

        /// <summary>
        /// Gets or Sets Spare2
        /// </summary>
        [DataMember(Name = "Spare2", IsRequired = true, EmitDefaultValue = true)]
        public int Spare2 { get; set; }

        /// <summary>
        /// Gets or Sets TxRxMode
        /// </summary>
        [DataMember(Name = "TxRxMode", IsRequired = true, EmitDefaultValue = true)]
        public int TxRxMode { get; set; }

        /// <summary>
        /// Gets or Sets ReportingInterval
        /// </summary>
        [DataMember(Name = "ReportingInterval", IsRequired = true, EmitDefaultValue = true)]
        public int ReportingInterval { get; set; }

        /// <summary>
        /// Gets or Sets QuietTime
        /// </summary>
        [DataMember(Name = "QuietTime", IsRequired = true, EmitDefaultValue = true)]
        public int QuietTime { get; set; }

        /// <summary>
        /// Gets or Sets Spare3
        /// </summary>
        [DataMember(Name = "Spare3", IsRequired = true, EmitDefaultValue = true)]
        public int Spare3 { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class GroupAssignmentCommand {\n");
            sb.Append("  MessageID: ").Append(MessageID).Append("\n");
            sb.Append("  RepeatIndicator: ").Append(RepeatIndicator).Append("\n");
            sb.Append("  UserID: ").Append(UserID).Append("\n");
            sb.Append("  Valid: ").Append(Valid).Append("\n");
            sb.Append("  Spare1: ").Append(Spare1).Append("\n");
            sb.Append("  Longitude1: ").Append(Longitude1).Append("\n");
            sb.Append("  Latitude1: ").Append(Latitude1).Append("\n");
            sb.Append("  Longitude2: ").Append(Longitude2).Append("\n");
            sb.Append("  Latitude2: ").Append(Latitude2).Append("\n");
            sb.Append("  StationType: ").Append(StationType).Append("\n");
            sb.Append("  ShipType: ").Append(ShipType).Append("\n");
            sb.Append("  Spare2: ").Append(Spare2).Append("\n");
            sb.Append("  TxRxMode: ").Append(TxRxMode).Append("\n");
            sb.Append("  ReportingInterval: ").Append(ReportingInterval).Append("\n");
            sb.Append("  QuietTime: ").Append(QuietTime).Append("\n");
            sb.Append("  Spare3: ").Append(Spare3).Append("\n");
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