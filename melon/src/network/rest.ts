import axios from 'axios';

import { UserApi } from './userapi';

export class AppApi {
  public userApi: UserApi;

  public constructor(serverURL: string) {
    const axiosInstance = axios.create({ baseURL: serverURL });

    this.userApi = new UserApi(serverURL, axiosInstance);
  }
}
