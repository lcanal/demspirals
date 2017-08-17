import React, { Component } from 'react';
import { Line } from 'react-chartjs-2';

class GraphPlayers extends Component {
    state = {
        players: [],
        playerHeaders: [],
        showTable: false,
        startHidden: true,
        loadState: 0.0,
        numLimit: 0,
    }

    setStateAsync(state) {
        return new Promise((resolve) => {
            this.setState(state, resolve)
        });
    }

    async componentDidMount(){
        await this.grabPlayerData(this.props.position,"stats");
    }
    
    //Main data gathering function
    async grabPlayerData(position,statSet){
        return new Promise(async (resolve) => {
            var apiURL = "/api/topplayers/" + this.props.position
            const response  = await fetch(apiURL);
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
                                if (statkey === "Games"){              //Special exception for games since they're not computed as having a "value"
                                    player[statkey] = parseFloat(stats[pidx].statnum)
                                }else{
                                    player[statkey] = parseFloat(stats[pidx].value)
                                }
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
                "showTable": true,
            }); //setAsync
            return resolve;
        }); //Promise
    }



    render(){
        var playerData = [];
        var headers = [];
        var options = {
            responsive: true,
            maintainAspectRatio: true,
            legend: {
                position: 'bottom',
                labels: {
                    usePointStyle: true
                }
            }
        }

        //Build headers as we get them from the api
        for (var idx in this.state.playerHeaders) {
            if (this.state.playerHeaders.hasOwnProperty(idx)) {
                var header = this.state.playerHeaders[idx]
                headers.push(header)
            }
        }
        
        //Main logic of adding player data to display ///////////////
        for (var id in this.state.players) {
            if (this.state.players.hasOwnProperty(id)) {
                var player = this.state.players[id];
                var stats = [];
                //Add data to display to chart. Match stat by the header type.
                for (var key in player) {
                    
                    if (player.hasOwnProperty(key)) {
                        //Match stat with header
                        for (var h in this.state.playerHeaders) {
                            if (this.state.playerHeaders.hasOwnProperty(h)) {
                                //Negative numbers OK. PassYards greatly skews graph..
                                if (key === this.state.playerHeaders[h] && this.state.playerHeaders[h] !== "PassYards" ){
                                    stats.push(player[key])
                                } // endif
                            }
                        }//end headerfor
                    }
                }//end playerfor DATA
                //Construct datapoint for use in charts
                var datapoint = {
                    label: player.name,
                    data: stats,
                    backgroundColor: randomColorGenerator(0.2),
                    hidden: this.state.startHidden
                }
                playerData.push(datapoint)
            }
        }//endofor/////////////
        
        var data = {
            labels: headers,
            datasets: playerData
        }

        return(
            <div>
                <Line data={data} options={options} />
            </div>
        )
    }
}

//Generate consistent colors based on the header type.
function randomColorGenerator(alphaLevel){
    //Generate random colors between 1-255
    var r = Math.floor(Math.random() * 255)
    var b = Math.floor(Math.random() * 255)
    var g = Math.floor(Math.random() * 255)
    var colorString = "rgba("+r+","+b+","+g+","+alphaLevel+")"
      return colorString
}

export default GraphPlayers;