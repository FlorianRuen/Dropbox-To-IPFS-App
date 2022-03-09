import React, { Component } from 'react';
import { Container, Row, Col, Button } from 'reactstrap';

import { CheckLoginToken } from '../../shared/services/login';
import { GetFilesMigratedForUser } from '../../shared/services/files';

import logoMin from '../../images/logo_min.png'
import FileList from './FileList';

export default class FilesListContainer extends Component {

  constructor(props) {
    super(props);

    this.state = {
      isLoggedIn: false,
      user: null,
      files: null,
      token: null
    };
  }

  handleInputValueChange = (event) => {
    this.setState({ token: event.target.value })
  }

  handleLogin = async () => {
    const tokenToCheck = { value: this.state.token };

    try {
      const response = await CheckLoginToken(tokenToCheck);

      if (response.data) {
        this.loadFiles();
        this.setState({ user: response.data, isLoggedIn: true });
      }

    } catch (error) {
      console.log(error);
    }
  }

  loadFiles = async () => {

    const tokenToCheck = { value: this.state.token };

    try {
      const response = await GetFilesMigratedForUser(tokenToCheck);

      if (response.data) {
        console.log(response.data);
        this.setState({ files: response.data });
      }

    } catch (error) {
      console.log(error);
    }

  }

  render() {
    const { isLoggedIn, user, files } = this.state;

    return (
      <Container>
        <Row>
          <Col md="4">
            <img src={logoMin} className="main-logo" height="300px" alt="logo" />
          </Col>

          <Col md="8">
            <div className="jumbotron mt-5">
              <div className="text-center">

                {isLoggedIn && user ? (
                  
                  <>
                    <h2 className="mt-4">Welcome {user.token_type}</h2>
                    
                    {files && files.length > 0 ? (
                      <>
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

                            {files.map((file, indexFile) => (
                              <tr>
                                <th scope="row">{file.estuaryId}</th>
                                <td>{file.name}</td>
                                <td>{file.cid}</td>
                                <td>{file.size}</td>
                              </tr>
                            ))}

                          </tbody>
                        </table>
                      </>
                    
                    ) : (
                      <>NO FILES</>
                    )}

                  </>

                ) : (

                  <>
                    <h2 className="mt-4">Login to check files status</h2>

                    <Row className="mt-4">
                      <Col md="12">
                        <input 
                          type="text" className="form-control" 
                          placeholder="Enter your token" 
                          onChange={this.handleInputValueChange}
                        />
                      </Col>
                    </Row>

                    <Row className="mt-4">
                      <Col md="12">
                        <Button onClick={this.handleLogin} color="primary" size="lg">
                          <span className="as--light">
                            Login
                          </span>
                        </Button>
                      </Col>
                    </Row>
                  </>
                )}

              </div>
            </div>
          </Col>
        </Row>
      </Container>
    );
  }
}