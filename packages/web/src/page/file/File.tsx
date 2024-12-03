import { ActionIcon, Anchor, Box, Breadcrumbs, Divider, Flex, Grid, Group } from '@mantine/core';
import { IconFolder, IconHome, IconLayoutGrid, IconList } from '@tabler/icons-react';
import { Folder } from '../../component/Folder.tsx';
import { useQuery } from '@tanstack/react-query';
import { fileList } from '../../api/file.ts';
import { useState } from 'react';

const folders = [
  { name: 'Documents', icon: <IconFolder size={30} />, files: 12 },
  { name: 'Photos', icon: <IconFolder size={30} />, files: 43 },
  { name: 'Music', icon: <IconFolder size={30} />, files: 10 },
  { name: 'Videos', icon: <IconFolder size={30} />, files: 28 },
  { name: 'Downloads', icon: <IconFolder size={30} />, files: 5 },
  { name: 'Projects', icon: <IconFolder size={30} />, files: 22 },
];

const File = () => {
  const { data } = useQuery({
    queryKey: ['filelist'],
    queryFn: fileList,
  });
  const [isListLayout, setIsListLayout] = useState(false);
  console.log(data);

  return (
    <>
      <Box>
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
            <ActionIcon onClick={() => setIsListLayout(!isListLayout)} variant={'light'}>
              {isListLayout ? <IconLayoutGrid /> : <IconList />}
            </ActionIcon>
          </Group>
        </Flex>
        <Divider mt={'sm'} />
        <Grid gutter="md" p={20}>
          {/* If isListLayout is true, render list layout; otherwise, render grid layout */}
          {isListLayout
            ? folders.map((folder, index) => (
                <Grid.Col span={12} key={index}>
                  {/* Here you can use a custom component or just text for folder */}
                  <div>{folder.name}</div>
                </Grid.Col>
              ))
            : folders.map((folder, index) => (
                <Grid.Col span={2} key={index}>
                  {/* Render Folder component, assuming you have it */}
                  <Folder {...folder} />
                </Grid.Col>
              ))}
        </Grid>
      </Box>
    </>
  );
};

export { File };
