import { configureStore } from '@reduxjs/toolkit';
import { thunk } from 'redux-thunk';

import {
  authReducer,
  callsReducer,
  contactReducers,
  activeChatReducer,
} from './slices';

const store = configureStore({
  reducer: {
    auth: authReducer,
    contacts: contactReducers,
    calls: callsReducer,
    activeChat: activeChatReducer,
  },
  middleware: (getDefaultMiddleware) => {
    return getDefaultMiddleware().concat(thunk);
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export type AppStore = typeof store;

export default store;
