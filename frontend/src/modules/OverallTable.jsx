import React, { Component } from 'react';
import { Table,Fade,ProgressBar } from 'react-bootstrap';
import '../css/TopOverall.css';

class OverallTable extends Component {
    state = {
        players: [],
        showTable: false,
        loadState: 0
    }

    updateProgress(oEvent){
        if (oEvent.lengthComputable) {
        var percentComplete = oEvent.loaded / oEvent.total;
           this.setState({
               "loadState": percentComplete
           })
        } else {
            //console.log("Byte size for computing % done is unknown.")
            this.setState({
                "loadState": this.state.loadState + 60
            })
        }
    }

    async componentDidMount(){
        
        //const response  = await fetch(this.props.apiURL)
        //const json      = await response.json()
        const response = await fetchOverallPlayers(this.props.apiURL,this.updateProgress.bind(this))
       // const json      = await response.json()
        this.setState({
            "players" : response.playerdata,
            "showTable": true
        })

        
    }

    render(){
        var rows = [];
        for (var index = 0; index < this.state.players.length; index++) {
            if (index > 10) {
                break;
            }
            var player = this.state.players[index]
            rows.push(<ResultEntry player={player} key={player.id} />)
        }
        
        return (
            <div>
            <Fade in={!this.state.showTable} unmountOnExit={true} >
                <ProgressBar className="stats-table-load-status" now={this.state.loadState} max={this.state.numLimit} />
            </Fade>
            <Fade in={this.state.showTable} transitionAppear={true}>
            <Table className="stats-table" condensed hover bordered responsive >
                <thead>
                    <tr><th colSpan="100"><h4>Top 10 {this.props.title}</h4></th></tr>
                    <tr>
                    <th>Player</th>
                    <th>Age</th>
                    <th>Number</th>
                    <th>Height</th>
                    <th>Weight</th>
                    <th>Position</th>
                    <th>Team</th>
                    <th>Total Fantasy Points</th>
                    </tr>
                </thead>
                <tbody>
                    {rows}
                </tbody>
            </Table>
            </Fade>
            </div>
        )
    }
}


class ResultEntry extends Component {
  render(){
      var nflprofileURL = "http://www.nfl.com/player/pid/"+this.props.player.nflid+"/profile"
      var nameField = <a href={nflprofileURL} target="_">{this.props.player.firstname} {this.props.player.lastname}</a>
       
    return(
      <tr className="stat-row" id={this.id}>
        <td>{nameField}</td>
        <td>{this.props.player.age}</td>
        <td>{this.props.player.jerseynum}</td>
        <td>{this.props.player.height}</td>
        <td>{this.props.player.weight}</td>
        <td>{this.props.player.position}</td>
        <td>{this.props.player.teamname}</td>
       <td>{this.props.player.totalfantasypoints.toFixed(2)}</td>
      </tr>
    )
  }
}

//This is to implement progress 
export const fetchOverallPlayers = (apiURL,progress) => {
    return new Promise((resolve,reject) => {
        var req = new XMLHttpRequest();
        req.addEventListener("progress",progress);
        //req.addEventListener("load",transferComplete);

        req.open("GET",apiURL);
        req.send();
        req.onreadystatechange = function(){
            if (req.readyState === XMLHttpRequest.DONE){
                let data = JSON.parse(req.responseText)
                resolve(data)
            }
        }
    });
}


/*
function transferComplete(evt) {
        console.log("The transfer is complete.");
        console.log("Event print ",evt)
        this.setState({
            "loadState": 100,
            "showTable": "true"
        })
    }
*/
export default OverallTable;