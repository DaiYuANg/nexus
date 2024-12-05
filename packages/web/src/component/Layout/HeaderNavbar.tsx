import { ActionIcon, ActionIconGroup } from '@mantine/core';
import {
  IconCalendar,
  IconFolder,
  IconLayoutDashboard,
  IconMessage2,
  IconPhoto,
  IconShape2,
  IconTimeline,
} from '@tabler/icons-react';
import { useNavigate } from 'react-router';

const HeaderNavbar = () => {
  const navigate = useNavigate();
  return (
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
          navigate('/flow');
        }}
        variant={'transparent'}
      >
        <IconTimeline />
      </ActionIcon>
      <ActionIcon
        onClick={() => {
          navigate('/flow');
        }}
        variant={'transparent'}
      >
        <IconPhoto />
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
  );
};

export { HeaderNavbar };
