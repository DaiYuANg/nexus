import { ActionIcon, Group, Textarea } from '@mantine/core';
import { IconSend } from '@tabler/icons-react';
import { KeyboardEvent, useCallback, useState } from 'react';
import { useChatContext } from './ChatContext.tsx';

const ChatInput = () => {
  const { sendMessage } = useChatContext();

  const [messageInput, setMessageInput] = useState('');
  const handleKeyDown = useCallback((event: KeyboardEvent) => {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault();
      handleSend();
    }
  }, []);

  const handleSend = useCallback(() => {
    sendMessage(messageInput);
  }, []);

  return (
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
      <ActionIcon color="blue" size="lg" onClick={handleSend}>
        <IconSend size={20} />
      </ActionIcon>
    </Group>
  );
};

export { ChatInput };
