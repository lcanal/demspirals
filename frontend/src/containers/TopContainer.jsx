import React,{ Component } from 'react';
import StatTable from '../modules/StatTable';
import "../css/TopContainer.css";

class TopContainer extends Component{
    apiURL = this.props.dataFrom
    render(){
        return(
            <div className="top-overall-container">
                <StatTable />
            </div>
        )
    }
}

export default TopContainer;