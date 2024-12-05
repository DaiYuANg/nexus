import { Box } from '@mantine/core';
import { IconFolder } from '@tabler/icons-react';
import { FileManager } from '../../component';

const folders = [
  { id: '1', name: 'Documents', icon: <IconFolder size={30} />, files: 12 },
  { id: '2', name: 'Photos', icon: <IconFolder size={30} />, files: 43 },
  { id: '3', name: 'Music', icon: <IconFolder size={30} />, files: 10 },
  { id: '4', name: 'Videos', icon: <IconFolder size={30} />, files: 28 },
  { id: '5', name: 'Downloads', icon: <IconFolder size={30} />, files: 5 },
  { id: '6', name: 'Projects', icon: <IconFolder size={30} />, files: 22 },
];

const File = () => {
  return (
    <Box>
      <FileManager files={folders} />
    </Box>
  );
};

export default File;
export { File };
