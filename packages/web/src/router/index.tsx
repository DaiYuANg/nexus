import { createBrowserRouter } from 'react-router';
import { Login } from '../page/auth/Login.tsx';
import { Layout } from '../component/Layout.tsx';
import { File } from '../page/file/File.tsx';
import { Flow } from '../page/flow/Flow.tsx';
import { CalendarScreen } from '../page/calendar/Calendar.tsx';
import { Chat } from '../page/chat/Chat.tsx';

const router = createBrowserRouter([
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        path: '/file',
        element: <File />,
      },
      {
        path: '/flow',
        element: <Flow />,
      },
      {
        path: '/calendar',
        element: <CalendarScreen />,
      },
      {
        path: '/chat',
        element: <Chat />,
      },
    ],
  },
]);

export { router };
