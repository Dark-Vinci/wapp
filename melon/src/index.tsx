import { StatusBar } from 'expo-status-bar';
import {
  Button,
  Dimensions,
  Platform,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import { useState, JSX, useEffect } from 'react';

import { If } from '@components';
import { WS } from '@network';
import { Message, MESSAGES_TYPE } from '@types';

// import { BACKEND_URL } from '@env';

export function Application(): JSX.Element {
  // console.log({ BACKEND_URL });
  const [count, setCount] = useState(0);
  const [messages, setMessages] = useState<any[]>([]);
  const wsManager = new WS('wss://your-websocket-server.com');

  useEffect(() => {
    // Connect WebSocket and add a listener
    wsManager.initiateConnection();

    const handleMessage = (data: any) => {
      setMessages((prev: any[]) => [...prev, data]);
    };

    wsManager.addListener(handleMessage);

    // Cleanup on component unmount
    return () => {
      wsManager.removeListener(handleMessage);
      wsManager.closeConnection();
    };
  }, []);

  const sendMessage = () => {
    const message: Message<any> = {
      X_AUTH_TOKEN: '',
      message: undefined,
      type: MESSAGES_TYPE.MEL,
      userId: '',
    };

    wsManager.sendMessage(message);
  };

  console.log({ messages, count, sendMessage });

  return (
    <If
      condition={false}
      element={
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
      }
    />
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
