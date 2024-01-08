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
    /// BaseStationReport
    /// </summary>
    [DataContract(Name = "BaseStationReport")]
    public partial class BaseStationReport : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="BaseStationReport" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected BaseStationReport() { }
        /// <summary>
        /// Initializes a new instance of the <see cref="BaseStationReport" /> class.
        /// </summary>
        /// <param name="messageID">messageID (required).</param>
        /// <param name="repeatIndicator">repeatIndicator (required).</param>
        /// <param name="userID">userID (required).</param>
        /// <param name="valid">valid (required).</param>
        /// <param name="utcYear">utcYear (required).</param>
        /// <param name="utcMonth">utcMonth (required).</param>
        /// <param name="utcDay">utcDay (required).</param>
        /// <param name="utcHour">utcHour (required).</param>
        /// <param name="utcMinute">utcMinute (required).</param>
        /// <param name="utcSecond">utcSecond (required).</param>
        /// <param name="positionAccuracy">positionAccuracy (required).</param>
        /// <param name="longitude">longitude (required).</param>
        /// <param name="latitude">latitude (required).</param>
        /// <param name="fixType">fixType (required).</param>
        /// <param name="longRangeEnable">longRangeEnable (required).</param>
        /// <param name="spare">spare (required).</param>
        /// <param name="raim">raim (required).</param>
        /// <param name="communicationState">communicationState (required).</param>
        public BaseStationReport(int messageID = default(int), int repeatIndicator = default(int), int userID = default(int), bool valid = default(bool), int utcYear = default(int), int utcMonth = default(int), int utcDay = default(int), int utcHour = default(int), int utcMinute = default(int), int utcSecond = default(int), bool positionAccuracy = default(bool), double longitude = default(double), double latitude = default(double), int fixType = default(int), bool longRangeEnable = default(bool), int spare = default(int), bool raim = default(bool), int communicationState = default(int))
        {
            this.MessageID = messageID;
            this.RepeatIndicator = repeatIndicator;
            this.UserID = userID;
            this.Valid = valid;
            this.UtcYear = utcYear;
            this.UtcMonth = utcMonth;
            this.UtcDay = utcDay;
            this.UtcHour = utcHour;
            this.UtcMinute = utcMinute;
            this.UtcSecond = utcSecond;
            this.PositionAccuracy = positionAccuracy;
            this.Longitude = longitude;
            this.Latitude = latitude;
            this.FixType = fixType;
            this.LongRangeEnable = longRangeEnable;
            this.Spare = spare;
            this.Raim = raim;
            this.CommunicationState = communicationState;
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
        /// Gets or Sets UtcYear
        /// </summary>
        [DataMember(Name = "UtcYear", IsRequired = true, EmitDefaultValue = true)]
        public int UtcYear { get; set; }

        /// <summary>
        /// Gets or Sets UtcMonth
        /// </summary>
        [DataMember(Name = "UtcMonth", IsRequired = true, EmitDefaultValue = true)]
        public int UtcMonth { get; set; }

        /// <summary>
        /// Gets or Sets UtcDay
        /// </summary>
        [DataMember(Name = "UtcDay", IsRequired = true, EmitDefaultValue = true)]
        public int UtcDay { get; set; }

        /// <summary>
        /// Gets or Sets UtcHour
        /// </summary>
        [DataMember(Name = "UtcHour", IsRequired = true, EmitDefaultValue = true)]
        public int UtcHour { get; set; }

        /// <summary>
        /// Gets or Sets UtcMinute
        /// </summary>
        [DataMember(Name = "UtcMinute", IsRequired = true, EmitDefaultValue = true)]
        public int UtcMinute { get; set; }

        /// <summary>
        /// Gets or Sets UtcSecond
        /// </summary>
        [DataMember(Name = "UtcSecond", IsRequired = true, EmitDefaultValue = true)]
        public int UtcSecond { get; set; }

        /// <summary>
        /// Gets or Sets PositionAccuracy
        /// </summary>
        [DataMember(Name = "PositionAccuracy", IsRequired = true, EmitDefaultValue = true)]
        public bool PositionAccuracy { get; set; }

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
        /// Gets or Sets FixType
        /// </summary>
        [DataMember(Name = "FixType", IsRequired = true, EmitDefaultValue = true)]
        public int FixType { get; set; }

        /// <summary>
        /// Gets or Sets LongRangeEnable
        /// </summary>
        [DataMember(Name = "LongRangeEnable", IsRequired = true, EmitDefaultValue = true)]
        public bool LongRangeEnable { get; set; }

        /// <summary>
        /// Gets or Sets Spare
        /// </summary>
        [DataMember(Name = "Spare", IsRequired = true, EmitDefaultValue = true)]
        public int Spare { get; set; }

        /// <summary>
        /// Gets or Sets Raim
        /// </summary>
        [DataMember(Name = "Raim", IsRequired = true, EmitDefaultValue = true)]
        public bool Raim { get; set; }

        /// <summary>
        /// Gets or Sets CommunicationState
        /// </summary>
        [DataMember(Name = "CommunicationState", IsRequired = true, EmitDefaultValue = true)]
        public int CommunicationState { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class BaseStationReport {\n");
            sb.Append("  MessageID: ").Append(MessageID).Append("\n");
            sb.Append("  RepeatIndicator: ").Append(RepeatIndicator).Append("\n");
            sb.Append("  UserID: ").Append(UserID).Append("\n");
            sb.Append("  Valid: ").Append(Valid).Append("\n");
            sb.Append("  UtcYear: ").Append(UtcYear).Append("\n");
            sb.Append("  UtcMonth: ").Append(UtcMonth).Append("\n");
            sb.Append("  UtcDay: ").Append(UtcDay).Append("\n");
            sb.Append("  UtcHour: ").Append(UtcHour).Append("\n");
            sb.Append("  UtcMinute: ").Append(UtcMinute).Append("\n");
            sb.Append("  UtcSecond: ").Append(UtcSecond).Append("\n");
            sb.Append("  PositionAccuracy: ").Append(PositionAccuracy).Append("\n");
            sb.Append("  Longitude: ").Append(Longitude).Append("\n");
            sb.Append("  Latitude: ").Append(Latitude).Append("\n");
            sb.Append("  FixType: ").Append(FixType).Append("\n");
            sb.Append("  LongRangeEnable: ").Append(LongRangeEnable).Append("\n");
            sb.Append("  Spare: ").Append(Spare).Append("\n");
            sb.Append("  Raim: ").Append(Raim).Append("\n");
            sb.Append("  CommunicationState: ").Append(CommunicationState).Append("\n");
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