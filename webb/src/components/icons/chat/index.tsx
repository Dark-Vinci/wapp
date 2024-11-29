import { BiMessageSquareDetail } from "react-icons/bi";

interface chatIconProps {
  readonly unreadMessageCount: number;
  readonly lightMode: boolean;
  readonly active: boolean;
}

export function ChatIcon({ unreadMessageCount }: chatIconProps): JSX.Element {
  return (
    <div>
      <div>
        <BiMessageSquareDetail color="" enableBackground={""} />
      </div>
      <div>
        <p>{unreadMessageCount}</p>
      </div>
    </div>
  );
}
