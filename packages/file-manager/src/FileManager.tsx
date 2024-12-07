import { Box, Divider } from '@mantine/core';
import { FileManagerProvider } from './FileManagerContext.tsx';
import { FileManagerToolBar } from './FileManagerToolBar.tsx';
import { RenderProps } from './type.ts';
import { LayoutManager } from './LayoutManager.tsx';
import { ReactElement } from 'react';

type FileManagerProps = {} & RenderProps;

const FileManager = ({ files }: FileManagerProps): ReactElement => {
  return (
    <FileManagerProvider files={files}>
      <Box>
        <FileManagerToolBar />
        <Divider mt={'sm'} />
        <LayoutManager />
      </Box>
    </FileManagerProvider>
  );
};

export default FileManager;
export type { FileManagerProps };
