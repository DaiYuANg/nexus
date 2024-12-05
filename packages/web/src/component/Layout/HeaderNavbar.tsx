import { ActionIcon, ActionIconGroup } from '@mantine/core';
import { useNavigate } from 'react-router';
import type { HeaderNav } from './HeaderNavDefine.tsx';
import { HeaderNavDefine } from './HeaderNavDefine.tsx';

type HeaderNavProps = {} & HeaderNav;

const HeaderNav = ({ path, icon }: HeaderNavProps) => {
  const navigate = useNavigate();
  return (
    <ActionIcon
      onClick={() => {
        navigate(path);
      }}
      variant={'transparent'}
    >
      {icon}
    </ActionIcon>
  );
};

const HeaderNavbar = () => {
  return (
    <ActionIconGroup>
      {HeaderNavDefine.map((item, index) => (
        <HeaderNav key={index} {...item} />
      ))}
    </ActionIconGroup>
  );
};

export { HeaderNavbar };
