import React, { Component } from 'react';
import { Table,Fade } from 'react-bootstrap';
import '../css/TopOverall.css';

class OverallTable extends Component {
    state = {
        players: [],
        showTable: false
    }

    async componentDidMount(){
        const response  = await fetch(this.props.apiURL)
        const json      = await response.json()
        
        this.setState({
            "players" : json.playerdata,
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

export default OverallTable;