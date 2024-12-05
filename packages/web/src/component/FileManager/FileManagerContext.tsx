import { createContext, ReactNode, useCallback, useContext, useState } from 'react';
import { FileInfo, layout, RenderProps } from './type.ts';

type FileManagerContextProp = {
  layout: layout;
  setLayout: (layout: layout) => void;
  files: FileInfo[];
  refreshKey: number;
  incrementRefreshKey: () => void;
};

const FileManagerContext = createContext<FileManagerContextProp | undefined>(undefined);

// 定义 provider 的 props
type FileManagerProviderProps = {
  children: ReactNode;
} & RenderProps;

const FileManagerProvider = ({ children, files }: FileManagerProviderProps) => {
  const [layout, setLayout] = useState<layout>('grid');

  // 定义更新 currentCategory 的回调函数
  const setCurrent = useCallback(
    (layout: layout) => {
      setLayout(layout);
    },
    [setLayout],
  );

  const [refreshKey, setRefreshKey] = useState(0);
  const handleAddSuccess = useCallback(() => {
    setRefreshKey((prevKey) => prevKey + 1);
  }, []);

  return (
    <FileManagerContext.Provider
      value={{
        refreshKey,
        incrementRefreshKey: handleAddSuccess,
        layout,
        setLayout: setCurrent,
        files,
      }}
    >
      {children}
    </FileManagerContext.Provider>
  );
};

// 自定义 Hook，方便在组件中使用
const useFileManagerContext = () => {
  const context = useContext(FileManagerContext);
  if (!context) {
    throw new Error('useDriverBoardContext must be used within a ProductProvider');
  }
  return context;
};

export { FileManagerProvider, useFileManagerContext };
