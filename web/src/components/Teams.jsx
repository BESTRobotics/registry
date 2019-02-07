import React, { useEffect, useState } from "react";
import axios from "axios";
import { Card, Button, Header, Grid } from "semantic-ui-react";
// import NewTeamForm from "./NewTeamForm";

const Teams = () => {
  const [teams, setTeams] = useState([]);
  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/teams`)
      .then(response => {
        setTeams(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  return (
    <Grid padded>
      <Grid.Row centered>
        <Grid.Column width={3}>
          <Header>Add a new Team</Header>
          {/* <NewTeamForm /> */}
        </Grid.Column>
      </Grid.Row>
      <Grid.Row>
        <Card.Group>
          {teams.map(team => (
            <Card key={team.ID}>
              <Card.Content>
                <Card.Header>{team.Name}</Card.Header>
                <Card.Meta>
                  {team.Location} &middot; {team.Director.FirstName}{" "}
                  {team.Director.LastName}
                </Card.Meta>
                <Card.Description>{team.Description}</Card.Description>
              </Card.Content>
              <Card.Content extra>
                <Button primary>Team Details</Button>
              </Card.Content>
            </Card>
          ))}
        </Card.Group>
      </Grid.Row>
    </Grid>
  );
};

export default Teams;
