import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, Form } from "semantic-ui-react";

const NewHubForm = () => {
  const [users, setUsers] = useState([]);
  const [name, setName] = useState("");
  const [location, setLocation] = useState("");
  const [founded, setFounded] = useState("");
  const [director, setDirector] = useState(null);

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/users`)
      .then(response => {
        setUsers(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  const submitForm = () => {
    axios
      .post(`http://${process.env.REACT_APP_API_URL}/v1/hubs`, {
        Name: name,
        Location: location,
        Founded: founded,
        Director: { ID: director }
      })
      .then(response => {
        console.log(response.data);
      })
      .catch(e => console.log(e));
  };

  return (
    <Form onSubmit={submitForm}>
      <Form.Input
        label="Name"
        value={name}
        onChange={(_, { value }) => setName(value)}
      />
      <Form.Input
        label="Location"
        value={location}
        onChange={(_, { value }) => setLocation(value)}
      />
      <Form.Input
        type="date"
        label="Founded"
        value={founded}
        onChange={(_, { value }) => setFounded(value)}
      />
      <Form.Dropdown
        label="Director"
        search
        loading={!users}
        options={users.map(u => ({
          text: `${u.FirstName} ${u.LastName}`,
          value: u.ID
        }))}
        selection
        value={director}
        onChange={(_, { value }) => setDirector(value)}
      />
      <Button color="green">Add Hub</Button>
    </Form>
  );
};

export default NewHubForm;
