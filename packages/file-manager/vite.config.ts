import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { resolve } from 'path';
import dts from 'vite-plugin-dts';
// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    dts({
      include: ['src/*'],
      rollupTypes: true,
      logLevel: 'info',
      insertTypesEntry: true,
    }),
  ],
  build: {
    copyPublicDir: false,
    lib: {
      entry: resolve(__dirname, 'src/index.ts'),
      name: 'file-manager',
    },
    rollupOptions: {
      external: ['react', 'react-dom', '@mantine/core', '@mantine/hook'],
    },
  },
});
