import { createSlice } from '@reduxjs/toolkit';

export interface IAuthState {
  X_AUTH_TOKEN: string;
  username: string;
  permissions: Record<string, string>;
  error: string | null;
  loading: boolean;
}

const initialState: IAuthState = {
  X_AUTH_TOKEN: '',
  username: '',
  permissions: {},
  error: null,
  loading: false,
};

const authSlice = createSlice({
  initialState: initialState,
  name: 'AUTH',
  reducers: {},
});

export const authReducer = authSlice.reducer;
export const authActions = authSlice.actions;
