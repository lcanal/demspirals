import React,{ Component } from 'react';
import StatsTable from '../modules/StatsTable';
//import "../css/TopPosContainer.css";

class TopPosContainer extends Component{
    render(){
        var apiURL = "/api/topplayers/" + this.props.position
        return(
            <div className="top-overall-container">
                <StatsTable apiURL={apiURL} />
            </div>
        )
    }
}

export default TopPosContainer;