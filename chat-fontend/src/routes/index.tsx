import { Title } from "@solidjs/meta";
import { clientOnly } from "@solidjs/start";

const WsButton = clientOnly(() => {
  return import("~/components/WebsocketButton");
});

const ChatWindow = clientOnly(() => {
  return import("~/components/ChatWindow");
});

export default function Home() {

  return (
    <main class="flex justify-center">
      <Title>Some chat site</Title>
      <section class="h-screen w-3/4 p-10">
        <ChatWindow />
      </section>
    </main>
  );
}
