import axios from "axios";
import {nprogress} from "@mantine/nprogress";

const req = axios.create({
  baseURL: '/api'
})
// 请求拦截器
axios.interceptors.request.use(
  (config) => {
    nprogress.set(50)
    return config;
  },
  (error) => {
    nprogress.reset();    // 重置进度条
    return Promise.reject(error);
  }
);

req.interceptors.response.use((resp) => {
  nprogress.set(100)
  return resp.data;
})

export {req}