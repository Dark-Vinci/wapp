import React, { JSX } from 'react';
import { View, Text, StyleSheet } from 'react-native';

export function Auth(): JSX.Element {
  return (
    <View style={style.container}>
      <Text>This is the auth page</Text>
    </View>
  );
}

const style = StyleSheet.create({
  container: {},
});
