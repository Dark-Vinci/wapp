import { JSX, ReactNode } from 'react';
import { SafeAreaView, View, StyleSheet, Platform } from 'react-native';
import Constants from 'expo-constants';

export interface ScreenInterface {
  readonly children: ReactNode;
}

export function Screen({ children }: ScreenInterface): JSX.Element {
  return (
    <SafeAreaView style={style.container}>
      <View style={style.container_container}>{children}</View>
    </SafeAreaView>
  );
}

const style = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingTop: Platform.OS === 'android' ? Constants.statusBarHeight * 1.2 : 0,
  },

  container_container: {
    width: '100%',
    height: '100%',
    flex: 1,
  },
});
