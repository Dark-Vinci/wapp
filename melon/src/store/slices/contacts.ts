import { createSlice } from '@reduxjs/toolkit';

enum ContactType {
  PERSON = 'PERSON',
  GROUP = 'GROUP',
  CHANNEL = 'CHANNEL',
}

export interface contactStructure {
  name: string;
  lastMessage: string;
  profileURL: string;
  hasActiveStory: boolean;
  hasUnseenStory: boolean;
  type: ContactType;
}

const initial: contactStructure[] = [];

const contactSlice = createSlice({
  initialState: initial,
  name: 'CONTACTS',
  reducers: {},
});

export const contactActions = contactSlice.actions;
export const contactReducers = contactSlice.reducer;
