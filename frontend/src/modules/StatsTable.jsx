import React, { Component } from 'react';
import { Table,Fade,ProgressBar } from 'react-bootstrap';
import "../css/StatsTable.css";

class StatTable extends Component {
    state = {
        players: [],
        playerHeaders: [],
        rows: [],
        showTable: false,
        loadState: 0.0,
        numLimit: 20
    }

    async componentDidMount(){
        const response  = await fetch(this.props.apiURL)
        const json      = await response.json()
        var playerHeaders = []
        var rows = []; 
        var headerset = false;
        var position = this.props.position

        for (var index = 0; index < json.playerdata.length; index++) {
            if (index > this.state.numLimit) {
                break;
            }
            var player = json.playerdata[index]
            player.name = <strong>{player.firstname} {player.lastname}</strong>
            
            //console.log("Fetching: http://localhost:8080/api/player/"+position+"/"+player.id)
            var playerstatsresponse =  await fetch("/api/player/"+position+"/"+player.id)
            var stats = await playerstatsresponse.json()
            rows.push(<ResultEntry player={player} key={player.id} pid={player.id} playerstats={stats} />)

            if (!headerset){
               for (var idx in stats) {
                   if (stats.hasOwnProperty(idx)) {
                       if (stats[idx].leaguename.length > 0){
                            playerHeaders.push(stats[idx].leaguename)
                       }else{
                           playerHeaders.push(stats[idx].name)
                       }
                   }
               }
               headerset = true;
            }

            this.setState({
                "loadState": index
            })
        }
       
        
        this.setState({
            "players" : json.playerdata,
            "playerHeaders" : playerHeaders,
            "rows" : rows,
            "showTable": true
        })
    }

    render(){
        var headers = []
        headers.push(<th key="Player">Player</th>)
        headers.push(<th key="Position">Position</th>)
        headers.push(<th key="Team">Team</th>)
        //Build headers as we get them from the api
        this.state.playerHeaders.forEach(function(header) {
            headers.push(<th key={header}>{header}</th>)
        }, this);
        headers.push(<th key="TotalPoints">Total Points</th>)
        return (
            <div>
            <Fade in={!this.state.showTable} unmountOnExit={true} >
                <ProgressBar className="stats-table-load-status" now={this.state.loadState} max={this.state.numLimit} />
            </Fade>
            <Fade in={this.state.showTable} transitionAppear={true} >
            <Table className="stats-table"  condensed hover bordered responsive >
                <thead><tr>
                    {headers}
                </tr></thead>
                <tbody>
                    {this.state.rows}
                </tbody>
            </Table>
            </Fade>
            </div>
        )
    }
}


class ResultEntry extends Component {
  render(){
      var statsTD = []
      var stats = this.props.playerstats
      for (var index in stats) {
          if (stats.hasOwnProperty(index)) {
             if (stats[index].leaguename.length > 0) {
                statsTD.push(<td key={stats[index].playerid+"-"+stats[index].name}>{stats[index].value.toFixed(2)}</td>)
             }else{
                statsTD.push(<td key={stats[index].playerid+"-"+stats[index].name}>{stats[index].statnum.toFixed(2)}</td>)
             }
          }
      }
    return(
      <tr className="stat-row" key={this.props.pid}>
        <td>{this.props.player.name}</td>
        <td>{this.props.player.position}</td>
        <td>{this.props.player.teamname} </td>
        {statsTD}
        <td>{this.props.player.totalfantasypoints.toFixed(2)}</td>
      </tr>
    )
  }
}

export default StatTable;