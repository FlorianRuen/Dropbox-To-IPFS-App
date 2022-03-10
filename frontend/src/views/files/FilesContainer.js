import React, { Component } from 'react';
import { Container, Row, Col, Button, Alert } from 'reactstrap';

import { CheckLoginAccountId, GetUserDetails } from '../../shared/services/users';
import { GetFilesMigratedForUser } from '../../shared/services/files';

import logoMin from '../../images/logo_min.png'

export default class FilesListContainer extends Component {

  constructor(props) {
    super(props);

    this.state = {
      isLoggedIn: false,
      user: null,
      files: null,
      accountId: null,
      isLoading: true,
      isError: false,
      errMessage: ''
    };
  }

  componentDidMount = () => {

    // Check localStorage if user is logged in
    const userAccountId = localStorage.getItem('dropboxToIPFSAccountId');
    
    if (userAccountId != null) {
      this.loadUser();
      this.loadFiles();
      this.setState({ isLoggedIn: true, isLoading: false });
    }
  }

  handleInputValueChange = event => {
    this.setState({ accountId: event.target.value })
  }

  handleDisconnect = () => {
    localStorage.removeItem('dropboxToIPFSAccountId');
    this.setState({ isLoggedIn: false });
  }

  handleLogin = async () => {

    const accountIdToCheck = { 
      value: this.state.accountId 
    };

    try {
      const response = await CheckLoginAccountId(accountIdToCheck);

      if (response.data) {

        // Add the token to the local storage
        localStorage.setItem("dropboxToIPFSAccountId", this.state.accountId)
        
        // Load files for the user
        this.loadUser();
        this.loadFiles();
        this.setState({ isLoggedIn: true, isLoading: false });

      } else {
        this.setState({ isError: true, errMessage: 'No user found with this account id' });
      }

    } catch (error) {
      this.setState({ isError: true, errMessage: 'No user found with this account id' });
      console.log(error);
    }
  }

  loadUser = async () => {

    const userAccountId = { 
      value: localStorage.getItem('dropboxToIPFSAccountId') 
    };

    try {
      const response = await GetUserDetails(userAccountId);

      if (response.data) {
        this.setState({ user: response.data });
      }

    } catch (error) {
      console.log(error);
    }

  }

  loadFiles = async () => {

    this.setState({ isLoading: true });

    const userAccountId = { 
      value: localStorage.getItem('dropboxToIPFSAccountId') 
    };

    try {
      const response = await GetFilesMigratedForUser(userAccountId);

      if (response.data) {
        this.setState({ files: response.data, isLoading: false });
      }

    } catch (error) {
      console.log(error);
    }

  }

  render() {
    const { isLoading, isError, errMessage, isLoggedIn, user, files } = this.state;

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

                {isLoggedIn && user ? (

                  <>
                    <h2 className="mb-4">Welcome back {user.firstName} !</h2>

                    {isLoading ? (

                      <Alert color="warning">
                        Loading files ...
                      </Alert>

                    ) : (

                      <>
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

                                {files.map((file) => (
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
                          <Alert color="danger">
                            No migrated files has been found on your account
                          </Alert>
                        )}

                        <Row className="mt-4">
                          <Col md="6">
                            <Button onClick={this.handleDisconnect} color="danger" size="lg">
                              <span className="as--light">
                                Log out
                              </span>
                            </Button>
                          </Col>

                          <Col md="6">
                            <Button onClick={this.loadFiles} color="primary" size="lg">
                              <span className="as--light">
                                Refresh file list
                              </span>
                            </Button>
                          </Col>
                        </Row>
                      </>
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

                    {/* Display error message for login form */}
                    {isError && (
                      <Row className="mt-4">
                        <Col md="12">
                          <Alert color="danger">
                            {errMessage}
                          </Alert>
                        </Col>
                      </Row>
                    )}

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