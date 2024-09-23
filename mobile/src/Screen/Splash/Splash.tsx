import React, { JSX } from 'react';
import { View, Text, StyleSheet } from 'react-native';

import { Screen } from '@components';

export function Splash(): JSX.Element {
  return (
    <Screen>
      <View style={style.view_container}>
        <Text>Linkedin</Text>
      </View>
    </Screen>
  );
}

const style = StyleSheet.create({
  view_container: {
    flex: 1,
  },
});
