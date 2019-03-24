import React from "react";
import Hubs from "../itemTables/Hubs";
import Teams from "../itemTables/Teams";
import Seasons from "../itemTables/Seasons";
import Users from "../itemTables/Users";
import Events from "../itemTables/Events";
import { Switch, Redirect, Route } from "react-router-dom";

const Super = () => (
  <Switch>
    <Redirect exact path="/" to="/hubs" />
    <Redirect path="/login" to="/hubs" />
    <Redirect path="/register" to="/hubs" />
    <Route path="/hubs" component={Hubs} />
    <Route path="/teams" component={Teams} />
    <Route path="/seasons" component={Seasons} />
    <Route path="/users" component={Users} />
    <Route path="/events" component={Events} />
    <Route default render={() => <div>No route at path.</div>} />
  </Switch>
);

export default Super;
