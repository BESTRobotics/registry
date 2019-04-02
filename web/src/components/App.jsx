import React from "react";
import {
  HashRouter as Router,
  Switch,
  Route,
  Redirect
} from "react-router-dom";
import { connect } from "react-redux";
import { Header, Message, Container } from "semantic-ui-react";
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

const App = ({ token, superAdmin, message }) => {
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
          {message && (
            <Container>
              <Message {...message} />
              <br />
            </Container>
          )}
          <Switch>
            <Redirect path="/login" to="/" />
            <Redirect path="/register" to="/" />
            <Route exact path="/" component={Homepage} />
            <Route exact path={["/hubs", "/hubs/:id"]} component={Hubs} />
            <Route path="/teams" component={Teams} />
            <Route path="/hubs/:id/brc/:season" component={BrcHub} />
            {/* These below guys need to be more reduxy */}
            {superAdmin && (
              <>
                <Route
                  path="/admin/hubs"
                  render={p => <HubsTables {...p} token={token} />}
                />
                <Route
                  path="/admin/teams"
                  render={p => <TeamsTables {...p} token={token} />}
                />
                <Route
                  path="/admin/seasons"
                  render={p => <SeasonsTables {...p} token={token} />}
                />
                <Route
                  path="/admin/users"
                  render={p => <UsersTables {...p} token={token} />}
                />
                <Route
                  path="/admin/events"
                  render={p => <EventsTables {...p} token={token} />}
                />
              </>
            )}
            <Route default render={() => <div>No route at path.</div>} />
          </Switch>
        </section>
      ) : (
        <Switch>
          <Route path="/login" component={Login} />
          <Route path="/register" component={Register} />
          <Redirect path="/" to="/login" />
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
  teams: loginReducer.teams,
  message: loginReducer.message
});

const mapDispatchToProps = dispatch => ({});
// ...
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App);
