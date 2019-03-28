import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Grid, Card, Header, Item, Icon } from "semantic-ui-react";
import { getMyTeams } from "../../redux/teams/reducer";
import FakeItemGroup from "./FakeItemGroup";

const Teams = ({ teams, teamsLength, getMyTeams }) => {
  useEffect(() => {
    (teams && teams.length) || getMyTeams();
  }, []);
  return (
    <Grid columns={2} centered>
      <Grid.Row>
        <Grid.Column>
          <Card fluid color="yellow">
            <Card.Content>
              <Card.Header as={Header} size="huge">
                My Teams
              </Card.Header>
              <Card.Description>
                {teams && teams.length ? (
                  <Item.Group divided>
                    {teams.map(t => (
                      <Item>
                        <Item.Content>
                          <Item.Header>{t.StaticName}</Item.Header>
                          <Item.Meta>{t.School && t.School.Name}</Item.Meta>
                          <Item.Description>
                            <Icon color="green" name="check" />
                            Ready to go
                          </Item.Description>
                        </Item.Content>
                      </Item>
                    ))}
                  </Item.Group>
                ) : (
                  <FakeItemGroup rows={teamsLength} />
                )}
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = ({ loginReducer, teamsReducer }) => ({
  teamsLength: loginReducer.teams ? loginReducer.teams.length : 0,
  teams: teamsReducer.myTeams
});

const mapDispatchToProps = {
  getMyTeams: () => getMyTeams.request()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Teams);
