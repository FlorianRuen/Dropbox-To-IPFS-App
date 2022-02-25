import React, { Component } from 'react';
import { StyleSheet, Text, View, Linking, Button, Platform } from 'react-native';
import uuidV4 from 'uuid/v4';
import shittyQs from 'shitty-qs';

const DROPBOX_APP_KEY = 'c3hbpngaqu240bf';
export default class DropboxOAuthSample extends Component {

  constructor(props) {
    super(props);
    this.state = {
      isDropboxInit: false,
      apiToken: '',
      verification: uuidV4(),
    }
  }

  componentDidMount() {
    Linking.addEventListener('url', (event) => this.handleLinkingUrl(event));
  }

  componentWillUnmount() {
    Linking.removeEventListener('url', (event) => this.handleLinkingUrl(event));
  }

  handleLinkingUrl(event) {
    var [, query_string] = event.url.match(/(.*)/);
    var query = shittyQs(query_string);
    if (this.state.verification === query.state) {
      this.setState({isDropboxInit:true,apiToken:query.access_token})
    } else {
      alert("Verification non égale");
    }
  }

  loginWithDropbox() {
    const redirect_uri = Platform.OS === 'ios' ? 'dropboxoauthsample://open' : 'https://www.dropboxoauthsample.com/open';
    const url = `https://www.dropbox.com/oauth2/authorize?response_type=token&client_id=${DROPBOX_APP_KEY}&redirect_uri=${redirect_uri}&state=${this.state.verification}`;
    Linking.openURL(url).catch(err => console.error('An error occurred', err));
  }

  render() {
    const instruction = this.state.isDropboxInit ? 'Dropbox API token : ' + this.state.apiToken : 'Vous ne vous êtes pas encore connecté';
    return (
      <View style={styles.container}>
        <Button title='Se connecter avec Dropbox' onPress={() => this.loginWithDropbox()} />
        <Text style={styles.instructions}>
          {instruction}
        </Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#F5FCFF',
  },
  instructions: {
    marginTop: 32,
    textAlign: 'center',
    color: '#333333',
    marginBottom: 5,
  },
});