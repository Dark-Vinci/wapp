import { JSX } from "react";

import style from "./Menu.module.scss";

import {
  AiIcon,
  ChannelIcon,
  SettingsIcon,
  StatusIcon,
  ChatIcon,
} from "../../icons";

interface MenuProps {
  profileURL: string;
}

export function Menu({ profileURL }: MenuProps): JSX.Element {
  return (
    <div className={style.container}>
      <div className={style.top}>
        <div>
          <ChatIcon />
        </div>
        <div>
          <StatusIcon />
        </div>

        <div>
          <ChannelIcon />
        </div>
        <div>
          <AiIcon />
        </div>
      </div>

      <div className={style.bottom}>
        <div>
          <SettingsIcon />
        </div>
        <div>
          <img src={profileURL} alt="profile image" />
        </div>
      </div>
    </div>
  );
}
