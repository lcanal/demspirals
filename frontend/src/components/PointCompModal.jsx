import React, {Component } from 'react';
import {Table,Button,Modal} from 'react-bootstrap';
import { Bar } from 'react-chartjs-2';

import "../css/Modals.css";

class PointCompModal extends Component {
  state = {
    topNumber: 0,
  }

  calcTopNumber(){
    var topNumber = this.state.topNumber;
    //Find largest number to keep scaling consistent
      for (var id in this.props.players) {
          var player = this.props.players[id];
        if (this.props.players.hasOwnProperty(id)) {
          for (var key in player) {
            if (player.hasOwnProperty(key)) {
              for (var h in this.props.headers) {
                if (this.props.headers.hasOwnProperty(h)) {
                  if (key === this.props.headers[h] ){    
                  //Let's keep the highest number to keep all scales the same.
                    if(player[key] > topNumber){
                      topNumber = player[key]
                    }
                  }
                }
              }
            }
          }
        }
      }

      return topNumber;
  }
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
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          ticks: {
            suggestedMax: this.calcTopNumber(),
            type: 'logarithmic'
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
                //Negative numbers OK
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
        var nflprofileURL = "http://www.nfl.com/player/pid/"+player.nflid+"/careerstats"
        var nameField = <a href={nflprofileURL} target="_">{player.name}</a>

        playas.push(
          <tr key={player.id}>
            <td className="modal-player">{nameField} <br /><img src={player.picurl} alt=" " /></td>
            <td><div className="bar-div"><Bar data={data} options={options}/></div></td>
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
    case "RY10":                        //Rushing
      return 'rgba(191, 191, 74, 0.2)'
    case "RTD":
      return 'rgba(74, 191, 74, 0.2)'
    case "REC":                         //Receiving
      return 'rgba(42, 127, 247, 0.2)'
    case "REY10":
      return 'rgba(24, 153, 1, 0.2)'
    case "RETD":
      return 'rgba(196, 0, 189, 0.2)'
    case "FUML":                        //Fumble
      return 'rgba(247, 0, 0, 0.2)'
    case "PC":                          //Passing
      return 'rgba(51, 234, 118, 0.2)'
    case "PY20":
      return 'rgba(40, 70, 140, 0.2)'
    case "PTD":
      return 'rgba(188, 185 , 1, 0.2)'
    case "INT":
      return 'rgba(181, 42 , 193, 0.2)'
    case "SK":
      return 'rgba(142, 78 , 18, 0.2)'
    default:
      return 'rgba(78, 192, 192, 0.2)'
  }
}

export default PointCompModal;