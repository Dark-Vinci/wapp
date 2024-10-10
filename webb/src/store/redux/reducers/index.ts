import { combineReducers } from "redux";
import {
  useSelector as useReduxSelector,
  TypedUseSelectorHook,
} from "react-redux";

import { authReducer } from "./auth";

export const rootReducer = combineReducers({
  auth: authReducer,
});

export const useSelector: TypedUseSelectorHook<ReturnType<typeof rootReducer>> =
  useReduxSelector;
