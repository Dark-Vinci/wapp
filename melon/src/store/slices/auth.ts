import { createSlice } from '@reduxjs/toolkit';

export interface IAuthState {
  X_AUTH_TOKEN: string;
  username: string;
  permissions: Record<string, string>;
}

const authSlice = createSlice({
  initialState: {
    X_AUTH_TOKEN: '',
    username: '',
    permissions: {},
  },
  name: 'AUTH',
  reducers: {},
});

export const authReducer = authSlice.reducer;
export const authActions = authSlice.actions;
