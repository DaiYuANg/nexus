import { scan } from 'react-scan';
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './App.css';
import { App } from './App.tsx';
import { MantineProvider } from '@mantine/core';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Notifications } from '@mantine/notifications';
import { NavigationProgress } from '@mantine/nprogress';
import { theme } from './theme.ts';
import { colorSchemeManager } from './colorSchemeManager.ts';
import { ContextMenuProvider } from 'mantine-contextmenu';

if (import.meta.env.MODE === 'development') {
  scan({
    enabled: true,
    log: true,
  });
}
const queryClient = new QueryClient();
createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <MantineProvider defaultColorScheme={'auto'} theme={theme} colorSchemeManager={colorSchemeManager}>
        <Notifications position={'top-right'} />
        <NavigationProgress />
        <ReactQueryDevtools initialIsOpen={false} />
        <ContextMenuProvider>
          <App />
        </ContextMenuProvider>
      </MantineProvider>
    </QueryClientProvider>
  </StrictMode>,
);
