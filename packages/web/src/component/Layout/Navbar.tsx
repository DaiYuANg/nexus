import { Box } from '@mantine/core';
import { useLocation } from 'react-router';
import { navbarContents } from './NavbarContent.tsx';

const NavbarContent = () => {
  const location = useLocation();
  console.log(location.pathname);
  return navbarContents.find((item) => {
    return item.path === location.pathname;
  })?.element;
};

const Navbar = () => {
  return (
    <Box>
      <NavbarContent />
    </Box>
  );
};

export { Navbar };
