import React, { Component } from 'react';
import { Navbar,Nav,NavItem } from 'react-bootstrap';
import { Link,Route,BrowserRouter as Router } from 'react-router-dom';
import { LinkContainer } from 'react-router-bootstrap';
import Home from './components/Home';
import TopOverall from './components/TopOverall';
import './css/App.css';

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
            <LinkContainer to="/topOverall">
              <NavItem eventKey={1} href="/topOverall">Top Overall</NavItem>
            </LinkContainer>
            <LinkContainer to="/topQB">
              <NavItem eventKey={2} href="/topQB">Top Quarterbacks</NavItem>
            </LinkContainer>
            <LinkContainer to="/topWR">
              <NavItem eventKey={3} href="/topWR">Top Wideouts</NavItem>
            </LinkContainer>
            <LinkContainer to="/topRB">
              <NavItem eventKey={4} href="/topRB">Top Rushers</NavItem>
            </LinkContainer>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
          <Route exact path="/"     component={Home} />
          <Route path="/topOverall" component={TopOverall} />
          <Route path="/topQB"      component={() => <TopOverall statfilter="qb" />}/>
          <Route path="/topWR"      component={() => <TopOverall statfilter="wr" />}/>
          <Route path="/topRB"      component={() => <TopOverall statfilter="rb" />}/>
      </div>
    </Router>
    );
  }
}



export default App;
