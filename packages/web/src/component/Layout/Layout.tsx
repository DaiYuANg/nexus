import { AppShell } from '@mantine/core';
import { Outlet } from 'react-router';
import { BackgroundUpload } from '../BackgroundUpload.tsx';
import { GlobalSpotlight } from '../GlobalSpotlight.tsx';
import { Header } from './Header.tsx';
import { Navbar } from './Navbar.tsx';
import { LayoutProvider } from './LayoutContext.tsx';

const headerHeight = 60;
const navbarWidth = 300;
const Layout = () => {
  return (
    <LayoutProvider headerHeight={headerHeight} navbarWidth={navbarWidth}>
      <AppShell
        header={{ height: headerHeight }}
        navbar={{ width: navbarWidth, breakpoint: 'sm', collapsed: { desktop: false } }}
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
    </LayoutProvider>
  );
};

export default Layout;
export { Layout };
