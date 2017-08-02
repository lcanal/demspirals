import React, { Component } from 'react';
import { Table } from 'react-bootstrap';

class StatTable extends Component {
    state = {
        players: []
    }

    async componentDidMount(){
        const response  = await fetch("/api/topten")
        const json      = await response.json()
        
        this.setState({
            "players" : json
        })
    }

    render(){
        var rows = [];
        this.state.players.forEach( player => {
            rows.push(<ResultEntry player={player} key={player.slug} />)
        });
        return (
            <Table condensed >
                <thead><tr>
                    <th>Player</th>
                    <th>Pos</th>
                    <th>Team</th>
                    <th>Rec.Yards</th>
                    <th>Rush.Yards</th>
                    <th>Rush.Attmp</th>
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
      <tr id={this.key}>
        <td>{this.props.player.name}</td>
        <td>{this.props.player.pos}</td>
        <td>{this.props.player.team.name}</td>
        <td>{this.props.player.stats.Receptionyards}</td>
        <td>{this.props.player.stats.Rushyards}</td>
        <td>{this.props.player.stats.Rushattempts}</td>
      </tr>
    )
  }
}

export default StatTable;