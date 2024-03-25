import { JSX } from 'react';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';

import { persist, store } from '@store';
import { Splash } from '@screens';

function App(): JSX.Element {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persist}>
        <Splash />
      </PersistGate>
    </Provider>
  );
}

export default App;
