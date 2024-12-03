import { ActionIcon, Box, Group, ScrollArea, Text, Textarea } from '@mantine/core';
import { KeyboardEvent, useState } from 'react';
import { ChatBubble } from './Bubble.tsx';
import { IconSend } from '@tabler/icons-react';

const Chat = () => {
  const [messages, setMessages] = useState<string[]>([]);
  const [messageInput, setMessageInput] = useState('');

  const handleSendMessage = () => {
    if (messageInput.trim() === '') return; // Do nothing if input is empty
    setMessages([...messages, messageInput]);
    setMessageInput('');
  };

  const handleKeyDown = (event: KeyboardEvent) => {
    // Check if Enter key was pressed (without shift, so it creates a new line in textarea)
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault(); // Prevent the default behavior of Enter key (adding a new line)
      handleSendMessage(); // Send the message
    }
  };

  return (
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
          {messages.map((msg, index) => (
            <ChatBubble key={index} message={msg} isOwnMessage={index % 2 === 0} />
          ))}
        </Box>
      </ScrollArea>

      {/* Message input */}
      <Group style={{ width: '100%' }}>
        <Textarea
          variant="filled"
          placeholder="Type a message"
          value={messageInput}
          onChange={(event) => setMessageInput(event.currentTarget.value)}
          autosize
          autoFocus
          onKeyDown={handleKeyDown}
          minRows={2}
          style={{ flexGrow: 1 }}
        />
        <ActionIcon color="blue" size="lg" onClick={handleSendMessage}>
          <IconSend size={20} />
        </ActionIcon>
      </Group>
    </Box>
  );
};

export { Chat };
