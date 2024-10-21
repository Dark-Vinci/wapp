import { JSX } from "react";

import style from "./verify.module.scss";
import { useInput } from "@utils";

export function Verify(): JSX.Element {
  const { input: otp, onChange: otpChange } = useInput("");

  const handleOTPSubmission = (): void => {
    console.log({ otp });
  };

  return (
    <div className={style.container}>
      <div className={style.container_}>
        <p>verify your OTP</p>

        <div>
          <div>
            <p>otp</p>
            <input onChange={otpChange} value={otp} placeholder="----" />
          </div>

          <button onClick={handleOTPSubmission}>submit</button>
        </div>
      </div>
    </div>
  );
}
