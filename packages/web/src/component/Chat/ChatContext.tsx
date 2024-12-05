import { createContext, ReactNode, useContext } from 'react';
import { useUserStore } from '../../store/useUserStore.tsx';
import useWebSocket from 'react-use-websocket';
import { WebSocketMessage } from 'react-use-websocket/dist/lib/types';
import { notifications } from '@mantine/notifications';

type ChatContextProp = {
  sendMessage: (message: WebSocketMessage) => void;
};

const ChatContext = createContext<ChatContextProp | undefined>(undefined);

type ChatProviderProps = {
  children: ReactNode;
};

const ChatProvider = ({ children }: ChatProviderProps) => {
  const { user } = useUserStore();
  const url = `http://localhost:8080/chat/${user?.userId}`;
  const { sendMessage } = useWebSocket(url, {
    onOpen() {
      notifications.show({
        title: 'Connection Success',
        message: 'Do not forget to star Mantine on GitHub! 🌟',
      });
    },
  });
  return (
    <ChatContext.Provider
      value={{
        sendMessage,
      }}
    >
      {children}
    </ChatContext.Provider>
  );
};

const useChatContext = () => {
  const context = useContext(ChatContext);
  if (!context) {
    throw new Error('useChatContext must be used within a ProductProvider');
  }
  return context;
};

export { ChatProvider, useChatContext };
