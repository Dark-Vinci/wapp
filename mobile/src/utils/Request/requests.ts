import { Axios } from 'axios';

import { instantiate } from './axios';
import { RequestURI } from '../contants';

export class Requests {
  private axios: Axios;

  public constructor(token?: string) {
    this.axios = instantiate(token);
  }

  public async ping(): Promise<string> {
    const data = await this.axios.get(RequestURI.PING, {});

    return data as unknown as string;
  }
}
