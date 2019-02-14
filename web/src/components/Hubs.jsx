import React, { useEffect, useState } from "react";
import axios from "axios";
import { Header, Grid, Table } from "semantic-ui-react";
import NewHubForm from "./NewHubForm";
import FakeRows from "./FakeRows";

const Hubs = ({ setMessage }) => {
  const [hubs, setHubs] = useState([]);
  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/hubs`)
      .then(response => {
        setHubs(response.data);
        setMessage(null);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: "Problem getting hubs",
          content: e.response.message || e.message
        });
      });
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
        <Table>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Name</Table.HeaderCell>
              <Table.HeaderCell>Location</Table.HeaderCell>
              <Table.HeaderCell>Director</Table.HeaderCell>
              <Table.HeaderCell>Description</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          {hubs && hubs.length ? (
            <Table.Body>
              {hubs.map(hub => (
                <Table.Row key={hub.ID}>
                  <Table.Cell>{hub.Name}</Table.Cell>
                  <Table.Cell>{hub.Location}</Table.Cell>
                  <Table.Cell>
                    {hub.Director.FirstName} {hub.Director.LastName}
                  </Table.Cell>
                  <Table.Cell>{hub.Description}</Table.Cell>
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

export default Hubs;
