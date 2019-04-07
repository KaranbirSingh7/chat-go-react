import React, { Component } from "react";
// import './App.css';
import Header from "./components/Header/Header";

import { connect, sendMsg } from "./api";
import ChatHistory from "./components/ChatHistory/ChatHistory";

class App extends Component {
  state = {
    chatHistory: []
  };

  componentDidMount() {
    connect(msg => {
      console.log("New Message");
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }));
      console.log(this.state);
    });
  }

  // send func to send messages
  send() {
    sendMsg("HELLO from React ");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Send Message</button>
      </div>
    );
  }
}

export default App;
