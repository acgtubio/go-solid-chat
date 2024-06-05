export type MessageType = {
  content: string
  author: string
}

export default function Message(props: MessageType) {
  return (
    <>
      <p class="text-white">{props.content}</p>
    </>
  )
}
