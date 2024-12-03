import { Card, Center, Menu, rem, Text } from '@mantine/core';
import { ReactNode } from 'react';
import {
  IconArrowsLeftRight,
  IconMessageCircle,
  IconPhoto,
  IconSearch,
  IconSettings,
  IconTrash,
} from '@tabler/icons-react';
import { useDisclosure } from '@mantine/hooks';

type FolderProps = {
  name: string;
  icon: ReactNode;
  files: number;
};

const Folder = ({ name, icon, files }: FolderProps) => {
  const [opened, { open }] = useDisclosure(false);
  return (
    <>
      <Menu opened={opened} shadow="md" width={200}>
        <Menu.Target>
          <Card
            onContextMenu={(e) => {
              e.preventDefault();
              open();
            }}
            color={'blue'}
            component={'a'}
            href="#"
            shadow="sm"
            padding="md"
            radius="md"
            style={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              cursor: 'pointer',
            }}
            c={'blue'}
          >
            <Center style={{ marginBottom: '10px' }}>{icon}</Center>
            <Text size="lg" style={{ marginBottom: '8px' }}>
              {name}
            </Text>
            <Text size="sm" color="dimmed">
              {files} files
            </Text>
          </Card>
        </Menu.Target>

        <Menu.Dropdown>
          <Menu.Label>Application</Menu.Label>
          <Menu.Item leftSection={<IconSettings style={{ width: rem(14), height: rem(14) }} />}>Settings</Menu.Item>
          <Menu.Item leftSection={<IconMessageCircle style={{ width: rem(14), height: rem(14) }} />}>
            Messages
          </Menu.Item>
          <Menu.Item leftSection={<IconPhoto style={{ width: rem(14), height: rem(14) }} />}>Gallery</Menu.Item>
          <Menu.Item
            leftSection={<IconSearch style={{ width: rem(14), height: rem(14) }} />}
            rightSection={
              <Text size="xs" c="dimmed">
                ⌘K
              </Text>
            }
          >
            Search
          </Menu.Item>

          <Menu.Divider />

          <Menu.Label>Danger zone</Menu.Label>
          <Menu.Item leftSection={<IconArrowsLeftRight style={{ width: rem(14), height: rem(14) }} />}>
            Transfer my data
          </Menu.Item>
          <Menu.Item color="red" leftSection={<IconTrash style={{ width: rem(14), height: rem(14) }} />}>
            Delete my account
          </Menu.Item>
        </Menu.Dropdown>
      </Menu>
    </>
  );
};
export { Folder };
