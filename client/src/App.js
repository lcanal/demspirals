import React, { Component } from 'react';
import { Navbar,Nav,NavItem } from 'react-bootstrap';
import { Link,Route,BrowserRouter as Router } from 'react-router-dom';
import { LinkContainer } from 'react-router-bootstrap';
import Home from './components/Home';
import Top10 from './components/Top10';
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
      <Router basename={process.env.PUBLIC_URL}>
        <div className="App">
        <Navbar inverse collapseOnSelect>
          <Navbar.Header>
            <Navbar.Brand>
              <Link to="/">Home</Link>
            </Navbar.Brand>
          <Navbar.Toggle />
          </Navbar.Header>
          <Navbar.Collapse>
            <Nav>
            <LinkContainer to="/top10">
              <NavItem eventKey={1} href="/top10">Top 10 Overall</NavItem>
            </LinkContainer>
            <LinkContainer to="/topQB">
              <NavItem eventKey={2} href="/topQB">Top QBs</NavItem>
            </LinkContainer>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
        <Route exact path="/" component={Home} />
        <Route path="/top10" component={Top10} />
      </div>
    </Router>
    );
  }
}



export default App;
