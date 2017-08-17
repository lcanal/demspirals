import React, { Component } from 'react';
import { Navbar,Nav,NavItem } from 'react-bootstrap';
import { Link,Route,BrowserRouter as Router } from 'react-router-dom';
import { LinkContainer } from 'react-router-bootstrap';
import Home from './components/Home';
import TopOverall from './components/TopOverall';
import TopPosition from './components/TopPosition';
import GraphPlayers from './components/GraphPlayers';
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
            <LinkContainer to="/app/topOverall">
              <NavItem eventKey={1} href="/app/topOverall">Top Overall</NavItem>
            </LinkContainer>
            <LinkContainer to="/app/topQB">
              <NavItem eventKey={2} href="/app/topQB">Top Quarterbacks</NavItem>
            </LinkContainer>
            <LinkContainer to="/app/topWR">
              <NavItem eventKey={3} href="/app/topWR">Top Wideouts</NavItem>
            </LinkContainer>
            <LinkContainer to="/app/topRB">
              <NavItem eventKey={4} href="/app/topRB">Top Rushers</NavItem>
            </LinkContainer>
            <LinkContainer to="/app/topTE">
              <NavItem eventKey={5} href="/app/topTE">Top Tight Ends</NavItem>
            </LinkContainer>
            </Nav>
            <Nav pullRight>
              <LinkContainer to="/app/graphQB">
                <NavItem eventKey={6} href="/app/graphQB">QB Chart</NavItem>
              </LinkContainer>
              <LinkContainer to="/app/graphWR">
                <NavItem eventKey={6} href="/app/graphWR">WR Chart</NavItem>
              </LinkContainer>
              <LinkContainer to="/app/graphRB">
                <NavItem eventKey={6} href="/app/graphRB">RB Chart</NavItem>
              </LinkContainer>
              <LinkContainer to="/app/graphTE">
                <NavItem eventKey={6} href="/app/graphTE">TE Chart</NavItem>
              </LinkContainer>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
        
          <Route exact path="/"     component={Home} />
          <Route path="/app/topOverall" component={TopOverall} />
          <Route path="/app/topQB"      component={() => <TopPosition position="qb" />}/>
          <Route path="/app/topWR"      component={() => <TopPosition position="wr" />}/>
          <Route path="/app/topRB"      component={() => <TopPosition position="rb" />}/>
          <Route path="/app/topTE"      component={() => <TopPosition position="te" />}/>

          <Route path="/app/graphQB"      component={() => <GraphPlayers position="qb" />}/>
          <Route path="/app/graphWR"      component={() => <GraphPlayers position="wr" />}/>
          <Route path="/app/graphRB"      component={() => <GraphPlayers position="rb" />}/>
          <Route path="/app/graphTE"      component={() => <GraphPlayers position="te" />}/>
      </div>
    </Router>
    );
  }
}



export default App;
