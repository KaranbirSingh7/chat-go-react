import React, { Component } from "react";
import "./ChatHistory.scss";
import Message from "../Message/Message";

export default class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory.map((msg, index) =>{
      return <Message key={index} message={msg.data}/>
    });

    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {messages}
      </div>
    );
  }
}
