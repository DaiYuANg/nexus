import { Avatar, Group, Menu, rem, Text, UnstyledButton } from '@mantine/core';
import {
  IconArrowsLeftRight,
  IconChevronRight,
  IconLogout,
  IconMessageCircle,
  IconPhoto,
  IconSearch,
  IconSettings2,
} from '@tabler/icons-react';
import { useUserStore } from '../store/useUserStore.tsx';
import { useNavigate } from 'react-router';

const UserCard = () => {
  const { user, clearUser } = useUserStore();
  const navigate = useNavigate();
  return (
    <>
      <Menu shadow="md" width={200}>
        <Menu.Target>
          <UnstyledButton p={'md'}>
            <Group>
              <Avatar src={user?.avatar} radius="xl" />

              <div style={{ flex: 1 }}>
                <Text size="sm" fw={500}>
                  Harriette Spoonlicker
                </Text>

                <Text c="dimmed" size="xs">
                  {user?.email}
                </Text>
              </div>

              <IconChevronRight size={14} stroke={1.5} />
            </Group>
          </UnstyledButton>
        </Menu.Target>

        <Menu.Dropdown>
          <Menu.Label>{user?.email}</Menu.Label>
          <Menu.Item
            leftSection={<IconSettings2 style={{ width: rem(14), height: rem(14) }} />}
            onClick={() => navigate('/settings')}
          >
            Settings
          </Menu.Item>
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
          <Menu.Item
            onClick={() => {
              clearUser();
            }}
            color="red"
            leftSection={<IconLogout style={{ width: rem(14), height: rem(14) }} />}
          >
            Logout
          </Menu.Item>
        </Menu.Dropdown>
      </Menu>
    </>
  );
};

export { UserCard };
