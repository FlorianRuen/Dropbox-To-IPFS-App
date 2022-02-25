import React from "react";
import { Switch, Route } from "react-router-dom";

import { DropboxAuth } from "./views/dropbox-auth";
import "./scss/app.scss";

const App = () => {
  return (
    <div className="app-container d-flex flex-column">

      <div>
        <Switch>
          
          {/* Following routes are protected only by Auth0 */}
          <Route path="/" component={DropboxAuth} />
        </Switch>
      </div>
    </div>
  );
};

export default App;