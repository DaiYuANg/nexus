import { Card, Center, Checkbox, HoverCard, Text } from '@mantine/core';
import { useContextMenu } from 'mantine-contextmenu';
import { IconCopy, IconDownload } from '@tabler/icons-react';
import { useState } from 'react';
import { FileInfo } from './type.ts';

type FileCardProps = Partial<
  {
    files: number;
    showCheckbox?: boolean;
  } & FileInfo
>;

const FileGridCard = ({ name, icon, files, showCheckbox }: FileCardProps) => {
  const { showContextMenu } = useContextMenu();
  const [checked, setChecked] = useState(false);
  return (
    <HoverCard openDelay={1000}>
      <HoverCard.Target>
        <Card
          onContextMenu={showContextMenu(
            [
              {
                key: 'copy',
                icon: <IconCopy />,
                onClick: () => {
                  console.log(123);
                },
              },
              {
                key: 'download',
                icon: <IconDownload />,
                onClick: () => {
                  console.log(312);
                },
              },
            ],
            { zIndex: 3000 },
          )}
          shadow="md"
          padding="sm"
          radius="sm"
          style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            cursor: 'pointer',
            position: 'relative',
          }}
          c={'blue'}
        >
          {showCheckbox && (
            <Checkbox
              checked={checked}
              onChange={(e) => setChecked(e.target.checked)}
              style={{
                position: 'absolute', // 设置绝对定位
                top: 10, // 距离顶部 10px
                left: 10, // 距离左侧 10px
              }}
            />
          )}
          <Center style={{ marginBottom: '10px' }}>{icon}</Center>
          <Text size="lg" style={{ marginBottom: '8px' }}>
            {name}
          </Text>
          <Text size="sm">{files} files</Text>
        </Card>
      </HoverCard.Target>
      <HoverCard.Dropdown>
        <Text size="sm">File Size:{files}</Text>
      </HoverCard.Dropdown>
    </HoverCard>
  );
};
export { FileGridCard };
