import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Card, Grid, Header, Icon, List, Item } from "semantic-ui-react";
import { Link } from "react-router-dom";
import { getMyHubs } from "../../redux/hubs/reducer";
import { getMyTeams } from "../../redux/hubs/reducer";

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
                  <List>
                    <List.Item>
                      <List.Icon name="marker" />
                      <List.Content>
                        <List.Header as="a" Capital BEST />
                        <List.Description>Austin, Texas</List.Description>
                      </List.Content>
                    </List.Item>
                    <List.Item>
                      <List.Icon name="marker" />
                      <List.Content>
                        <List.Header as="a">SA BEST</List.Header>
                        <List.Description>San Antonio, Texas</List.Description>
                      </List.Content>
                    </List.Item>
                  </List>
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
                <Card fluid color="orange">
                  <Card.Content>
                    <Card.Header as={Header} size="huge">
                      My Teams
                    </Card.Header>
                    <Card.Description>
                      <Item.Group divided>
                        <Item>
                          <Item.Content>
                            <Item.Header>LASA Robotics</Item.Header>
                            <Item.Meta>33 Students</Item.Meta>
                            <Item.Description>
                              <Icon color="green" name="check" />
                              Ready to go
                            </Item.Description>
                          </Item.Content>
                        </Item>
                        <Item>
                          <Item.Content>
                            <Item.Header>One Day Academy</Item.Header>
                            <Item.Meta>2 Students</Item.Meta>
                            <Item.Description>
                              <Icon color="orange" name="warning" />
                              Action Required: <Link to="/">Form 1090-T</Link>
                            </Item.Description>
                          </Item.Content>
                        </Item>
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
  fetchHubs: () => getMyHubs.request,
  fetchTeams: () => getMyTeams.request
};
// ...
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Homepage);
