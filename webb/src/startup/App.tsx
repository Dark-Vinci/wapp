import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { JSX } from "react";

import { If } from "@components";
import "./App.scss";

export const router = createBrowserRouter([
  {
    path: "/login",
    element: <div>THIS IS THE LOGIN ELEMENT</div>,
    children: [],
  },
  {
    path: "/signup",
    element: <div>THIS IS THE SIGNUP ELEMENT</div>,
    children: [],
  },
  {
    path: "/otp",
    element: <div>THIS IS THE OTP ELEMENT</div>,
    children: [],
  },
  {
    path: "/app",
    element: <div>THIS IS THE DEFAULT ELEMENT</div>,
    children: [],
  },

  {
    path: "*",
    element: <div>THIS IS PAGE 404</div>,
  },
]);

function Def(): JSX.Element {
  return (
    <div className="App">
      <p>WELCOME TO WHATSAPP</p>
      <If element={<i>THIS IS THE I CONTENT</i>} condition={true} />
    </div>
  );
}

export function App() {
  return <RouterProvider router={router} fallbackElement={<Def />} />;
}
