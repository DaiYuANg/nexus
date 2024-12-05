import { Box } from '@mantine/core';
import { ChatRoot } from '../../component/Chat/ChatRoot.tsx';

const Chat = () => {
  // const [messages, setMessages] = useState<string[]>([]);
  // const [messageInput, setMessageInput] = useState('');

  // const handleSendMessage = () => {
  //   if (messageInput.trim() === '') return; // Do nothing if input is empty
  //   setMessages([...messages, messageInput]);
  //   setMessageInput('');
  // };
  //
  // const handleKeyDown = (event: KeyboardEvent) => {
  //   // Check if Enter key was pressed (without shift, so it creates a new line in textarea)
  //   if (event.key === 'Enter' && !event.shiftKey) {
  //     event.preventDefault(); // Prevent the default behavior of Enter key (adding a new line)
  //     handleSendMessage(); // Send the message
  //   }
  // };

  return (
    <Box style={{ height: '80vh', display: 'flex', flexDirection: 'column' }}>
      <ChatRoot />
    </Box>
  );
};

export default Chat;
export { Chat };
