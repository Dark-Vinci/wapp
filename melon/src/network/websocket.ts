export class WS {
  private readonly serverURL: string;

  public constructor(serverURL: string) {
    this.serverURL = serverURL;
  }

  public async initiateConnection(): Promise<void> {
    console.log({ serverURL: this.serverURL });
  }

  public async closeConnection(): Promise<void> {}
}
