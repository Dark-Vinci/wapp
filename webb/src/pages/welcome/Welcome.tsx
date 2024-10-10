import style from "./welcome.module.scss";

export function Welcome(): JSX.Element {
  return (
    <div className={style.container}>
      <p>THIS IS THE WELCOME PAGE</p>
    </div>
  );
}
