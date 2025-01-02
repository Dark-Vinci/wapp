
export interface Message<T extends any> {
  message: T;
  type: string;
  userId: string;
  to: string;
  media: number[];
}

export class WsHandler {
  private url: string;
  private socket: WebSocket | null | undefined;
  private listeners: ((m: any) => any)[] = [];

  public constructor(url: string) {
    this.url = url;
    console.log("Websocket constructor");
  }

  public initiateConnection(): void {
    if (this.socket) {
      console.log({ message: 'connection already established' });
      return;
    }

    this.socket = new WebSocket(this.url);

    this.socket!.onopen = (e) => {
      console.log({ message: 'connection established', event: e });
    };

    this.socket.onmessage = (event: MessageEvent) => {
      const data = JSON.parse(event.data);

      this.listeners.forEach((callback) => callback(data));
    };

    this.socket.onclose = (event) => {
      console.log('WebSocket connection closed', event);
      this.socket = null;
    };

    this.socket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };
  }

  public addListener(callback: (m: any) => any): void {
    this.listeners.push(callback);
  }

  public removeListener(callback: (m: any) => any): void {
    this.listeners = this.listeners.filter((listener) => listener !== callback);
  }

  public closeConnection(): void {
    if (this.socket) {
      this.socket.close();
    }
  }

  public sendMessage<T extends {} = {}>(message: Message<T>): void {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(message));
    } else {
      console.warn('Cannot send message: WebSocket is not connected');
    }
  }
}
