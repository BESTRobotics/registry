import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Card, Grid, Header, Icon, List, Item } from "semantic-ui-react";
import { Link } from "react-router-dom";
import { getMyHubs } from "../../redux/hubs/reducer";
import { getMyTeams } from "../../redux/teams/reducer";

const Homepage = ({
  hubs,
  teams,
  hubsLength,
  teamsLength,
  getMyHubs,
  getMyTeams
}) => {
  useEffect(() => {
    getMyHubs();
    getMyTeams();
  }, []);

  return (
    <Grid centered columns={2}>
      <Grid.Row>
        <Header as="h1">BEST Registry</Header>
      </Grid.Row>
      {hubsLength ? (
        <Grid.Row>
          <Grid.Column>
            <Card fluid color="orange">
              <Card.Content>
                <Card.Header as={Header} size="huge">
                  My Hubs
                </Card.Header>
                <Card.Description>
                  <Item.Group divided>
                    {hubs && hubs.length
                      ? hubs.map(h => (
                          <Item>
                            <Item.Content>
                              <Item.Header>{h.Name}</Item.Header>
                              <Item.Meta>{h.Description}</Item.Meta>
                              <Item.Description />
                            </Item.Content>
                          </Item>
                        ))
                      : "Loading"}
                  </Item.Group>
                </Card.Description>
              </Card.Content>
            </Card>
          </Grid.Column>
        </Grid.Row>
      ) : null}
      <Grid.Row>
        <Grid.Column>
          {teamsLength ? (
            <Grid.Row>
              <Grid.Column>
                <Card fluid color="yellow">
                  <Card.Content>
                    <Card.Header as={Header} size="huge">
                      My Teams
                    </Card.Header>
                    <Card.Description>
                      <Item.Group divided>
                        {teams && teams.length
                          ? teams.map(t => (
                              <Item>
                                <Item.Content>
                                  <Item.Header>{t.StaticName}</Item.Header>
                                  <Item.Meta>
                                    {t.School && t.School.Name}
                                  </Item.Meta>
                                  <Item.Description>
                                    <Icon color="green" name="check" />
                                    Ready to go
                                  </Item.Description>
                                </Item.Content>
                              </Item>
                            ))
                          : "Loading"}
                      </Item.Group>
                    </Card.Description>
                  </Card.Content>
                </Card>
              </Grid.Column>
            </Grid.Row>
          ) : null}
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = ({ loginReducer, hubsReducer, teamsReducer }) => ({
  token: loginReducer.token,
  super: loginReducer.superAdmin,
  hubsLength: loginReducer.hubs ? loginReducer.hubs.length : 0,
  teamsLength: loginReducer.teams ? loginReducer.teams.length : 0,
  hubs: hubsReducer.myHubs,
  teams: teamsReducer.myTeams
});

const mapDispatchToProps = {
  getMyHubs: () => getMyHubs.request(),
  getMyTeams: () => getMyTeams.request()
};
// ...
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Homepage);
