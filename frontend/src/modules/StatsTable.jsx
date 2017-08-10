import React, { Component } from 'react';
import { Fade,ProgressBar,Button } from 'react-bootstrap';
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
        numLimit: 0
    }
    
    options = {
      sortIndicator: true 
    };

    getInitialState() {
        return { lgShow: false };
    }


    setStateAsync(state) {
        return new Promise((resolve) => {
            this.setState(state, resolve)
        });
    }

    
    async componentDidMount(){
        const response  = await fetch(this.props.apiURL)
        const json      = await response.json()
        var playerHeaders = []
        var players = []

        var headerset = false;
        var position  = this.props.position

        await this.setStateAsync({"numLimit" : json.playerdata.length}); //causes setstate error.
        //this.setState({"numLimit" : json.playerdata.length}); //causes setstate error.

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
                var statkey
                   if (stats.hasOwnProperty(pidx)) {
                       if (stats[pidx].leaguename.length > 0){
                            statkey = stats[pidx].leaguename
                            player[statkey] = stats[pidx].value.toFixed(2)
                       }else{
                            statkey = stats[pidx].name
                            player[statkey] = stats[pidx].statnum.toFixed(2)
                       }
                   }
               }

            //Fix totalpoint. HARD assumption key exists
            player.totalfantasypoints = player.totalfantasypoints.toFixed(2)
            
            await this.setStateAsync({"loadState": index})
            players.push(player)
        }
       
        await this.setStateAsync({
            "players" : players,
            "playerHeaders" : playerHeaders,
            "showTable": true
        });
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
            headers.push(<TableHeaderColumn key={header} dataField={header} dataSort={true} caretRender={getCaret}>{header}</TableHeaderColumn>)
            
        }, this);
        headers.push(<TableHeaderColumn key="totalfantasypoints" dataField="totalfantasypoints" dataSort caretRender={getCaret}>Total Points</TableHeaderColumn>)
        
        
        return (
            <div>
            <Fade in={!this.state.showTable} unmountOnExit={true} >
                <ProgressBar className="stats-table-load-status" now={this.state.loadState} max={this.state.numLimit} />
            </Fade>
            <Fade in={this.state.showTable} transitionAppear={true} >
                <div>
                <Button bsStyle="primary" bsSize="small" onClick={()=>this.setState({ lgShow: true })} className="show-modal-button" >
                    Show Point Composition
                </Button>
                <PointCompModal show={this.state.lgShow} onHide={lgClose} players={playersToPointComp} />
                <BootstrapTable selectRow={ selectRowProp } 
                                data={this.state.players} 
                                options={this.options}
                                multiColumnSort={ 2 }
                                pagination condensed bordered hover responsive version='4'>
                    {headers}
                </BootstrapTable>
                </div>
            </Fade>
            </div>
        )
    }
}

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
    /*let rowStr = '';
    for (const prop in row) {
        rowStr += prop + ': "' + row[prop] + '"';
    }
    //console.log(e);
    //console.log(`is selected: ${isSelected} =======  ${rowStr}`);
    // console.log("ID I got: ", row["id"])*/
    //console.log("ID I got: ", row["id"])
    if (isSelected) {
        playersToPointComp[row["id"]] = row
    }else{
        delete playersToPointComp[row["id"]]
    }
    console.log("Player array id size..",Object.keys(playersToPointComp).length)
}

export default StatTable;