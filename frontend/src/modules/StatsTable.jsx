import React, { Component } from 'react';
import { Fade,ProgressBar } from 'react-bootstrap';
import { BootstrapTable, TableHeaderColumn } from 'react-bootstrap-table';
import '../css/react-bootstrap-table-all.min.css';
import "../css/StatsTable.css";

const selectRowProp = {
  mode: 'checkbox',
  clickToSelect: true  // enable click to select
};


class StatTable extends Component {
    state = {
        players: [],
        playerHeaders: [],
        showTable: false,
        loadState: 0.0,
        numLimit: 0
    }

    async componentDidMount(){
        const response  = await fetch(this.props.apiURL)
        const json      = await response.json()
        var playerHeaders = []
        var players = []

        var headerset = false;
        var position = this.props.position

        this.setState({
            "numLimit" : json.playerdata.length
        });

        for (var index = 0; index < json.playerdata.length; index++) {
            var player = json.playerdata[index]
            player.name = player.firstname + " " + player.lastname

            var playerstatsresponse =  await fetch("/api/player/"+position+"/"+player.id)
            var stats = await playerstatsresponse.json()

            //Set once for headers.
            if (!headerset){
               for (var hidx in stats) {
                   if (stats.hasOwnProperty(hidx)) {
                       if (stats[hidx].leaguename.length > 0){
                            playerHeaders.push(stats[hidx].leaguename)
                       }else{
                            playerHeaders.push(stats[hidx].name)
                       }
                   }
               }
               headerset = true;
            }

            //Actually grab data in headers
            for (var pidx in stats) {
                   if (stats.hasOwnProperty(pidx)) {
                       if (stats[pidx].leaguename.length > 0){
                            var statkey = stats[pidx].leaguename
                            player[statkey] = stats[pidx].value
                       }else{
                            var statkey = stats[pidx].name
                            player[statkey] = stats[pidx].statnum
                       }
                   }
               }


            this.setState({
                "loadState": index
            })

            players.push(player)
        }
       
        this.setState({
            "players" : players,
            "playerHeaders" : playerHeaders,
            "showTable": true
        })
    }

    render(){
        var headers = []
        headers.push(<TableHeaderColumn isKey={true} dataField="id">#</TableHeaderColumn>)
        headers.push(<TableHeaderColumn dataField="name">Player</TableHeaderColumn>)
       
        headers.push(<TableHeaderColumn dataField="position">Position</TableHeaderColumn>)
        headers.push(<TableHeaderColumn dataField="teamname">Team</TableHeaderColumn>)
        //Build headers as we get them from the api
        this.state.playerHeaders.forEach(function(header) {
            headers.push(<TableHeaderColumn dataField={header}>{header}</TableHeaderColumn>)
            
        }, this);
        headers.push(<TableHeaderColumn dataField="TotalPoints">Total Points</TableHeaderColumn>)
        return (
            <div>
            <Fade in={!this.state.showTable} unmountOnExit={true} >
                <ProgressBar className="stats-table-load-status" now={this.state.loadState} max={this.state.numLimit} />
            </Fade>
            <Fade in={this.state.showTable} transitionAppear={true} >
            <BootstrapTable className="stats-table" selectRow={ selectRowProp } data={this.state.players} condensed hover bordered responsive version='4'>
                    {headers}
            </BootstrapTable>
            </Fade>
            </div>
        )
    }
}

export default StatTable;