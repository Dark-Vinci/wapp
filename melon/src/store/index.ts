import { configureStore } from '@reduxjs/toolkit';
import { thunk } from 'redux-thunk';

import { authReducer } from './slices';

const store = configureStore({
  reducer: {
    auth: authReducer,
    // and many more actions
  },
  middleware: (getDefaultMiddleware) => {
    return getDefaultMiddleware().concat(thunk);
  },
});

export default store;
