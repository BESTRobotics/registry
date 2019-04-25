import React, { useEffect } from "react";
import { connect } from "react-redux";
import { getBrcTeam, registerBrcTeam } from "../../redux/teams/reducer";
import {
  Button,
  Icon,
  Header,
  Card,
  Grid,
  Divider,
  List
} from "semantic-ui-react";
import FakeItemGroup from "./FakeItemGroup";

const BrcDescription = ({ brcTeam, register, id, season }) => {
  console.log(brcTeam.brcTeam);
  return brcTeam.brcTeam ? (
    <Card fluid>
      <Card.Content>
        <Card.Header>{`${brcTeam.brcTeam.Team.StaticName}  |  ${
          brcTeam.brcTeam.Season.Name
        }`}</Card.Header>
        <Card.Meta>
          {brcTeam.brcTeam && brcTeam.brcTeam.BRIApproved ? (
            "Registration Approved"
          ) : (
            <>
              <Icon name="warning" /> Registration not yet approved
            </>
          )}
        </Card.Meta>
        <Card.Description>
          <Divider horizontal>
            <Header as="h4">
              <Icon name="calendar" />
              Events
            </Header>
          </Divider>
          <List divided verticalAlign="middle">
            {brcTeam.brcTeam.events
              ? brcTeam.brcTeam.events.map(e => (
                  <List.Item>
                    <List.Content>
                      <List.Header as="a">Event</List.Header>
                    </List.Content>
                  </List.Item>
                ))
              : "No events defined"}
          </List>
          <Divider horizontal>
            <Header as="h4">
              <Icon name="users" />
              Students
            </Header>
          </Divider>
          <List divided verticalAlign="middle">
            {brcTeam.brcTeam.roster ? (
              brcTeam.brcTeam.roster.map(s => (
                <List.Item>
                  <List.Content>
                    <List.Header>
                      {s.Firstname} {s.LastName}
                    </List.Header>
                  </List.Content>
                </List.Item>
              ))
            ) : (
              <FakeItemGroup rows={3} />
            )}
          </List>
        </Card.Description>
      </Card.Content>
    </Card>
  ) : (
    <Button onClick={() => register(id, season)}>Register</Button>
  );
};

const BrcTeam = ({
  allBrcTeams,
  getBrcTeam,
  registerBrc,
  match: {
    params: { id, season }
  }
}) => {
  useEffect(() => {
    (allBrcTeams && allBrcTeams[id]) || getBrcTeam(id);
  }, []);

  return (
    <Grid columns={2} centered>
      <Grid.Row>
        <Grid.Column>
          {allBrcTeams && allBrcTeams[id] ? (
            <BrcDescription
              brcTeam={allBrcTeams[id].find(s => String(s.ID) === season)}
              id={id}
              season={season}
              register={() => registerBrc(id)}
            />
          ) : (
            "Loading Team Info ..."
          )}
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = ({ teamsReducer }) => ({
  allBrcTeams: teamsReducer.allBrcTeams
});

const mapDispatchToProps = {
  getBrcTeam: id => getBrcTeam.request(id),
  registerBrc: (id, season) => registerBrcTeam.request(id, season)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(BrcTeam);
