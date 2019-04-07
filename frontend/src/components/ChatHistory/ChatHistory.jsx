import React, { Component } from "react";
import "./ChatHistory.scss";

export default class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory.map((msg, index) => {
      return <li key={index}>{msg.data}</li>;
    });

    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        <ul>{messages}</ul>
      </div>
    );
  }
}
