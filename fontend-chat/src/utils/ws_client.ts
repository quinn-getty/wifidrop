const url = `ws://${window.location.hostname}:${location.port}/ws`;
const wsClient = new WebSocket(url);

class WsClient {
  client: WebSocket
  constructor(client:WebSocket) {
    this.client = client
  }
  send(data:Record<string,string>) {
    this.client.send(JSON.stringify(data))
  }
  onMessage(fn: (data:Record<string,string>)=>void) {
    this.client.onmessage = ({ data }) => {
      fn(JSON.parse(data))
    }
  }
}


const promise = new Promise<WsClient>((resolve, reject) => {
  wsClient.onopen = () => {
    resolve(new WsClient(wsClient))
  }
  setTimeout(() => {
    reject(new Error('get ws connection timeout'))
  }, 10000)
})

export const getWsClient = () => promise