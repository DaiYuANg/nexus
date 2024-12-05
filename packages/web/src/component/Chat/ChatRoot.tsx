import { ChatProvider } from './ChatContext.tsx';
import { Box, ScrollArea, Text } from '@mantine/core';
import { ChatInput } from './ChatInput.tsx';

const ChatRoot = () => {
  return (
    <ChatProvider>
      <Box style={{ height: '80vh', display: 'flex', flexDirection: 'column' }}>
        {/* Chat Header */}
        <Box style={{ textAlign: 'center', marginBottom: '1rem' }}>
          <Text size="xl" fw={600}>
            Current Chat
          </Text>
        </Box>

        {/* Chat messages */}
        <ScrollArea style={{ flex: 1, marginBottom: '1rem' }}>
          <Box>
            {/*{messages.map((msg, index) => (*/}
            {/*  <ChatBubble key={index} message={msg} isOwnMessage={index % 2 === 0} />*/}
            {/*))}*/}
          </Box>
        </ScrollArea>
        <ChatInput />
      </Box>
    </ChatProvider>
  );
};

export { ChatRoot };
