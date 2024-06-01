import { createHistoryStore } from "~/service/stores/ChatStore";
import { For, createEffect, createSignal } from "solid-js";
import { createChatWebsocket } from "~/service/websocket";
//import { connect, sendMsg } from "~/service/websocket";

const ChatWindow = () => {
  const [messages, setMessageHistory, addMessage] = createHistoryStore();
  const [connect, sendMsg] = createChatWebsocket();
  const messageList = messages.messageHistory;

  let messageInput!: HTMLSpanElement;

  createEffect(() => {
    connect((msg: any) => {
      addMessage(msg, messageList.length);
    });
  });

  const focusOnInput = () => {
    messageInput.focus();
  }

  const messageInputKeyboardEvent = (e: KeyboardEvent) => {
    if (e.key != 'Enter') {
      return;
    }
    e.preventDefault();

    sendMsg(messageInput.textContent);
  }

  return (
    <div class="border-slate-800 border rounded-xl p-10 h-full relative">
      <div class="mb-2">
        <h1 class="text-white">Chat Logs</h1>
      </div>
      <div>
        <For each={messageList}>{(message, i) => {
          return <p class="text-white">{message}</p>
        }}</For>
      </div>
      <div
        class="absolute bottom-10 right-10 left-10 py-3 px-5 border border-slate-700 background rounded-xl has-[:focus]:border-slate-500"
        onClick={focusOnInput}
      >
        <span
          ref={messageInput}
          onKeyDown={messageInputKeyboardEvent}
          class="text-white w-full block focus:outline-none"
          contenteditable
        >
        </span>
      </div>
    </div>
  );
}

export default ChatWindow;
