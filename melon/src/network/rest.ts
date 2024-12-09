import axios from 'axios';

import { BACKEND_URL } from '@env';
import { UserApi } from './userapi';

export class AppApi {
  public static userApi: UserApi = new UserApi(BACKEND_URL, axios.create({}));

  public constructor(serverURL: string) {
    console.log({ serverURL });
  }
}
