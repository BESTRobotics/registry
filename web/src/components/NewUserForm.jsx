import React, { useState } from "react";
import axios from "axios";
import { Button, Form, Message } from "semantic-ui-react";

const NewUserForm = ({ addToList, name }) => {
  const [firstName, setFirstName] = useState(name ? name.split(" ")[0] : "");
  const [lastName, setLastName] = useState(name ? name.split(" ")[1] : "");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [birthdate, setBirthdate] = useState("");
  const [type, setType] = useState("");
  const [message, setMessage] = useState("");

  const submitForm = () => {
    const newUser = {
      Username: username,
      EMail: email,
      Type: type,
      FirstName: firstName,
      LastName: lastName,
      Birthdate: birthdate ? new Date(birthdate).toISOString() : null
    };
    axios
      .post(`http://${process.env.REACT_APP_API_URL}/v1/users`, newUser)
      .then(response => {
        newUser.ID = response.data.ID;
        addToList(newUser);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem creating user`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  };

  return (
    <React.Fragment>
      {message ? <Message {...message} /> : null}
      <Form onSubmit={submitForm}>
        <Form.Group>
          <Form.Input
            label="First Name"
            value={firstName}
            width={8}
            onChange={(_, { value }) => setFirstName(value)}
          />
          <Form.Input
            label="Last Name"
            value={lastName}
            width={8}
            onChange={(_, { value }) => setLastName(value)}
          />
        </Form.Group>
        <Form.Input
          label="Username"
          value={username}
          onChange={(_, { value }) => setUsername(value)}
        />
        <Form.Input
          label="Email"
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
    </React.Fragment>
  );
};

export default NewUserForm;
