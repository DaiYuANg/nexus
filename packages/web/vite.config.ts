import { defineConfig, PluginOption } from 'vite';
import react from '@vitejs/plugin-react';
import TurboConsole from 'unplugin-turbo-console/vite';
import { visualizer } from 'rollup-plugin-visualizer';
import { compression } from 'vite-plugin-compression2';
import tsconfigPaths from 'vite-tsconfig-paths';
import * as os from 'node:os';
import { VitePWA } from 'vite-plugin-pwa';
const mode = process.env.NODE_ENV;
// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    TurboConsole(),
    visualizer({
      emitFile: false,
    }) as PluginOption,
    compression(),
    tsconfigPaths(),
    VitePWA({ registerType: 'autoUpdate' }),
  ],
  esbuild: {
    drop: mode === 'production' ? ['console', 'debugger'] : [],
  },
  build: {
    reportCompressedSize: true,
    rollupOptions: {
      output: {
        experimentalMinChunkSize: 1024,
        manualChunks: {
          mantine: [
            '@mantine/core',
            '@mantine/dates',
            '@mantine/form',
            '@mantine/hooks',
            '@mantine/modals',
            '@mantine/notifications',
            '@mantine/nprogress',
            '@mantine/tiptap',
          ],
          'code-highlight': ['@mantine/code-highlight'],
          'mantine-extension': ['@mantine/charts', '@mantine/spotlight'],
        },
      },
    },
  },
  server: {
    warmup: {
      clientFiles: ['./src/components/*/*.tsx', './src/page/*/**.tsx'],
    },
    open: true,
    ...(os.platform() === 'darwin'
      ? {
          proxy: {
            '/api': {
              target: 'http://localhost:8080',
              changeOrigin: true,
              configure: (proxy, options) => {
                proxy.on('proxyReq', (proxyReq, _req, _res) => {
                  proxyReq.setHeader('origin', 'http://localhost:8080');
                  console.log('Sending Request to the Target:', options.target + proxyReq.path);
                });
              },
            },
          },
        }
      : {}),
  },
});
