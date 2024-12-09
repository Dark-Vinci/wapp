import { createSlice } from '@reduxjs/toolkit';

export interface IUser {
  profileURL: string;
  displayName: string;
  lastName: string;
  isOnline: boolean;
}

export interface IMessages {
  stringContent: string;
  timestamp: string;
  delivered: boolean;
  seen: boolean;
  sent: boolean;
}

export interface IActiveChat {
  user: IUser | null;
  messages: IMessages[];
}

const initialState: IActiveChat = {
  user: null,
  messages: [],
};

const activeChatSlice = createSlice({
  initialState: initialState,
  name: 'ACTIVE_CHAT',
  reducers: {},
});

export const activeChatReducer = activeChatSlice.reducer;
export const activeChatActions = activeChatSlice.actions;
