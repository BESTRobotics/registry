import React, { useState, useEffect } from "react";
import Hubs from "./Hubs";
import Teams from "./Teams";
import Schools from "./Schools";
import Seasons from "./Seasons";
import Users from "./Users";
import Topbar from "./Topbar";
import Login from "./Login";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect
} from "react-router-dom";

const App = () => {
  const [token, setToken] = useState(
    window.localStorage.getItem("token") || null
  );

  useEffect(() => {
    if (token) {
      window.localStorage.setItem("token", token);
    } else {
      window.localStorage.removeItem("token");
    }
  }, [token]);

  const logout = () => {
    setToken(null);
  };
  return (
    <Router>
      {token ? (
        <section className="root">
          <Topbar logout={logout} />
          <Switch>
            <Redirect exact path="/" to="/hubs" />
            <Route path="/hubs" render={p => <Hubs {...p} token={token} />} />
            <Route
              path="/schools"
              render={p => <Schools {...p} token={token} />}
            />
            <Route path="/teams" render={p => <Teams {...p} token={token} />} />
            <Route
              path="/seasons"
              render={p => <Seasons {...p} token={token} />}
            />
            <Route path="/users" render={p => <Users {...p} token={token} />} />
            <Route default render={() => <div>No route at path.</div>} />
          </Switch>
        </section>
      ) : (
        <Login setToken={setToken} />
      )}
    </Router>
  );
};

export default App;

// <Table.Cell>{item.Name}</Table.Cell>
// <Table.Cell>{hub.Location}</Table.Cell>
// <Table.Cell>
//   {hub.Director.FirstName} {hub.Director.LastName}
// </Table.Cell>
// <Table.Cell>{hub.Description}</Table.Cell>
