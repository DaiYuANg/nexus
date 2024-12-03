import { ActionIcon, ActionIconGroup, Burger, Flex, Group, Text } from '@mantine/core';
import {
  IconAdjustmentsSearch,
  IconCalendar,
  IconCloudSearch,
  IconFolder,
  IconLanguage,
  IconLayoutDashboard,
  IconMessage2,
  IconNotification,
  IconShape2,
} from '@tabler/icons-react';
import { openSpotlight } from '@mantine/spotlight';
import { ThemeSwitcher } from './ThemeSwitcher.tsx';
import { UserCard } from './UserCard.tsx';
import { useDisclosure } from '@mantine/hooks';
import { useNavigate } from 'react-router';

const Header = () => {
  const [opened, { toggle }] = useDisclosure();
  const navigate = useNavigate();
  return (
    <>
      <Flex justify={'space-between'} mih={70} pr={'md'} align={'center'}>
        <Group h="100%" px="md">
          <Burger opened={opened} onClick={toggle} size="sm" />
          <Text size="lg" fw={500}>
            Logo Placeholder
          </Text>
          <ActionIconGroup>
            <ActionIcon variant={'transparent'}>
              <IconLayoutDashboard />
            </ActionIcon>
            <ActionIcon
              onClick={() => {
                navigate('/file');
              }}
              variant={'transparent'}
            >
              <IconFolder />
            </ActionIcon>
            <ActionIcon
              onClick={() => {
                navigate('/chat');
              }}
              variant={'transparent'}
            >
              <IconMessage2 />
            </ActionIcon>
            <ActionIcon
              onClick={() => {
                navigate('/flow');
              }}
              variant={'transparent'}
            >
              <IconShape2 />
            </ActionIcon>
            <ActionIcon
              onClick={() => {
                navigate('/calendar');
              }}
              variant={'transparent'}
            >
              <IconCalendar />
            </ActionIcon>
          </ActionIconGroup>
        </Group>
        <Group>
          <ActionIconGroup>
            <ActionIcon onClick={openSpotlight} variant={'light'}>
              <IconAdjustmentsSearch />
            </ActionIcon>
            <ActionIcon variant={'light'}>
              <IconLanguage />
            </ActionIcon>
            <ActionIcon variant={'light'}>
              <IconNotification />
            </ActionIcon>
            <ActionIcon variant={'light'}>
              <IconCloudSearch />
            </ActionIcon>
          </ActionIconGroup>
          <ThemeSwitcher />
          <UserCard />
        </Group>
      </Flex>
    </>
  );
};

export { Header };
