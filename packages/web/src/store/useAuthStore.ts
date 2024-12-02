import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

interface AuthState {
  token: string | null;
  setToken: (newToken: string) => void;
  clearToken: () => void;
}

const useAuthStore = create<AuthState>()(
  devtools(
    persist(
      (set) => ({
        token: null,
        setToken: (newToken: string) => set({ token: newToken }),
        clearToken: () => set({ token: null }),
      }),
      {
        name: "auth-storage",
      },
    ),
    { name: "AuthStore" },
  ),
);

export { useAuthStore };
export type { AuthState };