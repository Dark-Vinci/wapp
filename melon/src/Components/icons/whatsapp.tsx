import FontAwesome from '@expo/vector-icons/FontAwesome';
import { View } from 'react-native';
import { JSX } from 'react';

export function WhatsappIcon(): JSX.Element {
  return (
    <View style={{ backgroundColor: 'green' }}>
      <FontAwesome name="whatsapp" size={24} color="white" />
    </View>
  );
}
