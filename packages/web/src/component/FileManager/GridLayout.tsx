import { Grid } from '@mantine/core';
import { FileGridCard } from './FileGridCard.tsx';
import { useFileManagerContext } from './FileManagerContext.tsx';

const GridLayout = () => {
  const { files } = useFileManagerContext();

  return (
    <Grid gutter="md" p={20}>
      {files.map((folder, index) => (
        <Grid.Col span={2} key={index}>
          <FileGridCard {...folder} />
        </Grid.Col>
      ))}
    </Grid>
  );
};

export { GridLayout };
