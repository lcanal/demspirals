import React, { Component } from 'react';
import { Panel } from 'react-bootstrap';
import TopContainer from '../containers/TopContainer';
class TopOverall extends Component {
    render(){
        return (
        <Panel>
        <TopContainer dataFrom="/api/topoverall" />  
        </Panel>
        );
    }
}

export default TopOverall;