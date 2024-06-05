import { createStore } from "solid-js/store";
import { MessageType } from "~/components/Message";

export const createHistoryStore = () => {
  const [messages, setMessages] = createStore({
    messageHistory: [],
  });

  // TODO: Set message history type
  const setMessageHistory = (messageHistory: MessageType[]) => {
    setMessages('messageHistory', (history) => [...messageHistory]);
  }

  const addMessage = (message: MessageType, index: number) => {
    setMessages('messageHistory', index, message);
  }

  return [messages, setMessageHistory, addMessage];
}

