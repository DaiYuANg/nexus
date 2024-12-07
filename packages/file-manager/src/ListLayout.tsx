import { Box } from '@mantine/core';
import { useFileManagerContext } from './FileManagerContext.tsx';
import { FileListCard } from './FileListCard.tsx';
import { DragDropContext, Droppable } from '@hello-pangea/dnd';
import { useListState } from '@mantine/hooks';

const ListLayout = () => {
  const { files } = useFileManagerContext();
  const [_states, handlers] = useListState(files);

  const items = files.map((file, index) => <FileListCard key={file.id} index={index} file={file} />);
  return (
    <Box>
      <DragDropContext
        onDragEnd={({ destination, source }) => {
          return handlers.reorder({ from: source.index, to: destination?.index || 0 });
        }}
      >
        <Droppable droppableId="dnd-list" direction="vertical">
          {(provided) => (
            <div {...provided.droppableProps} ref={provided.innerRef}>
              {items}
              {provided.placeholder}
            </div>
          )}
        </Droppable>
      </DragDropContext>
    </Box>
  );
};

export { ListLayout };
