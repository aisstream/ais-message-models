namespace aisStream.Client
{
    public class AisReceivedEventArgs
    {
        public AisReceivedEventArgs(object ais, object meta)
        {
            AIS = ais;
            Meta = meta;
        }

        public object AIS { get; }
        public object Meta { get; }
    }
}
