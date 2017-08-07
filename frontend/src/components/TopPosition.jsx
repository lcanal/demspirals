import React, { Component } from 'react';
import { Panel } from 'react-bootstrap';
import TopPosContainer from '../containers/TopPosContainer';
class TopPosition extends Component {
    render(){
        return (
        <Panel>
        <TopPosContainer position={this.props.position}/>  
        </Panel>
        );
    }
}

export default TopPosition;