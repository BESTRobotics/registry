import React, { useState } from "react";
import axios from "axios";
import { Button, Form, Message } from "semantic-ui-react";

const NewUserForm = ({ addToList, existingItem, name, token }) => {
  const headers = { authorization: token };
  const user = existingItem;
  const [firstName, setFirstName] = useState(
    user ? user.FirstName : name ? name.split(" ")[0] : ""
  );
  const [lastName, setLastName] = useState(
    user ? user.LastName : name ? name.split(" ")[1] : ""
  );
  const [username, setUsername] = useState(user ? user.Username : "");
  const [id, setId] = useState(user ? user.ID : "");
  const [email, setEmail] = useState(user ? user.EMail : "");
  const [birthdate, setBirthdate] = useState(
    user && user.Birthdate ? user.Birthdate.substring(0, 10) : ""
  );
  const [type, setType] = useState(user ? user.Type : "");
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
    let call = axios.post;
    let url = `${process.env.REACT_APP_API_URL}/v1/users`;
    if (id !== "") {
      newUser.ID = id;
      call = axios.put;
      url = `${process.env.REACT_APP_API_URL}/v1/users/${id}`;
    }
    call(url, newUser, { headers: headers })
      .then(response => {
        if (!newUser.ID) {
          newUser.ID = response.data.ID;
          setId(response.data.ID);
        }
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
          value={birthdate || ""}
          onChange={(_, { value }) => setBirthdate(value)}
        />
        <Form.Input
          label="Type"
          value={type}
          onChange={(_, { value }) => setType(value)}
        />
        <Button color="green">{id ? "Update User" : "Add User"}</Button>
      </Form>
    </React.Fragment>
  );
};

export default NewUserForm;
