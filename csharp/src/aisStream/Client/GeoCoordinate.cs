using System;

namespace aisStream.Client
{
    public class GeoCoordinate : IEquatable<GeoCoordinate>
    {
        private readonly double _latitude;
        private readonly double _longitude;

        public double Latitude { get  => _latitude; }
        public double Longitude { get => _longitude; }

        public GeoCoordinate(double latitude, double longitude)
        {
            this._latitude = latitude;
            this._longitude = longitude;
        }

        public override string ToString()
        {
            return string.Format("{0},{1}", Latitude, Longitude);
        }

        public override bool Equals(Object other)
        {
            return other is GeoCoordinate && Equals((GeoCoordinate)other);
        }

        public bool Equals(GeoCoordinate other)
        {
            return Latitude == other.Latitude && Longitude == other.Longitude;
        }

        public override int GetHashCode()
        {
            return Latitude.GetHashCode() ^ Longitude.GetHashCode();
        }

        internal System.Collections.Generic.List<double> ToList()
        {
            System.Collections.Generic.List<double> result = new System.Collections.Generic.List<double>();
            result.Add(Latitude);
            result.Add(Longitude);
            return result;
        }
    }
}
