import React, { JSX } from 'react';
import { View, Text, StyleSheet } from 'react-native';

export function Navigation(): JSX.Element {
  return (
    <View style={style.container}>
      <Text>This is text content</Text>
    </View>
  );
}

const style = StyleSheet.create({
  container: {},
});
