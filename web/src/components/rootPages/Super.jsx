import React from "react";
import Hubs from "../itemLists/Hubs";
import Teams from "../itemLists/Teams";
import Schools from "../itemLists/Schools";
import Seasons from "../itemLists/Seasons";
import Users from "../itemLists/Users";
import Events from "../itemLists/Events";
import { Switch, Redirect, Route } from "react-router-dom";

const Super = () => (
  <Switch>
    <Redirect exact path="/" to="/hubs" />
    <Redirect path="/login" to="/hubs" />
    <Redirect path="/register" to="/hubs" />
    <Route path="/hubs" component={Hubs} />
    <Route path="/schools" component={Schools} />
    <Route path="/teams" component={Teams} />
    <Route path="/seasons" component={Seasons} />
    <Route path="/users" component={Users} />
    <Route path="/events" component={Events} />
    <Route default render={() => <div>No route at path.</div>} />
  </Switch>
);

export default Super;
