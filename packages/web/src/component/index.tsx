import { LazyComponentWrapper } from './Lazy.tsx';

const Layout = () => {
  return <LazyComponentWrapper importFunc={() => import('./Layout/Layout.tsx')} props={undefined} />;
};
export { Layout };
