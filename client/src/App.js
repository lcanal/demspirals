import React, { Component } from 'react';
import StatTable from './components/StatTable';
import './App.css';

class App extends Component {
  state = {
    welcomeword: "None"
  }

  async componentDidMount(){
    const response = await fetch("/api/hello")
    const json = await response.json()   

    this.setState({
      "welcomeword" : json.hello
    })
  }

  render() {
    return (
      <div className="App">
        <StatTable />
      </div>
    );
  }
}

export default App;
