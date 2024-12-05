import { AppShell } from '@mantine/core';
import { Outlet } from 'react-router';
import { BackgroundUpload } from '../BackgroundUpload.tsx';
import { GlobalSpotlight } from '../GlobalSpotlight.tsx';
import { Header } from './Header.tsx';
import { Navbar } from './Navbar.tsx';

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
          <Navbar />
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

export default Layout;
export { Layout };
