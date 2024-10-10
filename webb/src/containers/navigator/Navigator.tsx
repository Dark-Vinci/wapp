import style from "./navigator.module.scss";

export function Navigator(): JSX.Element {
  return (
    <div className={style.container}>
      <div className={style.topContent}>
        <div>chat</div>
        <div>status</div>
        <div>channel</div>
        <div>GPT</div>
      </div>

      <div className={style.lowerContent}>
        <div>SETTINGS</div>
        <div>PROFILE</div>
      </div>
    </div>
  );
}
