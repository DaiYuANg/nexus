import { Avatar, Group, Paper, Text } from '@mantine/core';

const ChatBubble = ({ message, isOwnMessage }: { message: string; isOwnMessage: boolean }) => {
  return (
    <Group pos={'relative'}>
      <Avatar size="sm" src={isOwnMessage ? '/path-to-own-avatar.jpg' : '/path-to-other-avatar.jpg'} />
      <Paper p="xs" radius="md" style={{ backgroundColor: isOwnMessage ? '#e0f7fa' : '#f1f1f1' }}>
        <Text>{message}</Text>
      </Paper>
    </Group>
  );
};

export { ChatBubble };
