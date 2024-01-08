using System;

namespace aisStream.Client
{
    public class GeoBoundingBox : IEquatable<GeoBoundingBox>
    {
        private readonly GeoCoordinate _corner1;
        private readonly GeoCoordinate _corner2;

        public GeoCoordinate Corner1 { get => _corner1; }
        public GeoCoordinate Corner2 { get => _corner2; }

        public GeoBoundingBox()
        {
            this._corner1 = new GeoCoordinate(-90, -180);
            this._corner2 = new GeoCoordinate(90, 180);
        }

        public GeoBoundingBox(GeoCoordinate corner1, GeoCoordinate corner2)
        {
            this._corner1 = corner1;
            this._corner2 = corner2;
        }

        public override string ToString()
        {
            return string.Format("{0}/{1}", Corner1, Corner2);
        }

        public override bool Equals(Object other)
        {
            return other is GeoBoundingBox && Equals((GeoBoundingBox)other);
        }

        public bool Equals(GeoBoundingBox other)
        {
            return Corner1 == other.Corner1 && Corner2 == other.Corner2;
        }

        public override int GetHashCode()
        {
            return Corner1.GetHashCode() ^ Corner2.GetHashCode();
        }

        internal System.Collections.Generic.List<System.Collections.Generic.List<double>> ToList()
        {
            System.Collections.Generic.List<System.Collections.Generic.List<double>> result = new System.Collections.Generic.List<System.Collections.Generic.List<double>>();
            result.Add(Corner1.ToList());
            result.Add(Corner2.ToList());
            return result;
        }
    }
}
