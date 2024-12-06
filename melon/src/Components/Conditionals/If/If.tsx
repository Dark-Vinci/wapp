import { JSX, ReactNode } from 'react';

interface IfProps {
  condition: boolean;
  element: ReactNode;
}

export function If({ condition, element }: IfProps): JSX.Element {
  if (!condition) {
    return null as unknown as JSX.Element;
  }

  return element as JSX.Element;
}
