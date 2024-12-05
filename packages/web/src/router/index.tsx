import { createBrowserRouter } from 'react-router';
import { Layout } from '../component';
import { CalendarScreen, Chat, File, Flow, Login } from '../page';
import { NothingFound } from '../page/error/NothingFound.tsx';

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
  {
    path: '/admin',
  },
  { path: '/*', element: <NothingFound /> },
]);

export { router };
