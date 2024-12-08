import { createSlice } from '@reduxjs/toolkit';

import { AppApi } from '@network';
import { ToolKitAuthAsync } from '@/store/slices/async/auth';

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
// let kt = null;

const authSlice = createSlice({
  initialState: initialState,
  name: 'AUTH',
  reducers: {},
  extraReducers: (builder) => {
    const api = new AppApi('');
    const tk = ToolKitAuthAsync.loginUser(api);
    // kt = tk;

    builder.addCase(tk.pending, (state) => {
      state.loading = true;
      state.error = null;
    });

    builder.addCase(tk.fulfilled, (state, action) => {
      state.loading = false;
      state.error = null;
      state.X_AUTH_TOKEN = action.payload;
    });

    builder.addCase(tk.rejected, (state, action) => {
      state.loading = false;
      state.error = action.payload as string;
    });
  },
});

export const authReducer = authSlice.reducer;
export const authActions = authSlice.actions;
// export const loginAction = kt;

// let a = authActions.
