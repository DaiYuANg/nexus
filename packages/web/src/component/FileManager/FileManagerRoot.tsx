import { Box, Divider } from '@mantine/core';
import { FileManagerProvider } from './FileManagerContext.tsx';
import { FileManagerToolBar } from './FileManagerToolBar.tsx';
import { RenderProps } from './type.ts';
import { LayoutManager } from './LayoutManager.tsx';

type FileManagerProps = {} & RenderProps;

const FileManagerRoot = ({ files }: FileManagerProps) => {
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

export default FileManagerRoot;
export type { FileManagerProps };
