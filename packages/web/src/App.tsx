import { RouterProvider } from 'react-router';
import { router } from './router';
import './i18n.ts';

const App = () => {
  return <RouterProvider router={router} />;
};

export { App };
