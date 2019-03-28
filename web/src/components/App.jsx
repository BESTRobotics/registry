import React from "react";
import {
  HashRouter as Router,
  Switch,
  Route,
  Redirect
} from "react-router-dom";
import { connect } from "react-redux";
import { Header } from "semantic-ui-react";
import Login from "./rootPages/Login";
import Homepage from "./rootPages/Homepage";
import Register from "./rootPages/Register";
import Topbar from "./Topbar";
import Hubs from "./itemViews/Hubs";
import Teams from "./itemViews/Teams";
import HubsTables from "./itemTables/Hubs";
import TeamsTables from "./itemTables/Teams";
import SeasonsTables from "./itemTables/Seasons";
import UsersTables from "./itemTables/Users";
import EventsTables from "./itemTables/Events";
import BrcHub from "./itemViews/BrcHub";

const App = ({ token, superAdmin }) => {
  return (
    <Router
    // basename={process.env.PUBLIC_URL}
    >
      {token ? (
        <section className="root">
          <Topbar />
          <Header textAlign="center" as="h1">
            BEST Registry
          </Header>
          <Switch>
            <Redirect exact path="/" to="/hubs" />
            <Redirect path="/login" to="/hubs" />
            <Redirect path="/register" to="/hubs" />
            <Route exact path="/" component={Homepage} />
            <Route exact path={["/hubs", "/hubs/:id"]} component={Hubs} />
            <Route path="/teams" component={Teams} />
            <Route path="/hub/:id/brc" component={BrcHub} />
            {superAdmin && (
              <>
                <Route path="/admin/hubs" component={HubsTables} />
                <Route path="/admin/teams" component={TeamsTables} />
                <Route path="/admin/seasons" component={SeasonsTables} />
                <Route path="/admin/users" component={UsersTables} />
                <Route path="/admin/events" component={EventsTables} />
              </>
            )}
            <Route default render={() => <div>No route at path.</div>} />
          </Switch>
        </section>
      ) : (
        <Switch>
          <Redirect exact path="/" to="/login" />
          <Route path="/login" component={Login} />
          <Route path="/register" component={Register} />
          <Route default component={Login} />
        </Switch>
      )}
    </Router>
  );
};

const mapStateToProps = ({ loginReducer }) => ({
  token: loginReducer.token,
  superAdmin: loginReducer.superAdmin,
  hubs: loginReducer.hubs,
  teams: loginReducer.teams
});

const mapDispatchToProps = dispatch => ({});
// ...
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App);
