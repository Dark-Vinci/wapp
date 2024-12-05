import { StatusBar } from 'expo-status-bar';
import {
  Button,
  Dimensions,
  Platform,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import { useState, JSX } from 'react';

export function Application(): JSX.Element {
  const [count, setCount] = useState(0);

  return (
    <View style={styles.container}>
      <Text>Open up App.tsx to start working on your app!</Text>
      <StatusBar style="auto" />
      <Text>The value is: {count}</Text>
      <Text>Platform: {Platform.OS}</Text>
      <Text>
        HEIGHT: {Dimensions.get('window').height}: WIDTH:
        {Dimensions.get('window').width}
      </Text>
      <Button onPress={() => setCount(count + 1)} title={'The content'} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
