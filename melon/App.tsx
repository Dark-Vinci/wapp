import { JSX } from 'react';
import { Provider } from 'react-redux';

import store from '@/store';
import { Application } from '@/index';

export default function App(): JSX.Element {
  return (
    <Provider store={store}>
      <Application />
    </Provider>
  );
}
