const socket = new WebSocket("ws://localhost:8080/ws");

const connect = () => {
    console.log("Connecting...");

    socket.onopen = () => {
        console.log("Successfully connected");
    }

    socket.onmessage = (msg) => {
        console.log(msg);
    }

    socket.onclose = (e) => {
        console.log("Socket closed: ", e);
    }

    socket.onerror = (e) => {
        console.error("Socket error: ", e);
    }
}

const sendMsg = (msg) => {
    console.log("Sending message: ", msg);
    socket.send(msg);
}

export { connect, sendMsg };
