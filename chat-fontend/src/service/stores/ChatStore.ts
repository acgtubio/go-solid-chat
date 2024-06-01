import { createStore } from "solid-js/store";

export const createHistoryStore = () => {
  const [messages, setMessages] = createStore({
    messageHistory: [],
  });

  // TODO: Set message history type
  const setMessageHistory = (messageHistory: any[]) => {
    setMessages('messageHistory', (history) => [...messageHistory]);
  }

  const addMessage = (message: any, index: number) => {
    setMessages('messageHistory', index + 1, message);
  }

  return [messages, setMessageHistory, addMessage];
}

