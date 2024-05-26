import { createHistoryStore } from "~/service/stores/ChatStore";
import { For, createEffect, createSignal } from "solid-js";
import { createChatWebsocket } from "~/service/websocket";
//import { connect, sendMsg } from "~/service/websocket";

const ChatWindow = () => {
  const [messages, setMessageHistory, addMessage] = createHistoryStore();
  const [test1, settest1] = createSignal("asdf");
  const [connect, sendMsg] = createChatWebsocket();
  const messageList = messages.messageHistory;

  createEffect(() => {
    connect((msg: any) => {
      addMessage(msg, messageList.length);
    });
  });

  return (
    <div class="border-slate-800 border p-10 h-full">
      <div>
        <h1 class="text-white">Chat Logs</h1>
        <button onClick={() => settest1(test1() + " adf")} class="text-white"> Click Me </button>
        <button onClick={() => sendMsg("hey hey")} class="text-white"> Click Me </button>
      </div>
      <p class="text-white">Text: {test1()}</p>
      <div>
        <For each={messageList}>{(message, i) => {
          return <p class="text-white">{message}</p>
        }}</For>
      </div>
    </div>
  );
}

export default ChatWindow;
