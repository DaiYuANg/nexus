import axios from 'axios';
import { nprogress } from '@mantine/nprogress';
import { useAuthStore } from '../store/useAuthStore.ts';

const req = axios.create({
  baseURL: '/api',
});
const state = useAuthStore.getState();
// 请求拦截器
req.interceptors.request.use(
  (config) => {
    config.headers['Authorization'] = 'Bearer ' + state.token;
    nprogress.set(30);
    return config;
  },
  (error) => {
    nprogress.reset(); // 重置进度条
    return Promise.reject(error);
  },
);

req.interceptors.response.use((resp) => {
  nprogress.complete();
  nprogress.cleanup();
  return resp.data;
});

export { req };
