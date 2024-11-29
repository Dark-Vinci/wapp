import React, { JSX, useState } from "react";

import style from "./chats.module.scss";
import {
  AddChatIcon,
  BackArrowIcon,
  ClearIcon,
  GetForDevice,
  MenuIcon,
  RoundButtonText,
  SearchIcon,
} from "@components";
import { AccountChat } from "@/components/ui/accountchat/AccountChat";

interface ChatElement {
  userName: string;
  id: string;
  time: string;
  isPin: boolean;
  type: "group" | "favourite" | "unread" | "all";
}

export function Chats(): JSX.Element {
  const [search, setSearch] = useState<string>("");
  const [chats, setChats] = useState<ChatElement[]>([]);
  const [page, _setPage] = useState([
    { value: "All", isActive: true },
    { value: "Unread", isActive: false },
    { value: "Favourites", isActive: false },
    { value: "Groups", isActive: false },
  ]);

  const setHoverState = (id: string) => {
    setChats((prevItems) =>
      prevItems.map((item) => {
        if (item.id == id) {
          return { ...item, isHover: true };
        } else {
          return { ...item, isHover: false };
        }
      }),
    );
  };

  const searchChangeHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(e.currentTarget.value);
  };

  return (
    <div className={style.container}>
      <div className={style.top}>
        <div className={style.header}>
          <p>Chats</p>

          <div className={"right"}>
            <div>
              <AddChatIcon />
            </div>

            <div>
              <MenuIcon />
            </div>
          </div>
        </div>

        <div className={style.input}>
          {search.length > 0 ? (
            <div>
              <SearchIcon />
            </div>
          ) : (
            <div>
              <BackArrowIcon />
            </div>
          )}

          <input
            value={search}
            placeholder="Search..."
            onChange={searchChangeHandler}
          />

          {search.length > 0 && (
            <div>
              <ClearIcon />
            </div>
          )}
        </div>

        {/*  THIS IS FOR THE PAGES ALL, FAV, GROUP*/}
        <div>
          {page.map(({ value, isActive }, i) => {
            return (
              <div key={i}>
                <RoundButtonText content={value} isActive={isActive} />
              </div>
            );
          })}
        </div>
      </div>

      <div>
        <div>
          <div>archive icon</div>
          <p>Archived</p>
        </div>

        {/*  chat element */}
        <div className={"chat-list"}>
          {chats.map(({ id, userName, time, isPin }: ChatElement) => (
            <div key={id} onMouseEnter={() => setHoverState(id)}>
              <AccountChat
                name={userName}
                time={time}
                isPin={isPin}
                onOver={false}
                lastFromCurrentUser={false}
                lastMessage={""}
                status={"pending"}
                messageCount={3}
                imageUrl={""}
              />
            </div>
          ))}
        </div>
      </div>

      <div>
        <GetForDevice />
      </div>
    </div>
  );
}
