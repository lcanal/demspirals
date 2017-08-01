import React, { Component } from 'react';
import { Jumbotron } from 'react-bootstrap';


class Home extends Component {
    render(){
        return (
        <Jumbotron>
            <h1>You ready to lob these?!</h1>
            <p>
                Use the nav bars to the top to explore.
            </p>
        </Jumbotron>
        )
    }
}

export default Home;