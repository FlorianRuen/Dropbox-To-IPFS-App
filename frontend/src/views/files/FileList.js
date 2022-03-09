import React, { Component } from 'react';
import { Row, Col } from 'reactstrap';

export default class FileList extends Component {

  render() {
    return (
      <Row className="mt-4">
        <Col md="12">
          <table className="table">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Filename</th>
                <th scope="col">CID</th>
                <th scope="col">File size</th>
              </tr>
            </thead>

            <tbody>
              <tr>
                <th scope="row">1</th>
                <td>Mark</td>
                <td>Otto</td>
                <td>@mdo</td>
              </tr>
            </tbody>
          </table>
        </Col>
      </Row>
    );
  }
}