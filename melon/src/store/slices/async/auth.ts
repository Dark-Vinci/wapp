import { createAsyncThunk } from '@reduxjs/toolkit';

import { AppApi } from '@network';

export const createLoginUser = createAsyncThunk(
  'user/login',
  async (userData: any, {}) => {
    try {
      const res = await AppApi.userApi.loginUser(userData);

      return res.data.user;
    } catch (err: unknown) {
      console.log({ err });
    }
  },
);

export const createUserAccount = createAsyncThunk(
  'user/register',
  async (data: any, {}) => {
    try {
      const res = await AppApi.userApi.createAccount(data);

      return res.data.user;
    } catch (err: unknown) {
      console.log({ err });
    }
  },
);
