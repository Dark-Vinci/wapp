import { combineReducers } from 'redux';
import {
  TypedUseSelectorHook,
  useSelector as useReduxSelector,
} from 'react-redux';

export type AppState = {};

export const rootReducers = combineReducers<AppState>({});

export const useSelector: TypedUseSelectorHook<
  ReturnType<typeof rootReducers>
> = useReduxSelector;
