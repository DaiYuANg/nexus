import { ActionIcon, Anchor, Breadcrumbs, Flex, Group, TextInput, useMantineTheme } from '@mantine/core';
import { IconHome, IconLayoutGrid, IconList, IconSearch } from '@tabler/icons-react';
import { useCallback } from 'react';
import { useFileManagerContext } from './FileManagerContext.tsx';
import { NewFile } from './NewFile.tsx';

const FileManagerToolBar = () => {
  const { layout, setLayout } = useFileManagerContext();

  const theme = useMantineTheme();

  const switchLayout = useCallback(() => {
    // setIsListLayout(!isListLayout)
    if (layout === 'grid') {
      setLayout('list');
      return;
    }

    if (layout === 'list') {
      setLayout('grid');
      return;
    }
  }, [layout]);

  return (
    <Flex justify={'space-between'}>
      <Group align={'center'}>
        <Breadcrumbs separator="/">
          <Anchor href="/">
            <IconHome />
          </Anchor>
          <Anchor href="/category">Category</Anchor>
          <Anchor href="/category/item">Item</Anchor>
        </Breadcrumbs>
      </Group>
      <Group>
        <TextInput
          radius="xl"
          size="md"
          placeholder="Search Text"
          rightSectionWidth={42}
          leftSection={<IconSearch size={18} stroke={1.5} />}
          rightSection={
            <ActionIcon variant={'light'} size={32} radius="xl" color={theme.primaryColor}>
              <IconSearch size={18} stroke={1.5} />
            </ActionIcon>
          }
        />
        <NewFile />
        <ActionIcon onClick={switchLayout} variant={'light'}>
          {layout === 'grid' ? <IconLayoutGrid /> : <IconList />}
        </ActionIcon>
      </Group>
    </Flex>
  );
};

export { FileManagerToolBar };
