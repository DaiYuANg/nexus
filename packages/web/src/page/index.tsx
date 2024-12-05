import { LazyComponentWrapper } from '../component/Lazy.tsx';

const Login = (props: any) => {
  return <LazyComponentWrapper importFunc={() => import('./auth/Login.tsx')} props={props} />;
};

const CalendarScreen = (props: any) => {
  return <LazyComponentWrapper importFunc={() => import('./calendar/Calendar.tsx')} props={props} />;
};

const Chat = (props: any) => {
  return <LazyComponentWrapper importFunc={() => import('./chat/Chat.tsx')} props={props} />;
};

const File = (props: any) => {
  return <LazyComponentWrapper importFunc={() => import('./file/File.tsx')} props={props} />;
};

const Flow = (props: any) => {
  return <LazyComponentWrapper importFunc={() => import('./flow/Flow.tsx')} props={props} />;
};

export { Login, CalendarScreen, Flow, File, Chat };
