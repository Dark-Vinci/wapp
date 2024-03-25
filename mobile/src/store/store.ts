import thunk from 'redux-thunk';
import { applyMiddleware, compose, legacy_createStore } from 'redux';
import { persistStore, persistReducer } from 'redux-persist';
import AsyncStorage from 'async-storage';

import {rootReducers} from './reducers';

const persistConfig = {
  key: 'root',
  storage: AsyncStorage,
  whitelist: ['user'],
}


const persistedReducer = persistReducer(persistConfig, rootReducers);

const middleware = [thunk];
// @ts-ignore
const enhancers  = [applyMiddleware(...middleware)];

// @ts-ignore
export const store = legacy_createStore(persistedReducer, compose(...enhancers));

export const persist  = persistStore(store);