import React, { Component } from "react";
// import './App.css';
import Header from "./components/Header/Header";

import { connect, sendMsg } from "./api";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChanInput from "./components/ChatInput/ChatInput";

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
  send(event) {
    // if ENTER(13) is pressed
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }

  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChanInput send={this.send} />
      </div>
    );
  }
}

export default App;
