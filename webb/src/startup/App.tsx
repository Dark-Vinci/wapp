import { If } from "@components";
import "./App.scss";

export function App() {
  return (
    <div className="App">
      <p>WELCOME TO WHATSAPP</p>
      <If element={<i>THIS IS THE I CONTENT</i>} condition={true} />
    </div>
  );
}
