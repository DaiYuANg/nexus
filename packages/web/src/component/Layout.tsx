import { AppShell, NavLink } from '@mantine/core';
import { Outlet } from 'react-router';
import { IconChevronRight, IconClock24, IconFile, IconShare3, IconStar, IconTag } from '@tabler/icons-react';
import { BackgroundUpload } from './BackgroundUpload.tsx';
import { GlobalSpotlight } from './GlobalSpotlight.tsx';
import { Header } from './Header.tsx';

const Layout = () => {
  return (
    <>
      <AppShell
        header={{ height: 60 }}
        navbar={{ width: 300, breakpoint: 'sm', collapsed: { desktop: false } }}
        padding="md"
      >
        <AppShell.Header>
          <Header />
        </AppShell.Header>
        <AppShell.Navbar p="md">
          <NavLink
            label={'All files'}
            leftSection={<IconFile size="1rem" stroke={1.5} />}
            rightSection={<IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />}
            active
          />
          <NavLink
            label={'Recent'}
            leftSection={<IconClock24 size="1rem" stroke={1.5} />}
            rightSection={<IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />}
          />
          <NavLink
            label={'Share'}
            leftSection={<IconShare3 size="1rem" stroke={1.5} />}
            rightSection={<IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />}
          />
          <NavLink
            label={'Favorite'}
            leftSection={<IconStar size="1rem" stroke={1.5} />}
            rightSection={<IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />}
          />
          <NavLink
            label={'Tag'}
            leftSection={<IconTag size="1rem" stroke={1.5} />}
            rightSection={<IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />}
          />
        </AppShell.Navbar>
        <AppShell.Main>
          <Outlet />
        </AppShell.Main>
        <AppShell.Aside hidden={true}></AppShell.Aside>
      </AppShell>
      <BackgroundUpload />
      <GlobalSpotlight />
    </>
  );
};

export { Layout };
