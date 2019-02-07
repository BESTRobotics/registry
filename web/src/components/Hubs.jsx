import React, { useEffect, useState } from "react";
import axios from "axios";
import { Card, Button, Header, Grid } from "semantic-ui-react";
import NewHubForm from "./NewHubForm";

const Hubs = () => {
  const [hubs, setHubs] = useState([]);
  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/hubs`)
      .then(response => {
        setHubs(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  const addHub = hub => {
    setHubs([...hubs, hub]);
  };

  return (
    <Grid padded>
      <Grid.Row centered>
        <Grid.Column width={3}>
          <Header>Add a new Hub</Header>
          <NewHubForm addHub={addHub} />
        </Grid.Column>
      </Grid.Row>
      <Grid.Row>
        <Card.Group>
          {hubs.map(hub => (
            <Card key={hub.ID}>
              <Card.Content>
                <Card.Header>{hub.Name}</Card.Header>
                <Card.Meta>
                  {hub.Location} &middot; {hub.Director.FirstName}{" "}
                  {hub.Director.LastName}
                </Card.Meta>
                <Card.Description>{hub.Description}</Card.Description>
              </Card.Content>
              <Card.Content extra>
                <Button primary>Hub Details</Button>
              </Card.Content>
            </Card>
          ))}
        </Card.Group>
      </Grid.Row>
    </Grid>
  );
};

export default Hubs;
