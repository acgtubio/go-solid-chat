import { createStore } from "solid-js/store";

export const createHistoryStore = () => {
  const [messages, setMessages] = createStore({
    messageHistory: ["a", "b", "c", "d"],
  });

  // TODO: Set message history type
  const setMessageHistory = (messageHistory: any[]) => {
    setMessages('messageHistory', (history) => [...messageHistory]);
  }

  return [messages, setMessageHistory];
}

