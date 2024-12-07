import { createSlice } from '@reduxjs/toolkit';

export interface IAuthState {
  X_AUTH_TOKEN: string;
  username: string;
  permissions: Record<string, string>;
}

const initialState: IAuthState = {
  X_AUTH_TOKEN: '',
  username: '',
  permissions: {},
};

const authSlice = createSlice({
  initialState: initialState,
  name: 'AUTH',
  reducers: {},
});

export const authReducer = authSlice.reducer;
export const authActions = authSlice.actions;
