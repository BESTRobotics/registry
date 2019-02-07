import React, { Component } from "react";
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

class App extends Component {
  render() {
    return (
      <Router>
        <section>
          <Topbar />
          <Switch>
            <Redirect exact path="/" to="/hubs" />
            <Route path="/hubs" component={Hubs} />
            <Route path="/schools" component={Schools} />
            <Route path="/teams" component={Teams} />
            <Route path="/users" component={Users} />
            <Route default render={() => <div>No route at path.</div>} />
          </Switch>
        </section>
      </Router>
    );
  }
}

export default App;
