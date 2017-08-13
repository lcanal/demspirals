import React, { Component } from 'react';
import { Jumbotron } from 'react-bootstrap';

class Home extends Component {
    render(){
        return (
        <Jumbotron className="welcome-jumbotron" >
            <h1>You ready to lob these?!</h1>
            <img width={500} height={300} alt="500x300" src="http://www.pbh2.com/wordpress/wp-content/uploads/2013/12/doucet-across-the-middle.gif"/>
        </Jumbotron>
        )
    }
}

export default Home;