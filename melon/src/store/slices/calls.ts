import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ICall {
  withUser: string;
  type: 'video' | 'audio';
  bound: 'incoming' | 'outgoing';
  isMissed: boolean;
  timestamp: string;
  profileURL: string;
}

export interface ICalls {
  data: ICall[];
}

const initialState: ICalls = {
  data: [],
};

const callsSlice = createSlice({
  initialState: initialState,
  name: 'CALLS',
  reducers: {
    saveCallData(state, action: PayloadAction<any>) {
      state.data = action.payload;
    },
  },
});

export const callsReducer = callsSlice.reducer;
export const callsActions = callsSlice.actions;
