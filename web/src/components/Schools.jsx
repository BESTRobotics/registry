import React, { useEffect, useState } from "react";
import axios from "axios";
import { Card, Button, Header, Grid } from "semantic-ui-react";
import NewSchoolForm from "./NewSchoolForm";

const Schools = () => {
  const [schools, setSchools] = useState([]);
  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/schools`)
      .then(response => {
        setSchools(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  const addSchool = school => {
    setSchools([...schools, school]);
  };

  return (
    <Grid padded>
      <Grid.Row centered>
        <Grid.Column width={3}>
          <Header>Add a new School</Header>
          <NewSchoolForm addSchool={addSchool} />
        </Grid.Column>
      </Grid.Row>
      <Grid.Row>
        <Card.Group>
          {schools.map(school => (
            <Card key={school.ID}>
              <Card.Content>
                <Card.Header>{school.Name}</Card.Header>
                <Card.Meta>
                  <a href={school.Website}>{school.Website}</a>
                </Card.Meta>
                <Card.Description>{school.Address}</Card.Description>
              </Card.Content>
              <Card.Content extra>
                <Button primary>School Details</Button>
              </Card.Content>
            </Card>
          ))}
        </Card.Group>
      </Grid.Row>
    </Grid>
  );
};

export default Schools;
