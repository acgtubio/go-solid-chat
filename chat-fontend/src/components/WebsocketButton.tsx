import { createEffect } from "solid-js";
import { connect, sendMsg } from "~/service/websocket";

const WebSocketButton = () => {

  createEffect(() => {
    connect();
  });

  const send = (message: string) => {
    sendMsg(message);
  }

  return (
    <>
      <button onClick={() => { send("Wassup") }}>Send Message</button>
    </>
  )
}

export default WebSocketButton;
