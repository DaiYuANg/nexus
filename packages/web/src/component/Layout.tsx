import {ActionIcon, ActionIconGroup, AppShell, Burger, Flex, Group, NavLink, Text} from "@mantine/core";
import {useDisclosure} from "@mantine/hooks";
import {Outlet} from "react-router";
import {
  IconAdjustmentsSearch,
  IconChevronRight,
  IconCloudSearch,
  IconFile,
  IconLanguage,
  IconLayoutDashboard,
  IconMessage2,
  IconNotification
} from "@tabler/icons-react";
import {openSpotlight} from "@mantine/spotlight";
import {UserCard} from "./UserCard.tsx";
import {BackgroundUpload} from "./BackgroundUpload.tsx";
import {GlobalSpotlight} from "./GlobalSpotlight.tsx";
import {ThemeSwitcher} from "./ThemeSwitcher.tsx";

const Layout = () => {
  const [opened, {toggle}] = useDisclosure();
  return <>
    <AppShell
      header={{height: 60}}
      navbar={{width: 300, breakpoint: 'sm', collapsed: {desktop: false}}}
      padding="md"
    >
      <AppShell.Header>
        <Flex justify={'space-between'} mih={70} pr={'md'} align={'center'}>
          <Group h="100%" px="md">
            <Burger opened={opened} onClick={toggle} size="sm"/>
            <Text size="lg" fw={500}>Logo Placeholder</Text>
            <ActionIconGroup>
              <ActionIcon variant={'light'}>
                <IconLayoutDashboard/>
              </ActionIcon>
              <ActionIcon variant={'light'}>
                <IconMessage2/>
              </ActionIcon>
            </ActionIconGroup>
          </Group>
          <Group>
            <ActionIconGroup>
              <ActionIcon onClick={openSpotlight} variant={'light'}>
                <IconAdjustmentsSearch/>
              </ActionIcon>
              <ActionIcon variant={'light'}>
                <IconLanguage/>
              </ActionIcon>
              <ActionIcon variant={'light'}>
                <IconNotification/>
              </ActionIcon>
              <ActionIcon variant={'light'}>
                <IconCloudSearch/>
              </ActionIcon>
            </ActionIconGroup>
            <ThemeSwitcher/>
            <UserCard/>
          </Group>
        </Flex>
      </AppShell.Header>
      <AppShell.Navbar p="md">
        <NavLink
          label={'All files'}
          leftSection={<IconFile size="1rem" stroke={1.5}/>}
          rightSection={
            <IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl"/>
          }
          active
        />
      </AppShell.Navbar>
      <AppShell.Main>
        <Outlet/>
      </AppShell.Main>
      <AppShell.Aside hidden={true}>
      </AppShell.Aside>
    </AppShell>
    <BackgroundUpload/>
    <GlobalSpotlight/>
  </>
}

export {Layout}