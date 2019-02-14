import React, { useEffect, useState } from "react";
import axios from "axios";
import { Card, Button, Header, Grid, Table } from "semantic-ui-react";
import NewSchoolForm from "./NewSchoolForm";
import FakeRows from "./FakeRows";

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
        <Table>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Name</Table.HeaderCell>
              <Table.HeaderCell>Website</Table.HeaderCell>
              <Table.HeaderCell>Address</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          {schools && schools.length ? (
            <Table.Body>
              {schools.map(school => (
                <Table.Row key={school.ID}>
                  <Table.Cell>{school.Name}</Table.Cell>
                  <Table.Cell>{school.Website}</Table.Cell>
                  <Table.Cell>{school.Address}</Table.Cell>
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

export default Schools;
