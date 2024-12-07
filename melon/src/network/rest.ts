// import { createAsyncThunk } from '@reduxjs/toolkit';
// import axios, { AxiosError } from 'axios';
import { UserApi } from '@/network/userapi';
import axios from 'axios';

export class AppApi {
  public userApi: UserApi;

  public constructor(serverURL: string) {
    const axiosInstance = axios.create({});

    this.userApi = new UserApi(serverURL, axiosInstance);
  }
}
