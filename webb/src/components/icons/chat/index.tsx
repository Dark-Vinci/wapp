import { BiMessageSquareDetail } from "react-icons/bi";
import { JSX } from "react";

interface chatIconProps {
  // readonly unreadMessageCount: number;
  // readonly lightMode: boolean;
  // readonly active: boolean;
}

export function ChatIcon({}: chatIconProps): JSX.Element {
  return (
    <div>
      <div>
        <BiMessageSquareDetail color="" enableBackground={""} />
      </div>
      <div>{/*<p>{unreadMessageCount}</p>*/}</div>
    </div>
  );
}
