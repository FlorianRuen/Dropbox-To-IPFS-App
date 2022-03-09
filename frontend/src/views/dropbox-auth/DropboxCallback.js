import React, { PureComponent } from "react";
import { Container, Row, Col, Button } from 'reactstrap';

import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleCheck , faTriangleExclamation } from "@fortawesome/free-solid-svg-icons";

import logoMin from '../../images/logo_min.png'

class DropboxCallback extends PureComponent {

  constructor(props) {
    super(props);

    this.state = {
      callbackUID: null,
      isError: false
    }
  }

  componentDidMount() {
    const callbackUID = this.props.match.params.token;

    if (callbackUID === null || callbackUID === undefined) {
      this.setState({ isError: true });
    } else {
      this.setState({ callbackUID });
    }
  }

  render() {
    const { isError, callbackUID } = this.state;
    
    return (
      <Container>
        <Row>

          <Col md="4">
            <img src={logoMin} className="main-logo" height="300px" alt="logo" />
          </Col>

          <Col md="8">
            <div className="jumbotron mt-5 text-center">

              {!isError ? (
                <>
                  <div className="text-center">
                    <FontAwesomeIcon icon={faCircleCheck} size="3x" style={{ color: 'green' }} />
                    <h2 className="mt-4">Integration with your Dropbox account completed</h2>
                  </div>

                  <Row className="mt-4">
                    <Col md="12">
                      <p className="mt-4">Use this token to log in to the app and check the status of your migrations</p>
                      <p className="mb-4">Keep it in safe place, you won't be able to see it again</p>
                      
                      <pre className="multiline text-center col-12 text-light bg-dark p-4">
                        {callbackUID}
                      </pre>
                    </Col> 
                  </Row>

                  <Row className="text-center mt-4">
                    <Col md="12">
                      <Button tag={Link} to="/files" color="secondary" size="lg">
                        <span className="as--light">
                          Check my files status
                        </span>
                      </Button>
                    </Col>
                  </Row>
                </>

              ) : (
                
                <>
                  <div className="text-center">
                    <FontAwesomeIcon icon={faTriangleExclamation} size="3x" style={{ color: 'orange' }} />
                    <h2 className="mt-4">Something wrong during authentification flow</h2>
                  </div>

                  <Row className="mt-4">
                    <Col md="12">
                      <div className="card">
                        <div className="card-body text-center">
                          It seems an error occured when validating your authentification to your Dropbox
                          account. Please try again or create an issue on our Github repository
                          <br /><br />
                          <a href="https://github.com/FlorianRuen/Dropbox-To-IPFS-App">https://github.com/FlorianRuen/Dropbox-To-IPFS-App</a>
                        </div>
                      </div>
                    </Col>
                  </Row>

                  <Row className="text-center mt-4">
                    <Col md="12">
                      <Button tag={Link} to="/" color="warning" size="lg">
                        <span className="as--light">
                          Try authentification again
                        </span>
                      </Button>
                    </Col>
                  </Row>
                </>

              )}

            </div>
          </Col>       
        </Row>

      </Container>
    );
  }
}

export default DropboxCallback;