import { JSX } from "react";

interface AccountChatProps {
  name: string;
  time: string;
  isPin: boolean;
  onOver: boolean;
  lastFromCurrentUser: boolean;
  lastMessage: string;
  status: string | null;
  messageCount: number;
  imageUrl: string | null;
}

export function AccountChat({
  name,
  time,
  isPin,
  onOver,
  imageUrl,
  lastFromCurrentUser,
  lastMessage,
  messageCount,
}: AccountChatProps): JSX.Element {
  return (
    <div>
      <div className="account-profile">
        {imageUrl ? (
          <img src={imageUrl} alt={`${name} profile`} />
        ) : (
          <div>default image</div>
        )}
      </div>

      <div className={"chat-details"}>
        <div>
          <p>{name}</p>
          <div>{time}</div>
        </div>

        <div>
          {lastFromCurrentUser && <div>checkmark based on status</div>}

          <div
            style={{
              color: lastFromCurrentUser ? "red" : "green",
            }}
          >
            {lastMessage}
          </div>

          {!lastFromCurrentUser && <div>{messageCount}</div>}
          {isPin && <div>pin icon</div>}
          {onOver && <div>below icon</div>}
        </div>
      </div>
    </div>
  );
}
