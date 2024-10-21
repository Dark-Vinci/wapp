import { JSX } from "react";

import style from "./login.module.scss";
import { useInput } from "@utils";

export function Login(): JSX.Element {
  const { input: password, onChange: passwordChange } = useInput("");
  const { input: phoneNumber, onChange: phoneNumberChange } = useInput("");

  const submitHandler = (): void => {
    console.log({ password, phoneNumber });
  };

  return (
    <div className={style.container}>
      <div>
        <div className={style.left}>
          <img src={""} />
        </div>

        <div className={style.right_container}>
          <div className={style.right}>
            {/*  welcome */}
            <div>
              <div className={style.top_container}>
                <p>Welcome!!!</p>
                <p>Log in to your whatsapp account</p>
              </div>
            </div>

            {/*form*/}
            <div>
              <div>
                <div>
                  <p>phone number</p>
                  <div>
                    <input
                      type="string"
                      placeholder="phone number"
                      onChange={phoneNumberChange}
                      value={phoneNumber}
                    />
                  </div>
                </div>
                <div>
                  <p>password</p>
                  <div>
                    <input
                      type="password"
                      placeholder="******"
                      onChange={passwordChange}
                      value={password}
                    />
                    {/*    eye icon */}
                  </div>
                </div>

                <button onClick={submitHandler}>Log in</button>
              </div>

              <div>
                {/*  add link to change password */}
                <a href={"."}>Forgot password?</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
