import React, { Component } from 'react';
import { Container, Row, Col, Button } from 'reactstrap';
import { Link } from "react-router-dom";
import { DROPBOX_APP_KEY, SERVER_API_URL } from '../../config/constants';

import uuidV4 from 'uuid/v4';
import logoMin from '../../images/logo_min.png'

export default class DropboxOAuthSample extends Component {

  constructor(props) {
    super(props);

    this.state = {
      verification: uuidV4(),
    }
  }

  render() {
    
    const { verification } = this.state;
    const redirect_uri = `${SERVER_API_URL}/dropbox/oauth_callback`;
    const auth_url = `https://www.dropbox.com/oauth2/authorize?response_type=code&client_id=${DROPBOX_APP_KEY}&redirect_uri=${redirect_uri}&state=${verification}`;

    return (
      <Container>
        <Row>
          <Col md="12" className="text-center">
            <img src={logoMin} className="mt-4" height="100px" alt="logo" />
          </Col>
        </Row>

        <Row>
          <Col md="12">
            <div className="jumbotron mt-5">
              <div className="text-center">
                <h2 className="mt-4">Welcome on Dropbox to IPFS application</h2>

                  <Row className="mt-4">
                    <Col md="4">
                      <div className="card">
                        <div className="card-header font-weight-bold">
                          1. Copy files to app folder
                        </div>

                        <div className="card-body">
                          Copy all the files you want to migrate in the folder 
                          Applications/Send-To-IPFS
                        </div>
                      </div>
                    </Col>

                    <Col md="4">
                      <div className="card">
                        <div className="card-header font-weight-bold">
                          2. Files will be detected
                        </div>

                        <div className="card-body">
                          Our application will identify all the files you add to send them over IPFS
                        </div>
                      </div>
                    </Col>    

                    <Col md="4">
                      <div className="card">
                        <div className="card-header font-weight-bold">
                          3. Check status
                        </div>

                        <div className="card-body">
                          Use the identifier below, to access the list and status of
                          each migrated file
                        </div>
                      </div>
                    </Col> 
                  </Row>

                  <Row className="mt-4">
                    <Col md="6">
                      <a href={auth_url} className="btn btn-primary btn-lg mt-4">Connect my Dropbox account</a>
                    </Col>

                    <Col md="6" className="mt-4">
                      <Button tag={Link} to="/files" color="secondary" size="lg">
                        <span className="as--light">
                          Check my files status
                        </span>
                      </Button>
                    </Col>
                  </Row>
              </div>
            </div>
          </Col>
        </Row>
      </Container>
    );
  }
}