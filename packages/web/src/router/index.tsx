import {createBrowserRouter} from "react-router";
import {Login} from "../page/auth/Login.tsx";
import {Layout} from "../component/Layout.tsx";
import {File} from "../page/File.tsx";

const router = createBrowserRouter([
  {
    path: "/login",
    element: <Login/>,
  },
  {
    path: '/',
    element: <Layout/>,
    children: [
      {
        path: '/file',
        element: <File/>
      }
    ]
  }
])

export {router}