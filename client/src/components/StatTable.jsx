import React, { Component } from 'react';

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
        this.state.players.forEach(  player => {
            rows.push(<tr><td>{player.name}</td><td>{player.position}</td></tr>)
        });
        return (
            <table>
                <thead><tr>
                    <th>Player</th>
                    <th>Position</th>
                </tr></thead>
                <tbody>
                    {rows}
                </tbody>
            </table>
        )
    }
}

export default StatTable;