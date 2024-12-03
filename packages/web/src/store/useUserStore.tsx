import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';
import { loginResponse } from '../api/auth.ts';

interface UserState {
  user: loginResponse | null; // 存储整个登录响应对象
  setUser: (user: loginResponse) => void; // 更新 user 的函数
  clearUser: () => void; // 清除用户信息的函数
}

const useUserStore = create<UserState>()(
  devtools(
    persist(
      (set) => ({
        user: null, // 初始化 user 为 null
        setUser: (newUser: loginResponse) => set({ user: newUser }), // 设置 user
        clearUser: () => set({ user: null }), // 清除 user
      }),
      {
        name: 'user-storage', // 本地存储的名称
      },
    ),
    { name: 'UserStore' }, // Zustand devtools 的名称
  ),
);
export { useUserStore };
