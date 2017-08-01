import React, { Component } from 'react';
import { Panel } from 'react-bootstrap';
import TopContainer from '../containers/TopContainer';
class Top10 extends Component {
    render(){
        return (
        <Panel>
        <TopContainer dataFrom="/api/top10s" />  
        </Panel>
        );
    }
}

export default Top10;