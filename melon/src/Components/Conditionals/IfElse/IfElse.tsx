import { JSX, ReactNode } from 'react';

interface IfElseProps {
  condition: boolean;
  ifElement: ReactNode;
  elseElement: ReactNode;
}

export function IfElse({
  condition,
  ifElement,
  elseElement,
}: IfElseProps): JSX.Element {
  if (condition) {
    return ifElement as JSX.Element;
  }

  return elseElement as JSX.Element;
}
