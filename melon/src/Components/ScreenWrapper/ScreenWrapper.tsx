import { ReactNode } from 'react';
import { ScrollView } from 'react-native';

import { style } from './ScreenWrapper.style';

export function ScreenWrapper({ children }: { children: ReactNode }) {
  return <ScrollView style={style.container}>{children}</ScrollView>;
}
