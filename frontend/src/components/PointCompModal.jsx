import React, {Component } from 'react';
import {Table,Button,Modal} from 'react-bootstrap';
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
        playas.push(
          <tr key={player.id}>
            <td className="modal-player">{player.name}<br /><img src={player.picurl} alt={player.name}/></td>
            <td className="modal-data">{player.totalfantasypoints}</td>
          </tr>
          )
      }
    }
    return (
      <Modal {...this.props} bsSize="large" dialogClassName="custom-modal" aria-labelledby="contained-modal-title-lg">
        <Modal.Header  closeButton>
          <Modal.Title id="contained-modal-title-lg">Modal heading</Modal.Title>
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


export default PointCompModal;