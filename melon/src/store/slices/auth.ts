import { createSlice } from '@reduxjs/toolkit';
import { createLoginUser, createUserAccount } from '@/store/slices/async/auth';

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
  extraReducers: (builder) => {
    // login user
    builder
      .addCase(createLoginUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(createLoginUser.fulfilled, (state, action) => {
        state.loading = false;
        state.error = null;
        state.X_AUTH_TOKEN = action.payload;
      })
      .addCase(createLoginUser.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      });

    // create user account
    builder
      .addCase(createUserAccount.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(createUserAccount.fulfilled, (state, action) => {
        state.loading = false;
        state.error = null;
        state.X_AUTH_TOKEN = action.payload;
      })
      .addCase(createUserAccount.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      });
  },
});

export const authReducer = authSlice.reducer;
export const authActions = authSlice.actions;
