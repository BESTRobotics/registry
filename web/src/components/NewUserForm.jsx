import React, { useState } from "react";
import axios from "axios";
import { Button, Form } from "semantic-ui-react";

const NewUserForm = addUser => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [birthdate, setBirthdate] = useState("");
  const [type, setType] = useState("");

  const submitForm = () => {
    const newUser = {
      Username: username,
      EMail: email,
      Type: type,
      FirstName: firstName,
      LastName: lastName,
      Birthdate: birthdate
    };
    axios
      .post(`http://${process.env.REACT_APP_API_URL}/v1/users`, newUser)
      .then(response => {
        console.log(response.data);
        addUser(newUser);
      })
      .catch(e => console.log(e));
  };

  return (
    <Form onSubmit={submitForm}>
      <Form.Group inline>
        <Form.Input
          label="First Name"
          value={firstName}
          onChange={(_, { value }) => setFirstName(value)}
        />
        <Form.Input
          label="Last Name"
          value={lastName}
          onChange={(_, { value }) => setLastName(value)}
        />
      </Form.Group>
      <Form.Input
        label="username"
        value={username}
        onChange={(_, { value }) => setUsername(value)}
      />
      <Form.Input
        label="email"
        type="email"
        value={email}
        onChange={(_, { value }) => setEmail(value)}
      />
      <Form.Input
        type="date"
        label="Birthdate"
        value={birthdate}
        onChange={(_, { value }) => setBirthdate(value)}
      />
      <Form.Input
        label="Type"
        value={type}
        onChange={(_, { value }) => setType(value)}
      />
      <Button color="green">Add User</Button>
    </Form>
  );
};

export default NewUserForm;
