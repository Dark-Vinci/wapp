import { ReactNode, JSX } from 'react';

interface IfElseInterface {
  readonly condition: boolean;
  readonly If: ReactNode;
  readonly El: ReactNode;
}

export function IfElse({ If, El, condition }: IfElseInterface): JSX.Element {
  if (condition) {
    return If as JSX.Element;
  }

  return El as JSX.Element;
}
