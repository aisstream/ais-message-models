using System;

namespace aisStream.ModelExample
{
    internal class Program
    {
        static void Main(string[] args)
        {
            using (var client = new Client.WebsocketClient("<YOUR API KEY>", new Client.GeoBoundingBox(new Client.GeoCoordinate(-11.0, 178.0), new Client.GeoCoordinate(30.0, 74.0))))
            {
                client.AisReceived += Client_AisReceived;
                client.Start();

                Console.WriteLine("Press any key to close ...");
                Console.ReadKey();
            }
        }

        private static void Client_AisReceived(object sender, Client.AisReceivedEventArgs e)
        {
            Console.WriteLine(e.AIS);
        }
    }
}
