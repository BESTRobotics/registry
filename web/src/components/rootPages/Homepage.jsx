import React from "react";
import { connect } from "react-redux";
import { Grid, Header } from "semantic-ui-react";
import { Switch, Redirect, Route } from "react-router-dom";
import Hubs from "../itemViews/Hubs";
import Teams from "../itemViews/Teams";

const mapStateToProps = ({ loginReducer }) => ({
  hubsLength: loginReducer.hubs ? loginReducer.hubs.length : 0,
  teamsLength: loginReducer.teams ? loginReducer.teams.length : 0
});

const InnerHomepage = connect(mapStateToProps)(
  ({ hubsLength, teamsLength }) => {
    return (
      <React.Fragment>
        {hubsLength ? <Hubs /> : null}
        {teamsLength ? <Teams /> : null}
      </React.Fragment>
    );
  }
);

const Homepage = ({ hubsLength, teamsLength, getMyHubs, getMyTeams }) => {
  return (
    <Grid centered columns={2}>
      <Grid.Row>
        <Header as="h1">BEST Registry</Header>
      </Grid.Row>

      <Switch>
        <Redirect path="/login" to="/" />
        <Redirect path="/register" to="/" />
        <Route exact path="/" component={InnerHomepage} />
        <Route path="/hubs" component={Hubs} />
        <Route path="/teams" component={Teams} />
        <Route default render={() => <div>No route at path.</div>} />
      </Switch>
    </Grid>
  );
};

export default Homepage;
