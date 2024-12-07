import { ReactElement } from 'react';

type layout = 'grid' | 'list';

type RenderProps = {
  files: FileInfo[];
};

type FileInfo = {
  id: string;
  name: string;
} & Partial<{
  icon: ReactElement;
  isDir: boolean;
  updateAt: Date;
  size: number;
  extra: object;
}>;

export type { layout, RenderProps, FileInfo };
