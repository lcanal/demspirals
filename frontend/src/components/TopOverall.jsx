import React, { Component } from 'react';
import { Panel } from 'react-bootstrap';
import TopContainer from '../containers/TopContainer';
class TopOverall extends Component {
    render(){
        return (
        <Panel>
        <TopContainer statfilter={this.props.statfilter}/>  
        </Panel>
        );
    }
}

export default TopOverall;