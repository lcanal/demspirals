import React, { Component } from 'react';
import { Panel } from 'react-bootstrap';
import OverallTable from '../modules/OverallTable';
import "../css/TopOverall.css";
class TopOverall extends Component {
    render(){
        var apiURL = "/api/topplayers"
        var qbURL = apiURL + "/qb"
        var wrURL = apiURL + "/wr"
        var rbURL = apiURL + "/rb"
        
        return (
        <Panel>
            <div className="top-overall-container">
                <OverallTable apiURL={qbURL} title="Quarterbacks" />
                <OverallTable apiURL={wrURL} title="Wideouts" />
                <OverallTable apiURL={rbURL} title="Running Backs" />
            </div>
        </Panel>
        );
    }
}

export default TopOverall;