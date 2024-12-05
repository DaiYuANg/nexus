import { Box, NavLink } from '@mantine/core';
import { IconChevronRight, IconClock24, IconFile, IconShare3, IconStar, IconTag } from '@tabler/icons-react';

const Navbar = () => {
  return (
    <Box>
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
    </Box>
  );
};

export { Navbar };
