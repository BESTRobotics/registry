import React, { useState, useEffect } from "react";
import Hubs from "./Hubs";
import Teams from "./Teams";
import Schools from "./Schools";
import Seasons from "./Seasons";
import Users from "./Users";
import Topbar from "./Topbar";
import Login from "./Login";
import Register from "./Register";
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
    <Router basename={process.env.PUBLIC_URL}>
      {token ? (
        <section className="root">
          <Topbar logout={logout} />
          <Switch>
            <Redirect exact path="/" to="/hubs" />
            <Redirect path="/login" to="/hubs" />
            <Redirect path="/register" to="/hubs" />
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
        <Switch>
          <Redirect exact path="/" to="/login" />
          <Route
            path="/login"
            render={p => <Login {...p} setToken={setToken} />}
          />
          <Route path="/register" render={p => <Register {...p} />} />
          <Route default render={p => <Login {...p} setToken={setToken} />} />
        </Switch>
      )}
    </Router>
  );
};

export default App;
