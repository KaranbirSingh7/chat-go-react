let socketURL = "ws://localhost:8000/ws"
let socket = new WebSocket(socketURL)

// Connect method handles whole lifecycle of a single conenection
let connect = (cb) => {
  console.log(`Attempting connection at ${socketURL}`);

  // CONNECTION CREATED
  socket.onopen = () => {
    console.log(`Successfully connected`)
  };

  // MESSAGE RECEIVED
  socket.onmessage = (msg) => {
    console.log(msg);
    cb(msg)
  } 

  // CONNECTION CLOSED
  socket.onclose = (event) => {
    console.log(`Connection closed: ${event}`)    
  }

  socket.onerror = error => {
    console.log(`Socker error: ${error}`)
  }
};

// Method to send message
let sendMsg = msg => {
  console.log(`Sending msg: ${msg}`);
  socket.send(msg);
}

// Export the above created methods
export {connect, sendMsg};