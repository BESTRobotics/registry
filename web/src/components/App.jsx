import React, { useState } from "react";
import Hubs from "./Hubs";
import Teams from "./Teams";
import Schools from "./Schools";
import Users from "./Users";
import Topbar from "./Topbar";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect
} from "react-router-dom";
import { Message } from "semantic-ui-react";

const App = () => {
  const [message, setMessage] = useState(null);
  return (
    <Router>
      <section>
        <Topbar />
        {message ? <Message {...message} /> : null}
        <Switch>
          <Redirect exact path="/" to="/hubs" />
          <Route
            path="/hubs"
            render={p => <Hubs {...p} setMessage={setMessage} />}
          />
          <Route path="/schools" component={Schools} />
          <Route path="/teams" component={Teams} />
          <Route path="/users" component={Users} />
          <Route default render={() => <div>No route at path.</div>} />
        </Switch>
      </section>
    </Router>
  );
};

export default App;
