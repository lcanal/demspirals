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

    for (var id in this.props.players) {
      if (this.props.players.hasOwnProperty(id)) {
        var player = this.props.players[id]
        if (player.picurl.length <= 0){
          player.picurl = process.env.PUBLIC_URL + "/no-image.png"
        }
        playas.push(
          <tr key={player.id}>
            <td className="modal-player">{player.name}
              <br /><img src={player.picurl} alt=" " /></td>
            <td className="modal-data">{player.totalfantasypoints}</td>
          </tr>
          )
      }
    }

    var options = {
      responsive: true,
      maintainAspectRatio: false,
      stacked: true,
      spanGaps: true,
      lineTension: 0.1,
      scales: {
            yAxes: [{
                stacked: true
            }]
        }
    }

    var data = {
      labels: ["1","2","3","4"],
      dataSets: [
        {
          label: "Age",
          data: [10,3,5,10],
          backgroundColor: generateColors([10,3,5,10],'rgba(75, 192, 192, 0.2)')
        },
        {
        label: "Age2",
          data: [12,9,7,11],
          backgroundColor: generateColors([10,3,5,10],'rgba(255, 99, 132, 0.2)')
        }
      ]
    }// data

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
            <tbody></tbody>
          </Table>
          <Bar data={data} options={options} height={200} />
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={this.props.onHide}>Close</Button>
        </Modal.Footer>
      </Modal>
    );
  }
}

function generateColors(dataArray,colorString){
  var colorArray = [colorString];

  if(typeof dataArray === "undefined"){
    return colorArray.push(colorString);
  }

  for (var i = 0; i < dataArray.length; i++) {
    colorArray.push(colorString);
  }

  return colorArray;
}

export default PointCompModal;