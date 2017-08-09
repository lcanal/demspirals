import React,{ Component } from 'react';
import StatsTable from '../modules/StatsTable';
//import "../css/TopPosition.css";

class TopPosition extends Component{
    render(){
        var apiURL = "/api/topplayers/" + this.props.position
        return(
            <div className="top-overall-container">
                <StatsTable apiURL={apiURL} position={this.props.position}/>
            </div>
        )
    }
}

export default TopPosition;