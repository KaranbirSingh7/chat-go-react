import React, { Component } from 'react'
import "./ChatInput.scss";

export default class ChatInput extends Component {
  render() {
    return (
      <div className="ChatInput">
          <input type="text" onKeyDown={this.props.send} /> 
      </div>
    )
  }
}
