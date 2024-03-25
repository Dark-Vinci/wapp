import { createRef, RefObject } from 'react';
import {
  NavigationAction,
  NavigationContainerRef,
  StackActions,
} from '@react-navigation/native';

export const navRef: RefObject<NavigationContainerRef<any>> = createRef();

function navigate(name: string, params?: object): void {
  navRef?.current?.navigate(name, params);
}

function dispatch(action: NavigationAction): void {
  navRef?.current?.dispatch(action);
}

function goBack(): void {
  navRef?.current?.goBack();
}

function push(name: string, param?: object): void {
  navRef?.current?.dispatch(StackActions.push(name, param));
}

function replace(name: string, param?: object): void {
  navRef?.current?.dispatch(StackActions.replace(name, param));
}

export const nav = {
  navigate,
  dispatch,
  goBack,
  push,
  replace,
};
