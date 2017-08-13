import React, { Component } from 'react';
import { Navbar,Nav,NavItem } from 'react-bootstrap';
import { Link,Route,BrowserRouter as Router } from 'react-router-dom';
import { LinkContainer } from 'react-router-bootstrap';
import Home from './components/Home';
import TopOverall from './components/TopOverall';
import TopPosition from './components/TopPosition';
import './css/App.css';

class App extends Component {
  render() {
    return (
      <Router basename={process.env.PUBLIC_URL}>
        <div className="App">
        <Navbar collapseOnSelect>
          <Navbar.Header>
            <Navbar.Brand>
              <Link to="/">Home</Link>
            </Navbar.Brand>
          <Navbar.Toggle />
          </Navbar.Header>
          <Navbar.Collapse>
            <Nav>
            <LinkContainer to="/stats/topOverall">
              <NavItem eventKey={1} href="/stats/topOverall">Top Overall</NavItem>
            </LinkContainer>
            <LinkContainer to="/stats/topQB">
              <NavItem eventKey={2} href="/stats/topQB">Top Quarterbacks</NavItem>
            </LinkContainer>
            <LinkContainer to="/stats/topWR">
              <NavItem eventKey={3} href="/stats/topWR">Top Wideouts</NavItem>
            </LinkContainer>
            <LinkContainer to="/stats/topRB">
              <NavItem eventKey={4} href="/stats/topRB">Top Rushers</NavItem>
            </LinkContainer>
            <LinkContainer to="/stats/topTE">
              <NavItem eventKey={5} href="/stats/topTE">Top Tight Ends</NavItem>
            </LinkContainer>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
          <Route exact path="/"     component={Home} />
          <Route path="/stats/topOverall" component={TopOverall} />
          <Route path="/stats/topQB"      component={() => <TopPosition position="qb" />}/>
          <Route path="/stats/topWR"      component={() => <TopPosition position="wr" />}/>
          <Route path="/stats/topRB"      component={() => <TopPosition position="rb" />}/>
          <Route path="/stats/topTE"      component={() => <TopPosition position="te" />}/>
      </div>
    </Router>
    );
  }
}



export default App;
