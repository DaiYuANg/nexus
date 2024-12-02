import {RouterProvider} from "react-router";
import {router} from "./router";

const App = () => {
  console.log(123)
  return (
    <RouterProvider router={router}/>
  );
};

export {App}