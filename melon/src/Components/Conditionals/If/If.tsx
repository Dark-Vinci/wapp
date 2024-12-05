import { JSX, ReactNode } from 'react';

interface IIfProps {
  condition: boolean;
  element: ReactNode;
}

export function If({ condition, element }: IIfProps): JSX.Element {
  if (!condition) {
    return null as unknown as JSX.Element;
  }

  return element as JSX.Element;
}
