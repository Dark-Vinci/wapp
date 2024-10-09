import { JSX, ReactNode } from 'react';

interface IfProps {
  readonly If: ReactNode,
  readonly Else: ReactNode,
  readonly condition: boolean,
}

export function IfElse({ condition, If, Else }: IfProps): JSX.Element {
  if (condition) {
    return If as JSX.Element;
  }
  
  return Else as JSX.Element;
}