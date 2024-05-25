import { createHistoryStore } from "~/service/stores/ChatStore";
import { For } from "solid-js";

const ChatWindow = () => {
  const [messages, setMessageHistory] = createHistoryStore();
  const messageList = messages.messageHistory;

  return (
    <div class="border-slate-800 border p-10 h-full">
      <div>
        <h1 class="text-white">Chat Logs</h1>
      </div>
      <div>
        <For each={messageList}>{(message, i) => {
          return <p class="text-white">{message}</p>
        }}</For>
      </div>
    </div>
  );
}

export default ChatWindow;
