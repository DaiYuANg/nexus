import { LazyComponentWrapper } from './Lazy.tsx';
import { FileManagerProps } from './FileManager/FileManagerRoot.tsx';

const FileManager = (props: FileManagerProps) => {
  return <LazyComponentWrapper importFunc={() => import('./FileManager/FileManagerRoot.tsx')} props={props} />;
};

const Layout = () => {
  return <LazyComponentWrapper importFunc={() => import('./Layout/Layout.tsx')} props={undefined} />;
};
export { FileManager, Layout };
