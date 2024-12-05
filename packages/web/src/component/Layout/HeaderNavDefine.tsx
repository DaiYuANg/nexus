import {
  IconCalendar,
  IconFolder,
  IconLayoutDashboard,
  IconMail,
  IconMessage2,
  IconPhoto,
  IconShape2,
  IconTimeline,
} from '@tabler/icons-react';
import { ReactElement } from 'react';

type HeaderNav = {
  path: string;
  icon: ReactElement;
};

const HeaderNavDefine: Array<HeaderNav> = [
  {
    path: '/dashboard',
    icon: <IconLayoutDashboard />,
  },
  {
    path: '/file',
    icon: <IconFolder />,
  },
  {
    path: '/chat',
    icon: <IconMessage2 />,
  },
  {
    path: '/flow',
    icon: <IconShape2 />,
  },
  {
    path: '/timeline',
    icon: <IconTimeline />,
  },
  {
    path: '/photo',
    icon: <IconPhoto />,
  },
  {
    path: '/calendar',
    icon: <IconCalendar />,
  },
  {
    path: '/email',
    icon: <IconMail />,
  },
];
export { HeaderNavDefine };
export type { HeaderNav };
