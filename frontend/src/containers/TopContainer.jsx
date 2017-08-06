import React,{ Component } from 'react';
import StatTable from '../modules/StatTable';
import "../css/TopContainer.css";

class TopContainer extends Component{
    render(){
        var apiURL = "/api/topoverall"
        if (typeof this.props.statfilter !== 'undefined' && this.props.statfilter.length > 0){
            apiURL = apiURL + "/" + this.props.statfilter
        }
        return(
            <div className="top-overall-container">
                <StatTable apiURL={apiURL} />
            </div>
        )
    }
}

export default TopContainer;