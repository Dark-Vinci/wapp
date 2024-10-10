import { Auth } from "../actions";

export interface AuthState {
  authToken: string;
  userId: string;
}

const initialSate: AuthState = {
  authToken: "",
  userId: "",
};

export interface AuthAction {
  type: Auth;
  payload: any; //TODO: update accordingly
}

export function authReducer(
  state: AuthState = initialSate,
  action: AuthAction,
): AuthState {
  switch (action.type) {
    case Auth.LOGIN:
      return state;
    default:
      return state;
  }
}
