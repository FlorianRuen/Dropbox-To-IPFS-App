import React, { Component } from 'react';
import { Container, Row, Col } from 'reactstrap';
import uuidV4 from 'uuid/v4';

const DROPBOX_APP_KEY = 'c3hbpngaqu240bf';

export default class DropboxOAuthSample extends Component {

  constructor(props) {
    super(props);

    this.state = {
      verification: uuidV4(),
    }
  }

  render() {
    const redirect_uri = "http://localhost:3200/api/dropbox/oauth_callback"
    const auth_url = `https://www.dropbox.com/oauth2/authorize?response_type=code&client_id=${DROPBOX_APP_KEY}&redirect_uri=${redirect_uri}&state=${this.state.verification}`;

    return (
      <Container>
        <Row>
          <Col>
            <a href={auth_url} className="btn btn-primary btn-lg mt-4">Primary link</a>
          </Col>
        </Row>
      </Container>
    );
  }
}