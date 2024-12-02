import type {Config} from 'tailwindcss';

export default {
  darkMode: 'media',
  content: ['src/**/*.html', 'src/**/*.tsx', 'src/**/*.mdx'],
  theme: {
    extend: {},
  },
  plugins: [],
} satisfies Config;