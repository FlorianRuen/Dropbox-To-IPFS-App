import React from "react";
import { Switch, Route } from "react-router-dom";

import { DropboxAuth, DropboxCallback } from "./views/dropbox-auth";

const App = () => {
  return (
    <div className="app-container d-flex flex-column">

      <div>
        <Switch>
          <Route exact path="/" component={DropboxAuth} />
          <Route path="/callback" component={DropboxCallback} />
        </Switch>
      </div>
    </div>
  );
};

export default App;