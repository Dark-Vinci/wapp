import React, { JSX } from 'react';
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';

import { persist, store } from '@store';
import { Splash } from '@screens';
import { If, IfElse } from '@components';

function App(): JSX.Element {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persist}>
        <If children={<div>meme</div>} condition={true} />
        <IfElse
          IfComponent={<div> melon</div>}
          ElseComponent={<div> gret </div>}
          condition={true}
        />
        <Splash />
      </PersistGate>
    </Provider>
  );
}

export default App;
