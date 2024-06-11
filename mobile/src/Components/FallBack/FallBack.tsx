import React, { TouchableOpacity, Text, View } from 'react-native';
import { JSX } from 'react';

import { Screen } from '@components';

export function FallBack(): JSX.Element {
  return (
    <Screen>
      <View>SOMETHING WENT WRONG</View>
      <TouchableOpacity>
        <Text>Go Back</Text>
      </TouchableOpacity>
    </Screen>
  );
}
