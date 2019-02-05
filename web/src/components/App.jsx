import React, { Component } from "react";
import { Container } from "semantic-ui-react";
import Hubs from "./Hubs";
import Topbar from "./Topbar";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

class App extends Component {
  render() {
    return (
      <Router>
        <section>
          <Topbar />
          <Switch>
            <Route exact path="/" redirect="/Hubs" />
            <Route path="/Hubs" component={Hubs} />
          </Switch>
        </section>
      </Router>
    );
  }
}

export default App;
