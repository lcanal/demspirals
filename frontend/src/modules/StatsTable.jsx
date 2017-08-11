import React, { Component } from 'react';
import { Fade,ProgressBar,Button,DropdownButton,MenuItem } from 'react-bootstrap';
import { BootstrapTable, TableHeaderColumn } from 'react-bootstrap-table';
import PointCompModal from "../components/PointCompModal.jsx";
import '../css/react-bootstrap-table-all.min.css';
import "../css/StatsTable.css";

var playersToPointComp = {};

const selectRowProp = {
  mode: 'checkbox',
  hideSelectColumn: true,
  clickToSelect: true, // enable click to select
  onSelect: onRowSelect,
  showOnlySelected: true,
  bgColor: 'rgba(51, 122, 183, 0.48)',
  className: 'stat-player-selected'
};


class StatTable extends Component {
    state = {
        players: [],
        playerHeaders: [],
        showTable: false,
        loadState: 0.0,
        numLimit: 0,
        lgShow: false
    }
    
    tableOptions = {
      sortIndicator: true,
      clearSearch: true
    };

    pointFormatter(cell, row) { 
        return `${cell.toFixed(2)}`;
    }

    setStateAsync(state) {
        return new Promise((resolve) => {
            this.setState(state, resolve)
        });
    }

    
    async componentDidMount(){
        await this.grabPlayerData(this.props.position,"stats")    //Default stat set is stats
    }

    componentWillUnmount(){
        //Don't want to carry over other players from other tabs.
        playersToPointComp = {};
    }

    async recalcHeaders(eventKey,event){
        this.setState({ showTable: false,numLimit: 0,loadState: 0})
        playersToPointComp = {};
        await this.grabPlayerData(this.props.position,eventKey);
    }

    //Main data gathering function
    async grabPlayerData(position,statSet){
        return new Promise(async (resolve) => {
            const response  = await fetch(this.props.apiURL);
            const json      = await response.json();

            await this.setStateAsync({"numLimit" : json.playerdata.length}); //causes setstate error.

            var playerHeaders = []
            var players = []

            var headerset = false;
            var position  = this.props.position
            for (var index = 0; index < json.playerdata.length; index++) {
                var player = json.playerdata[index]
                player.name = player.firstname + " " + player.lastname

                var playerstatsresponse =  await fetch("/api/player/"+position+"/"+player.id)
                var stats = await playerstatsresponse.json()

                //Set once for headers.
                if (!headerset){
                for (var hidx in stats) {
                    if (stats.hasOwnProperty(hidx)) {
                        if (statSet === "stats"){
                                playerHeaders.push(stats[hidx].name)
                        }else if (stats[hidx].leaguename.length > 0) {
                                playerHeaders.push(stats[hidx].leaguename)
                        }
                    }
                }
                headerset = true;
                }

                //Actually grab data in headers
                for (var pidx in stats) {
                    var statkey
                    if (stats.hasOwnProperty(pidx)) {
                        if (statSet === "stats"){
                                statkey = stats[pidx].name
                                player[statkey] = parseFloat(stats[pidx].statnum)
                        }else if (stats[pidx].leaguename.length > 0) {
                                statkey = stats[pidx].leaguename
                                player[statkey] = parseFloat(stats[pidx].value)
                        }
                    }
                }

                //Fix totalpoint. HARD assumption key exists
                player.totalfantasypoints = parseFloat(player.totalfantasypoints)
                
                this.setStateAsync({"loadState": index})
                players.push(player)
            }// endFor

                this.setStateAsync({
                "players" : players,
                "playerHeaders" : playerHeaders,
                "showTable": true
            }); //setAsync
            return resolve;
        }); //Promise
    }


    render(){
        var headers = []
        let lgClose = () => this.setState({ lgShow: false });
        
        //Build some headers manually
        headers.push(<TableHeaderColumn key="id" isKey={true} dataField="id" hidden={true}>#</TableHeaderColumn>)
        headers.push(<TableHeaderColumn key="name" dataField="name">Player</TableHeaderColumn>)
       
        headers.push(<TableHeaderColumn key="position" dataField="position">Position</TableHeaderColumn>)
        headers.push(<TableHeaderColumn key="teamname" dataField="teamname">Team</TableHeaderColumn>)
        //Build headers as we get them from the api
        this.state.playerHeaders.forEach(function(header) {
            headers.push(<TableHeaderColumn key={header} dataFormat={ this.pointFormatter } dataField={header} dataSort caretRender={getCaret}>{header}</TableHeaderColumn>)
        }, this);
        headers.push(<TableHeaderColumn key="totalfantasypoints" dataFormat={ this.pointFormatter } dataField="totalfantasypoints" dataSort caretRender={getCaret}>Total Points</TableHeaderColumn>)
        
        return ( 
            <div>
            <Fade in={!this.state.showTable} unmountOnExit={true} >
                <ProgressBar className="stats-table-load-status" now={this.state.loadState} max={this.state.numLimit} />
            </Fade>
            <Fade in={this.state.showTable} transitionAppear={true} >
                <div>

                <DropdownButton bsSize="small" bsStyle="primary" title="Stats" key="stats" pullRight id="stats-dropdown"> 
                    <MenuItem eventKey="stats" active onSelect={this.recalcHeaders.bind(this)}>Player Stats</MenuItem>
                    <MenuItem eventKey="espn" onSelect={this.recalcHeaders.bind(this)}>ESPN Point Value</MenuItem>
                </DropdownButton>

                <Button id="show-modal-button" bsStyle="primary" bsSize="small" disabled={this.state.lgShow} onClick={()=>this.setState({ lgShow: true })} className="show-modal-button" >
                    Show Point Composition
                </Button>

                <PointCompModal show={this.state.lgShow} onHide={lgClose} players={playersToPointComp} headers={this.state.playerHeaders} />
                <BootstrapTable selectRow={ selectRowProp } 
                                data={this.state.players} 
                                options={this.tableOptions}
                                multiColumnSort={ 2 }
                                pagination search condensed bordered hover responsive version='4'>
                    {headers}
                </BootstrapTable>
                </div>
            </Fade>
            </div>
        )
    }
}


//Helper functions
function getCaret(direction) {
    if (direction === 'asc') {
        return (
        <span><strong> ^</strong></span>
        );
    }
    if (direction === 'desc') {
        return (
        <span><strong> v</strong></span>
        );
    }
  return (
    <span></span>
  );
}

function onRowSelect(row, isSelected, e) {
    if (isSelected) {
        playersToPointComp[row["id"]] = row
    }else{
        delete playersToPointComp[row["id"]]
    }
}

export default StatTable;