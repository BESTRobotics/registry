import React, { useEffect, useState } from "react";
import axios from "axios";
import { Card, Button, Header, Grid, Table } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";
import FakeRows from "./FakeRows";

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
        <Table>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Name</Table.HeaderCell>
              <Table.HeaderCell>Email</Table.HeaderCell>
              <Table.HeaderCell>Birthdate</Table.HeaderCell>
              <Table.HeaderCell>Type</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          {users && users.length ? (
            <Table.Body>
              {users.map(user => (
                <Table.Row key={user.ID}>
                  <Table.Cell>
                    {user.FirstName} {user.LastName}
                  </Table.Cell>
                  <Table.Cell>{user.EMail}</Table.Cell>
                  <Table.Cell>{user.Birthdate.substring(0, 10)}</Table.Cell>
                  <Table.Cell>{user.Type}</Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          ) : (
            <FakeRows cols={4} />
          )}
        </Table>
      </Grid.Row>
    </Grid>
  );
};

export default Users;
