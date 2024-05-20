import { Title } from "@solidjs/meta";
import { createEffect } from "solid-js";
import { effect } from "solid-js/web";
import Counter from "~/components/Counter";
import { connect, sendMsg } from "~/service/websocket";
import { clientOnly } from "@solidjs/start";

const WsButton = clientOnly(() => {
  return import("~/components/WebsocketButton");
});

export default function Home() {

  return (
    <main>
      <Title>Hello World</Title>
      <WsButton />
    </main>
  );
}
