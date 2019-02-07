import React, { useEffect, useState } from "react";
import axios from "axios";
import { Card, Button, Header, Grid } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";

const Users = () => {
  const [users, setUsers] = useState([]);
  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/users`)
      .then(response => {
        setUsers(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  const addUser = user => {
    setUsers([...users, user]);
  };

  return (
    <Grid padded>
      <Grid.Row centered>
        <Grid.Column width={3}>
          <Header>Add a new User</Header>
          <NewUserForm addUser={addUser} />
        </Grid.Column>
      </Grid.Row>
      <Grid.Row>
        <Card.Group>
          {users.map(user => (
            <Card key={user.ID}>
              <Card.Content>
                <Card.Header>
                  {user.FirstName} {Users.LastName}
                </Card.Header>
                <Card.Meta>
                  <a href={`mailto:${user.EMail}`}>{user.EMail}</a>
                </Card.Meta>
                <Card.Description>{user.Type}</Card.Description>
              </Card.Content>
              <Card.Content extra>
                <Button primary>User Details</Button>
              </Card.Content>
            </Card>
          ))}
        </Card.Group>
      </Grid.Row>
    </Grid>
  );
};

export default Users;
