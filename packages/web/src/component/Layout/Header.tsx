import { ActionIcon, ActionIconGroup, Burger, Flex, Group, Text } from '@mantine/core';
import { IconAdjustmentsSearch, IconCloudSearch, IconLanguage, IconNotification } from '@tabler/icons-react';
import { openSpotlight } from '@mantine/spotlight';
import { ThemeSwitcher } from '../ThemeSwitcher.tsx';
import { UserCard } from '../UserCard.tsx';
import { useDisclosure } from '@mantine/hooks';
import { HeaderNavbar } from './HeaderNavbar.tsx';

const Header = () => {
  const [opened, { toggle }] = useDisclosure();
  return (
    <>
      <Flex justify={'space-between'} mih={70} pr={'md'} align={'center'}>
        <Group h="100%" px="md">
          <Burger opened={opened} onClick={toggle} size="sm" />
          <Text size="lg" fw={500}>
            Logo Placeholder
          </Text>
          <HeaderNavbar />
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
