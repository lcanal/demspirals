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
        this.state.players.forEach( player => {
            rows.push(<ResultEntry player={player} key={player.slug} />)
        });
        return (
            <Table condensed hover bordered >
                <thead><tr>
                    <th>Player</th>
                    <th>Pos</th>
                    <th>Team</th>
                    <th>Gamesplayed</th>
                    <th>Rushattempts</th>
                    <th>Rushyards10p</th>
                    <th>Rushyards20p</th>
                    <th>Rushyards30p</th>
                    <th>Rushyards40p</th>
                    <th>Rushyards50p</th>
                    <th>Runs</th>
                    <th>Passes</th>
                    <th>Receptions</th>
                    <th>Receptionyards</th>
                    <th>Receptiontargets</th>
                    <th>Receptionyards10p</th>
                    <th>Receptionyards20p</th>
                    <th>Receptionyards30p</th>
                    <th>Receptionyards40p</th>
                    <th>Receptionyards50p</th>
                    <th>Rushyards</th>
                    <th>Rushattempts</th>
                    <th>Rushyards10p</th>
                    <th>Rushyards20p</th>
                    <th>Rushyards30p</th>
                    <th>Rushyards40p</th>
                    <th>Rushyards50p</th>
                    <th>Passyards</th>
                    <th>Passattempts</th>
                    <th>Passyards10p</th>
                    <th>Passyards20p</th>
                    <th>Passyards30p</th>
                    <th>Passyards40p</th>
                    <th>Passyards50p</th>
                    <th>Touchdownpasses</th>
                    <th>Touchdownrushes</th>
                    <th>Fumbles</th>
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
        <td>{this.props.player.stats.Gamesplayed}</td>
        <td>{this.props.player.stats.Rushattempts}</td>
        <td>{this.props.player.stats.Rushyards10p}</td>
        <td>{this.props.player.stats.Rushyards20p}</td>
        <td>{this.props.player.stats.Rushyards30p}</td>
        <td>{this.props.player.stats.Rushyards40p}</td>
        <td>{this.props.player.stats.Rushyards50p}</td>
        <td>{this.props.player.stats.Runs}</td>
        <td>{this.props.player.stats.Passes}</td>
        <td>{this.props.player.stats.Receptions}</td>
        <td>{this.props.player.stats.Receptionyards}</td>
        <td>{this.props.player.stats.Receptiontargets}</td>
        <td>{this.props.player.stats.Receptionyards10p}</td>
        <td>{this.props.player.stats.Receptionyards20p}</td>
        <td>{this.props.player.stats.Receptionyards30p}</td>
        <td>{this.props.player.stats.Receptionyards40p}</td>
        <td>{this.props.player.stats.Receptionyards50p}</td>
        <td>{this.props.player.stats.Rushyards}</td>
        <td>{this.props.player.stats.Rushattempts}</td>
        <td>{this.props.player.stats.Rushyards10p}</td>
        <td>{this.props.player.stats.Rushyards20p}</td>
        <td>{this.props.player.stats.Rushyards30p}</td>
        <td>{this.props.player.stats.Rushyards40p}</td>
        <td>{this.props.player.stats.Rushyards50p}</td>
        <td>{this.props.player.stats.Passyards}</td>
        <td>{this.props.player.stats.Passattempts}</td>
        <td>{this.props.player.stats.Passyards10p}</td>
        <td>{this.props.player.stats.Passyards20p}</td>
        <td>{this.props.player.stats.Passyards30p}</td>
        <td>{this.props.player.stats.Passyards40p}</td>
        <td>{this.props.player.stats.Passyards50p}</td>
        <td>{this.props.player.stats.Touchdownpasses}</td>
        <td>{this.props.player.stats.Touchdownrushes}</td>
        <td>{this.props.player.stats.Fumbles}</td>
      </tr>
    )
  }
}

export default StatTable;