import React,{ Component } from 'react';
import StatTable from '../modules/StatTable';

class TopContainer extends Component{
    apiURL = this.props.dataFrom
    render(){
        return(
            <StatTable />
        )
    }
}

export default TopContainer;