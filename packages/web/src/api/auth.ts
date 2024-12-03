import { Result } from '../type/result.ts';
import { req } from './req.ts';

type loginForm = {
  email: string;
  password: string;
};

type loginResponse = {
  token: string;
  email: string;
  avatar: string;
};

const login = (form: loginForm): Promise<Result<loginResponse>> => {
  return req.post('/user/login', form);
};

export { login };

export type { loginResponse, loginForm };
