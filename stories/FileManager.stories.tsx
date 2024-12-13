import type { Meta, StoryObj } from '@storybook/react';
import { FileManager } from '@nexus/file-manager';
import { ButtonProps } from './Button.tsx';

const meta = {
  title: 'Example/Control',
  component: FileManager,
  parameters: {
    // Optional parameter to center the component in the Canvas. More info: https://storybook.js.org/docs/configure/story-layout
    layout: 'centered',
  },
  // This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
  tags: ['autodocs'],
  // More on argTypes: https://storybook.js.org/docs/api/argtypes
  argTypes: {
    // backgroundColor: { control: 'color' },
  },
  // Use `fn` to spy on the onClick arg, which will appear in the actions panel once invoked: https://storybook.js.org/docs/essentials/actions#action-args
  // args: { onClick: fn() },
} satisfies Meta<typeof FileManager>;

export default meta;
export type Story = StoryObj<typeof meta>;

export const FileManagerExample = ({
  primary = false,
  size = 'medium',
  backgroundColor,
  label,
  ...props
}: ButtonProps) => {
  // const mode = primary ? 'storybook-button--primary' : 'storybook-button--secondary';
  return (
    <FileManager
      // className={['storybook-button', `storybook-button--${size}`, mode].join(' ')}
      // style={{ backgroundColor }}
      {...props}
      files={[]}
    ></FileManager>
  );
};
