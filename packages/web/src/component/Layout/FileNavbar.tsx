import { Box, NavLink } from '@mantine/core';
import { IconChevronRight, IconClock24, IconFile, IconShare3, IconStars, IconTags } from '@tabler/icons-react';
import { useLocation, useNavigate } from 'react-router';
import { ReactElement } from 'react';

type FileNav = {
  label: string;
  left: ReactElement;
  right: ReactElement;
  path: string;
};

const fileNavs: Array<FileNav> = [
  {
    label: 'ALL files',
    left: <IconFile size="1rem" stroke={1.5} />,
    right: <IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />,
    path: '/file',
  },
  {
    label: 'Recent',
    left: <IconClock24 size="1rem" stroke={1.5} />,
    right: <IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />,
    path: '/file/recent',
  },
  {
    label: 'Share',
    left: <IconShare3 size="1rem" stroke={1.5} />,
    right: <IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />,
    path: '/file/share',
  },
  {
    label: 'Favorite',
    left: <IconStars size="1rem" stroke={1.5} />,
    right: <IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />,
    path: '/file/favorite',
  },
  {
    label: 'Tag',
    left: <IconTags size="1rem" stroke={1.5} />,
    right: <IconChevronRight size="0.8rem" stroke={1.5} className="mantine-rotate-rtl" />,
    path: '/file/tag',
  },
];

const FileNavbar = () => {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  return (
    <Box>
      {fileNavs.map((fileNav, index) => (
        <NavLink
          key={index}
          label={fileNav.label}
          leftSection={fileNav.left}
          rightSection={fileNav.right}
          active={pathname === fileNav.path}
          onClick={() => navigate(fileNav.path)}
        />
      ))}
    </Box>
  );
};

export { FileNavbar };
