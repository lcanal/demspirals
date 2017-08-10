import React, {Component } from 'react';
import {Table,Button,Modal} from 'react-bootstrap';
import { Bar } from 'react-chartjs-2';

import "../css/Modals.css";

class PointCompModal extends Component {
  render() {
    var playas = [];
    var headers = [];
    
    //Build headers as we get them from the api
    for (var idx in this.props.headers) {
      if (this.props.headers.hasOwnProperty(idx)) {
        var header = this.props.headers[idx]
        headers.push(<th key={header}>{header}</th>)
      }
    }


    var options = {
      responsive: false,
      maintainAspectRatio: false,
      stacked: true,
      spanGaps: true,
      lineTension: 0.1,
       scales: {
            yAxes: [{
                ticks: {
                    suggestedMax: 350
                }
            }]
        }
    }

    //Main logic of adding player data to display
    for (var id in this.props.players) {
      if (this.props.players.hasOwnProperty(id)) {
        var player = this.props.players[id];
        var playerData = [];

        //Add data to display to chart. Match stat by the header type.
        for (var key in player) {
          if (player.hasOwnProperty(key)) {
            //Match stat with header
            for (var h in this.props.headers) {
              if (this.props.headers.hasOwnProperty(h)) {
                //Negative numbers OK, just don't want to fill chart with 0 bars.
                if (key === this.props.headers[h] ){
                  var stats = [];
                  stats.push(player[key])
                  //Construct datapoint for use in charts
                  var datapoint = {
                    label: this.props.headers[h],
                    data: stats,
                    backgroundColor: headerColorMap(this.props.headers[h])
                  }
                  playerData.push(datapoint)
                } // endif
              }
            }//end headerfor
          }
        }//end playerfor

        //Consturct data for particular chart
        var data = {
          labels: [ "Total Points: "+player.totalfantasypoints.toFixed(2) ],
          datasets: playerData
        }

        if (player.picurl.length <= 0){
          player.picurl = process.env.PUBLIC_URL + "/no-image.png"
        }
        var nflprofileURL = "http://www.nfl.com/player/pid/"+player.nflid+"/profile"
        var nameField = <a href={nflprofileURL} target="_">{player.name}</a>

        playas.push(
          <tr key={player.id}>
            <td className="modal-player">{nameField} <br /><img src={player.picurl} alt=" " /></td>
            <td><Bar data={data} options={options} height={300} width={1000}/></td>
          </tr>
          )
      }
    }

    return (
      <Modal {...this.props} bsSize="large" dialogClassName="custom-modal" aria-labelledby="contained-modal-title-lg">
        <Modal.Header  closeButton>
          <Modal.Title id="contained-modal-title-lg">Player Point Composition</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Table>
            <thead><tr>
              <th><strong>Player</strong></th>
              <th><strong>Point Composition</strong></th>
            </tr></thead>
            <tbody>{playas}</tbody>
          </Table>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={this.props.onHide}>Close</Button>
        </Modal.Footer>
      </Modal>
    );
  }
}

//Generate consistent colors based on the header type.
function headerColorMap(headerString){
  switch (headerString) {
    //Rushing stats
    case "RY10":
      return 'rgba(191, 191, 74, 0.2)'
    case "RTD":
      return 'rgba(74, 191, 74, 0.2)'
    default:
      return 'rgba(78, 192, 192, 0.2)'
  }
}

export default PointCompModal;