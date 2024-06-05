export const createChatWebsocket = () => {
    const socket = new WebSocket("ws://localhost:8080/ws");

    const connect = (cb) => {
        console.log("Connecting...");

        socket.onopen = () => {
            console.log("Successfully connected");
        }

        socket.onmessage = (msg) => {
            const parsedData = JSON.parse(msg.data);
            cb(parsedData.body);
        }

        socket.onclose = (e) => {
            console.log("Socket closed: ", e);
        }

        socket.onerror = (e) => {
            console.error("Socket error: ", e);
        }
    }

    const sendMsg = (msg) => {
        socket.send(JSON.stringify(msg));
    }

    return [connect, sendMsg];
}
