import { req } from './req.ts';

const fileList = () => {
  return req.get('/file/list');
};

export { fileList };
