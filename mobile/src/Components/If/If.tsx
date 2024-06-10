import { JSX, ReactNode } from 'react';

interface IfProps {
  readonly children: ReactNode;
  readonly condition: boolean;
}

export function If({ children, condition }: IfProps): JSX.Element {
  if (condition) {
    return children as JSX.Element;
  }

  return null as unknown as JSX.Element;
}
