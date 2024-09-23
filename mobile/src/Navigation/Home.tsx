import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

export function Home(): JSX.Element {
  return (
    <View style={style.container}>
      <Text>This is a text for the Home page</Text>
    </View>
  );
}

const style = StyleSheet.create({
  container: {
    borderWidth: 2,
    borderRadius: 20,
  },
});
