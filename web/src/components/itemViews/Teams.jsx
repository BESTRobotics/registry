import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Grid, Card, Header, Item } from "semantic-ui-react";
import { Link } from "react-router-dom";
import { getMyTeams } from "../../redux/teams/reducer";
import FakeItemGroup from "./FakeItemGroup";
import Team from "./Team.jsx";

const Teams = ({ teams, teamsLength, getMyTeams, match }) => {
  useEffect(() => {
    (teams && teams.length) || getMyTeams();
  }, []);
  if (teams && teams.length) {
    if (match && match.params && match.params.id) {
      const team = teams.find(h => String(h.ID) === match.params.id);
      return <Team team={team} />;
    }
    if (teams.length === 1) {
      return <Team team={teams[0]} />;
    }
  }
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
                          <Item.Header as={Link} to={`/teams/${t.ID}`}>
                            {t.StaticName}
                          </Item.Header>
                          <Item.Meta>{t.School && t.School.Name}</Item.Meta>
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
  teams: teamsReducer.myTeams,
  allBrcTeams: teamsReducer.allBrcTeams
});

const mapDispatchToProps = {
  getMyTeams: () => getMyTeams.request()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Teams);
