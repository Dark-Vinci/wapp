import { MESSAGES_TYPE } from './enums';

export interface Message<T extends any> {
  message: T;
  type: MESSAGES_TYPE;
  X_AUTH_TOKEN: string;
  userId: string;
}
