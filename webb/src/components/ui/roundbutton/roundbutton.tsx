import { JSX } from "react";

interface Props {
  content: string;
  isActive: boolean;
}

export function RoundButtonText({ content }: Props): JSX.Element {
  return (
    <div style={{ width: "auto", height: "30px", borderRadius: "50%" }}>
      {content}
    </div>
  );
}
