import { createTheme } from '@mantine/core';
import '@fontsource/jetbrains-mono/300.css';
import '@fontsource/jetbrains-mono/400.css';
import '@fontsource/jetbrains-mono/400-italic.css';

export const theme = createTheme({
  fontFamily: 'Jetbrains mono',
  focusRing: 'auto',
  fontSmoothing: true,
  autoContrast: true,
  cursorType: 'pointer',
});
