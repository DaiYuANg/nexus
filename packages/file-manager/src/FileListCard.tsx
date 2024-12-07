import { Card, Group, Text } from '@mantine/core';
import { FileInfo } from './type.ts';
import { Draggable } from '@hello-pangea/dnd';
import cx from 'clsx';
import classes from './DndList.module.css';
import { IconArrowAutofitContent } from '@tabler/icons-react';

type FileListCardProps = {
  file: FileInfo;
  index: number;
};

const FileListCard = ({ file, index }: FileListCardProps) => {
  return (
    <Draggable key={file.id} index={index} draggableId={file.id}>
      {(provided, snapshot) => (
        <div
          className={cx({ [classes.itemDragging]: snapshot.isDragging })}
          {...provided.draggableProps}
          {...provided.dragHandleProps}
          ref={provided.innerRef}
        >
          <Card
            onClick={(e) => {
              e.stopPropagation();
              console.log('click');
            }}
            m={'xs'}
            key={file.id}
            shadow="sm"
            radius="md"
          >
            <Group style={{ width: '100%' }} justify={'space-between'}>
              <Group>
                {file.icon}
                <Text fw={500} size="sm">
                  {file.name}
                </Text>
                <Text size="xs">type</Text>
                <Text size="xs">size</Text>
              </Group>
              <Group>
                <IconArrowAutofitContent />
              </Group>
            </Group>
          </Card>
        </div>
      )}
    </Draggable>
  );
};

export { FileListCard };
