import { createSlice } from '@reduxjs/toolkit';

export interface ICall {
  withUser: string;
  type: 'video' | 'audio';
  bound: 'incoming' | 'outgoing';
  isMissed: boolean;
  timestamp: string;
  profileURL: string;
}

const initialState: ICall[] = [];

const callsSlice = createSlice({
  initialState: initialState,
  name: 'CALLS',
  reducers: {},
});

export const callsReducer = callsSlice.reducer;
export const callsActions = callsSlice.actions;
