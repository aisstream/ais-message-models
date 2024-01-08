using aisStream.Model;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace aisStream.Client
{
    public partial class WebsocketClient : IDisposable
    {
        private bool _disposing;
        private bool _stopping;
        private ClientWebSocket? _client;
        private CancellationTokenSource? _cancellation;
        private CancellationTokenSource? _cancellationTotal;
        private GeoBoundingBox _boundingBox;
        private string _apiKey;

        public event EventHandler<AisReceivedEventArgs> AisReceived;

        public WebsocketClient(string apiKey, GeoBoundingBox boundingBox = null)
        {
            _apiKey = apiKey;
            _boundingBox = boundingBox == null ? new GeoBoundingBox() : boundingBox;
        }

        public Uri Url
        {
            get => new Uri("wss://stream.aisstream.io/v0/stream");
        }

        public bool IsStarted { get; private set; }

        public bool IsRunning { get; private set; }

        public ClientWebSocket? NativeClient
        {
            get
            {
                if (_client == null)
                    return null;
                var specific = _client as ClientWebSocket;
                if (specific == null)
                    throw new Exception("Cannot cast 'WebSocket' client to 'ClientWebSocket', provide correct type via factory or don't use this property at all.");
                return specific;
            }
        }

        public void Dispose()
        {
            _disposing = true;
            try
            {
                _cancellation?.Cancel();
                _cancellationTotal?.Cancel();
                _client?.Abort();
                _client?.Dispose();
                _cancellation?.Dispose();
                _cancellationTotal?.Dispose();
            }
            catch (Exception e)
            {
            }
            if (IsRunning)
            IsRunning = false;
            IsStarted = false;
        }

        public Task Start()
        {
            return StartInternal();
        }

        public async Task<bool> Stop(WebSocketCloseStatus status, string statusDescription)
        {
            var result = await StopInternal(_client, status, statusDescription, null, false, false).ConfigureAwait(false);
            return result;
        }

        public async Task<bool> StopOrFail(WebSocketCloseStatus status, string statusDescription)
        {
            var result = await StopInternal(_client, status, statusDescription, null, true, false).ConfigureAwait(false);
            return result;
        }

        private async Task StartInternal()
        {
            if (_disposing)
            {
                throw new Exception("Client is already disposed, starting not possible");
            }

            if (IsStarted) return;
            IsStarted = true;

            _cancellation = new CancellationTokenSource();
            _cancellationTotal = new CancellationTokenSource();
            await StartClient(_cancellation.Token).ConfigureAwait(false);
        }

        private async Task<bool> StopInternal(WebSocket? client, WebSocketCloseStatus status, string statusDescription, CancellationToken? cancellation, bool failFast, bool byServer)
        {
            if (_disposing)
            {
                throw new Exception("Client is already disposed, stopping not possible");
            }

            if (client == null)
            {
                IsStarted = false;
                IsRunning = false;
                return false;
            }

            if (!IsRunning)
            {
                IsStarted = false;
                return false;
            }

            var result = false;
            try
            {
                var cancellationToken = cancellation ?? CancellationToken.None;
                _stopping = true;
                if (byServer)
                    await client.CloseOutputAsync(status, statusDescription, cancellationToken);
                else
                    await client.CloseAsync(status, statusDescription, cancellationToken);
                result = true;
            }
            catch (Exception e)
            {
                if (failFast)
                {
                    throw new Exception($"Failed to stop Websocket client, error: '{e.Message}'", e);
                }
            }
            finally
            {
                IsRunning = false;
                _stopping = false;
                if (!byServer) IsStarted = false;
            }

            return result;
        }

        private async Task StartClient(CancellationToken token)
        {
            try
            {
                _client = new ClientWebSocket();
                await _client.ConnectAsync(Url, token).ConfigureAwait(false);
                var sub = new Model.SubscriptionMessage(_apiKey, new List<List<List<double>>>(new List<List<double>>[] { _boundingBox.ToList() }));
                await _client.SendAsync(new ArraySegment<byte>(Encoding.UTF8.GetBytes(sub.ToJson())), WebSocketMessageType.Text, true, token);
                _ = Listen(_client, token);
                IsRunning = true;
                IsStarted = true;
            }
            catch
            {
                IsRunning = false;
                IsStarted = false;
            }
        }

        private async Task Listen(WebSocket client, CancellationToken token)
        {
            try
            {
                var buffer = new ArraySegment<byte>(new byte[4096]);
                do
                {
                    WebSocketReceiveResult result;
                    byte[]? resultArrayWithTrailing = null;
                    var resultArraySize = 0;
                    var isResultArrayCloned = false;
                    MemoryStream? ms = null;

                    while (true)
                    {
                        result = await client.ReceiveAsync(buffer, token);
                        var currentChunk = buffer.Array;
                        var currentChunkSize = result.Count;

                        var isFirstChunk = resultArrayWithTrailing == null;
                        if (isFirstChunk)
                        {
                            resultArraySize += currentChunkSize;
                            resultArrayWithTrailing = currentChunk;
                            isResultArrayCloned = false;
                        }
                        else if (currentChunk == null)
                        {
                        }
                        else
                        {
                            if (ms == null)
                            {
                                ms = new MemoryStream();
                                ms.Write(resultArrayWithTrailing!, 0, resultArraySize);
                            }
                            ms.Write(currentChunk, buffer.Offset, currentChunkSize);
                        }

                        if (result.EndOfMessage)
                        {
                            break;
                        }

                        if (isResultArrayCloned)
                            continue;

                        resultArrayWithTrailing = resultArrayWithTrailing?.ToArray();
                        isResultArrayCloned = true;
                    }

                    ms?.Seek(0, SeekOrigin.Begin);

                    if (result.MessageType == WebSocketMessageType.Close)
                    {
                        if (!IsStarted || _stopping)
                        {
                            return;
                        }
                        await StopInternal(client, WebSocketCloseStatus.NormalClosure, "Closing", token, false, true);
                        return;
                    }
                    else
                    {
                        var json = ms != null ? Encoding.UTF8.GetString(ms.ToArray()) : resultArrayWithTrailing != null ? Encoding.UTF8.GetString(resultArrayWithTrailing, 0, resultArraySize) : null;
                        if (json != null && AisReceived != null)
                        {
                            try
                            {
                                AisStreamMessage aisStreamMessage = JsonConvert.DeserializeObject<AisStreamMessage>(json);
                                object ais = null;
                                string mt = Enum.GetName(typeof(AisMessageTypes), aisStreamMessage.MessageType);
                                System.Reflection.PropertyInfo prop = typeof(AisStreamMessageMessage).GetProperty(mt);
                                if (prop != null)
                                {
                                    ais = prop.GetValue(aisStreamMessage.Message, null);
                                    AisReceived(this, new AisReceivedEventArgs(ais, aisStreamMessage.MetaData));
                                }
                            }
                            catch
                            {
                            }
                        }
                    }
                    ms?.Dispose();
                } while (client.State == WebSocketState.Open && !token.IsCancellationRequested);
            }
            catch
            {
            }
        }
    }
}
