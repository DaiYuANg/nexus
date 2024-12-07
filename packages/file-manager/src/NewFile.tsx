import { ActionIcon, FileInput, Menu, rem } from '@mantine/core';
import { IconFileText, IconPlus, IconUpload } from '@tabler/icons-react';
import { useCallback, useRef } from 'react';

const NewFile = () => {
  const ref = useRef<HTMLButtonElement>(null);
  const handleUpload = useCallback(() => {
    ref.current?.click();
  }, [ref]);

  return (
    <Menu>
      <Menu.Target>
        <ActionIcon variant={'light'}>
          <IconPlus />
        </ActionIcon>
      </Menu.Target>
      <Menu.Dropdown>
        <Menu.Label>Create New File</Menu.Label>
        <Menu.Item onClick={handleUpload} leftSection={<IconUpload style={{ width: rem(14), height: rem(14) }} />}>
          Upload
        </Menu.Item>
        <Menu.Item leftSection={<IconFileText style={{ width: rem(14), height: rem(14) }} />}>New Text File</Menu.Item>
      </Menu.Dropdown>
      <FileInput ref={ref} display={'none'} />
    </Menu>
  );
};

export { NewFile };
