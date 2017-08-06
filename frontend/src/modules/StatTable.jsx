import React, { Component } from 'react';
import { Table } from 'react-bootstrap';

class StatTable extends Component {
    state = {
        players: []
    }

    async componentDidMount(){
        const response  = await fetch("/api/topoverall")
        const json      = await response.json()
        
        this.setState({
            "players" : json
        })
    }

    render(){
        var rows = []; 
        for (var index = 0; index < this.state.players.length; index++) {
            if (index > 20) {
                break;
            }
            var player = this.state.players[index]
            rows.push(<ResultEntry player={player} key={player.id} />)
        }

        /*this.state.players.forEach( player => {
            rows.push(<ResultEntry player={player} key={player.ID} />)
        });*/
        return (
            <Table className="stats-table" condensed hover bordered responsive >
                <thead><tr>
                    <th>Player</th>
                    <th>Position</th>
                    <th>Team</th>
                   <th>Total Fantasy Points</th>
                   <th>Field</th>
                   <th>Field</th>
                </tr></thead>
                <tbody>
                    {rows}
                </tbody>
            </Table>
        )
    }
}


class ResultEntry extends Component {
  render(){
    return(
      <tr className="stat-row" id={this.id}>
        <td>{this.props.player.firstname} {this.props.player.lastname}</td>
        <td>{this.props.player.position}</td>
        <td>{this.props.player.teamname} </td>
       <td>{this.props.player.totalfantasypoints.toFixed(2)}</td>
       <td>NA</td>
       <td>NA</td>
      </tr>
    )
  }
}

export default StatTable;