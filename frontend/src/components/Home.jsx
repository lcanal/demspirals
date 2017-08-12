import React, { Component } from 'react';
import { Jumbotron } from 'react-bootstrap';

class Home extends Component {
    render(){
        return (
        <Jumbotron className="welcome-jumbotron" >
            <h1>You ready to lob these?!</h1>
        </Jumbotron>
        )
    }
}

export default Home;