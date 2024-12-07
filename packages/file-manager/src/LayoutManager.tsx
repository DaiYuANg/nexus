import { useFileManagerContext } from './FileManagerContext.tsx';
import { Box } from '@mantine/core';
import { layout } from './type.ts';
import { GridLayout } from './GridLayout.tsx';
import { ListLayout } from './ListLayout.tsx';

type RenderLayoutProps = {
  layout: layout;
};

const RenderLayout = ({ layout }: RenderLayoutProps) => {
  switch (layout) {
    case 'grid':
      return <GridLayout />;
    case 'list':
      return <ListLayout />;
  }
};

const LayoutManager = () => {
  const { layout } = useFileManagerContext();

  return (
    <Box>
      <RenderLayout layout={layout} />
    </Box>
  );
};

export { LayoutManager };
