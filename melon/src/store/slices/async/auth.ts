import { createAsyncThunk } from '@reduxjs/toolkit';

import { AppApi } from '@network';

export class ToolKitAuthAsync {
  public static loginUser(net: AppApi) {
    return createAsyncThunk('user/login', async (userData: any, {}) => {
      try {
        const res = await net.userApi.loginUser(userData);

        return res.data.user;
      } catch (err: unknown) {
        console.log({ err });

        // We got validation errors, let's return those so we can reference in our component and set form errors
        // return rejectWithValue(error.response.data);
      }
    });
  }
}
