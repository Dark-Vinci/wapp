import React, { JSX } from 'react';
import { Provider } from 'react-redux';
import { View } from 'react-native';
import { PersistGate } from 'redux-persist/integration/react';

import { persist, store } from '@store';
import { Splash } from '@screens';
import { If, IfElse } from '@components';

function App(): JSX.Element {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persist}>
        <If children={<View>meme</View>} condition={true} />
        <IfElse
          If={<View> melon</View>}
          El={<View> gret </View>}
          condition={true}
        />
        <Splash />
      </PersistGate>
    </Provider>
  );
}

export default App;
