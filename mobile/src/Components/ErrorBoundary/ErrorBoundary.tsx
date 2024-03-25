import { JSX, ReactNode } from 'react';
import ErrorBoundary from 'react-native-error-boundary';

import { FallBack } from '../FallBack';

interface ErrorBoundaryInterface {
  readonly children: ReactNode;
}

const errorHandler = (error: Error, stackTrace: string) => {
  console.log({ error, stackTrace });
};

export function AppErrorBoundary({
  children,
}: ErrorBoundaryInterface): JSX.Element {
  return (
    <ErrorBoundary onError={errorHandler} FallbackComponent={FallBack}>
      // @ts-ignore
      {children}
    </ErrorBoundary>
  );
}
