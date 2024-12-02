import {Avatar, Group, Menu, rem, Text, UnstyledButton} from "@mantine/core"
import {
  IconArrowsLeftRight,
  IconChevronRight,
  IconMessageCircle,
  IconPhoto,
  IconSearch,
  IconSettings, IconTrash
} from "@tabler/icons-react";

const UserCard = () => {
  return <>
    <Menu shadow="md" width={200}>
      <Menu.Target>
        <UnstyledButton p={'md'}>
          <Group>
            <Avatar
              src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-8.png"
              radius="xl"
            />

            <div style={{flex: 1}}>
              <Text size="sm" fw={500}>
                Harriette Spoonlicker
              </Text>

              <Text c="dimmed" size="xs">
                hspoonlicker@outlook.com
              </Text>
            </div>

            <IconChevronRight size={14} stroke={1.5}/>
          </Group>
        </UnstyledButton>
      </Menu.Target>

      <Menu.Dropdown>
        <Menu.Label>Application</Menu.Label>
        <Menu.Item leftSection={<IconSettings style={{ width: rem(14), height: rem(14) }} />}>
          Settings
        </Menu.Item>
        <Menu.Item leftSection={<IconMessageCircle style={{ width: rem(14), height: rem(14) }} />}>
          Messages
        </Menu.Item>
        <Menu.Item leftSection={<IconPhoto style={{ width: rem(14), height: rem(14) }} />}>
          Gallery
        </Menu.Item>
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
        <Menu.Item
          leftSection={<IconArrowsLeftRight style={{ width: rem(14), height: rem(14) }} />}
        >
          Transfer my data
        </Menu.Item>
        <Menu.Item
          color="red"
          leftSection={<IconTrash style={{ width: rem(14), height: rem(14) }} />}
        >
          Delete my account
        </Menu.Item>
      </Menu.Dropdown>
    </Menu>

  </>
}

export {UserCard}