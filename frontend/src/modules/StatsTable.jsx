import React, { Component } from 'react';
import { Table } from 'react-bootstrap';

class StatTable extends Component {
    state = {
        players: [],
        headerType: "",
        playerHeaders: [],
        rows: []
    }

    async componentDidMount(){
        const response  = await fetch(this.props.apiURL)
        const json      = await response.json()
        var playerHeaders = []
        var rows = []; 
        var headerset = false;

        for (var index = 0; index < json.playerdata.length; index++) {
            /*if (index > 20) {
                break;
            }*/
            var player = json.playerdata[index]
            player.name = <strong>{player.firstname} {player.lastname}</strong>
            
            var playerstatsresponse =  await fetch("/api/player/"+player.id)
            var stats = await playerstatsresponse.json()
            rows.push(<ResultEntry player={player} key={player.id} playerstats={stats} />)

            if (!headerset){
               for (var idx in stats) {
                   if (stats.hasOwnProperty(idx)) {
                       if (stats[idx].LeagueName.length > 0){
                            playerHeaders.push(stats[idx].LeagueName)
                       }else{
                           playerHeaders.push(stats[idx].Name)
                       }
                   }
               }
               headerset = true;
            }
        }
       
        
        this.setState({
            "players" : json.playerdata,
            "playerHeaders" : playerHeaders,
            "headerType": json.datatype,
            "rows" : rows
        })
    }

    render(){
        var headers = []
        if (this.state.headerType === "topwr"){
            headers.push(<th>Player</th>)
            headers.push(<th>Position</th>)
            headers.push(<th>Team</th>)
            //Build headers as we get them from the api
            this.state.playerHeaders.forEach(function(header) {
                headers.push(<th>{header}</th>)
            }, this);
            headers.push(<th>Total Points</th>)
        }
        return (
            <Table className="stats-table"  hover bordered responsive >
                <thead><tr>
                    {headers}
                </tr></thead>
                <tbody>
                    {this.state.rows}
                </tbody>
            </Table>
        )
    }
}


class ResultEntry extends Component {
  render(){
      var statsTD = []
      var stats = this.props.playerstats
      for (var index in stats) {
          if (stats.hasOwnProperty(index)) {
             if (stats[index].LeagueName.length > 0) {
                statsTD.push(<td>{stats[index].Value.toFixed(2)}</td>)
             }else{
                statsTD.push(<td>{stats[index].StatNum.toFixed(2)}</td>)
             }
          }
      }
    return(
      <tr className="stat-row" id={this.id}>
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