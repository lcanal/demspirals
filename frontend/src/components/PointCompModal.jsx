import React, {Component } from 'react';
import {Button,Modal} from 'react-bootstrap';

class PointCompModal extends Component {
  render() {
    var playas = []
    for (var id in this.props.players) {
      if (this.props.players.hasOwnProperty(id)) {
        var player = this.props.players[id]
        playas.push(<p key={player.id}>Hello derr {player.name} </p>)
      }
    }
    return (
      <Modal {...this.props} bsSize="large" dialogClassName="custom-modal" aria-labelledby="contained-modal-title-lg">
        <Modal.Header  closeButton>
          <Modal.Title id="contained-modal-title-lg">Modal heading</Modal.Title>
        </Modal.Header>
        <Modal.Body>
         {playas}
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={this.props.onHide}>Close</Button>
        </Modal.Footer>
      </Modal>
    );
  }
}


export default PointCompModal;