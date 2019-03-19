import React from "react";
import Login from "./rootPages/Login";
import Homepage from "./rootPages/Homepage";
import NewUser from "./rootPages/NewUser";
import Register from "./rootPages/Register";
import Super from "./rootPages/Super";
import Topbar from "./Topbar";
import {
  HashRouter as Router,
  Switch,
  Route,
  Redirect
} from "react-router-dom";
import { connect } from "react-redux";

const App = ({ token, superAdmin, hubs, teams }) => {
  return (
    <Router
    // basename={process.env.PUBLIC_URL}
    >
      {token ? (
        <section className="root">
          <Topbar />
          {superAdmin ? (
            <Super />
          ) : (hubs && hubs.length) || (teams && teams.length) ? (
            <Homepage />
          ) : (
            <NewUser />
          )}
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
