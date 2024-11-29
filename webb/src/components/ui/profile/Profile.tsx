import { ChangeEventHandler, JSX } from "react";

import { CheckIcon, EditIcon } from "@/components/icons";

import style from "./Profile.module.scss";

interface ProfileProps {
  editAbout: boolean;
  aboutValue: string;
  userName: string;
  editUsername: boolean;
  profileURL: string;
  onUsernameChange: ChangeEventHandler<HTMLInputElement>;
  onAboutChange: ChangeEventHandler<HTMLInputElement>;
}

export function Profile({
  editAbout,
  aboutValue,
  userName,
  editUsername,
  profileURL,
  onUsernameChange,
  onAboutChange,
}: ProfileProps): JSX.Element {
  let about = (
    <div>
      <div>
        <p>About</p>
      </div>
      <div>
        <div>{aboutValue}</div>
        <EditIcon />
      </div>
    </div>
  );

  let username = (
    <div>
      <div>
        <p>Your name</p>
      </div>

      <div>
        <p>{aboutValue}</p>
        <EditIcon />
      </div>
    </div>
  );

  if (editUsername) {
    username = (
      <div>
        <div>
          <p>Your name</p>
        </div>

        <div>
          <input type={"text"} value={userName} onChange={onUsernameChange} />
          <div>8</div>
          <div>emoji icon</div>

          <div>
            <CheckIcon />
          </div>
        </div>
      </div>
    );
  }

  if (editAbout) {
    about = (
      <div>
        <div>
          <p>About</p>
        </div>
        <div>
          <input
            placeholder="About..."
            value={aboutValue}
            onChange={onAboutChange}
          />
          <div>
            <CheckIcon />
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className={style.container}>
      <div className="profile">
        <p>Profile</p>
      </div>

      <div className="profile_img">
        <img src={profileURL} alt="profile image" />
        <div>
          <div>icon photo</div>
          <p>CHANGE YOUR PROFILE PHOTO</p>
        </div>
      </div>

      <div>
        <div className="details_container">
          <div>
            <div>{username}</div>
            <div>
              <p>
                This is not your username or pin. The name will be visible to
                your WhatsApp contacts
              </p>
            </div>
          </div>

          {/*about; either input or display*/}
          <div>{about}</div>
        </div>
      </div>
    </div>
  );
}
