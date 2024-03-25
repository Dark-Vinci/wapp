import { ReactNode, JSX } from 'react';

interface IfElseInterface {
  readonly condition: boolean;
  readonly IfComponent: ReactNode;
  readonly ElseComponent: ReactNode;
}

export function IfElse({
  IfComponent,
  ElseComponent,
  condition,
}: IfElseInterface): JSX.Element {
  if (condition) {
    return IfComponent as JSX.Element;
  }

  return ElseComponent as JSX.Element;
}
