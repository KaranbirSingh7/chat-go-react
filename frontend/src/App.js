import React, { Component } from 'react';
import './App.css';
import {connect, sendMsg} from "./api"

class App extends Component {
  constructor(props){
    super(props)
    // Initialize socket connection
    connect()
  }

  // send func to send messages
  send(){
    sendMsg("HELLO from React ")
  }

  render() {
    return (
      <div className="App">
      <button onClick={this.send}>Send Message</button>
      </div>
    );
  }
}

export default App;
