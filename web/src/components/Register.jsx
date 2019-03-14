import React, { useState } from "react";
import {
  Grid,
  Header,
  Segment,
  Button,
  Form,
  Image,
  Message
} from "semantic-ui-react";
import { Link } from "react-router-dom";
import logo from "../assets/logo.jpg";
import axios from "axios";

const Register = ({ history }) => {
  const [username, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [type, setType] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [birthdate, setBirthdate] = useState("");
  const [message, setMessage] = useState("");
  const register = () => {
    const newUser = {
      Username: username,
      EMail: email,
      Type: type,
      FirstName: firstName,
      LastName: lastName,
      Birthdate: birthdate ? new Date(birthdate).toISOString() : null
    };
    axios
      .post(
        `http://${process.env.REACT_APP_API_URL}/v1/account/register/local`,
        { U: newUser, Password: password }
      )
      .then(response => {
        history.push("/login");
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Registration failed`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  };
  return (
    <Grid textAlign="center" verticalAlign="middle" style={{ height: "100%" }}>
      <Grid.Row>
        <Grid.Column style={{ maxWidth: 750 }}>
          <Header as="h2" color="teal" textAlign="center">
            <Image src={logo} />
            Create a new account
          </Header>
          <Form size="large" onSubmit={register}>
            <Segment stacked>
              <Form.Group widths="equal">
                <Form.Input
                  label="Username"
                  fluid
                  icon="user"
                  iconPosition="left"
                  placeholder="Username"
                  value={username}
                  onChange={(_, { value }) => setUserName(value)}
                />
                <Form.Input
                  label="Password"
                  fluid
                  icon="lock"
                  iconPosition="left"
                  placeholder="Password"
                  type="password"
                  value={password}
                  onChange={(_, { value }) => setPassword(value)}
                />
              </Form.Group>
              <Form.Group widths="equal">
                <Form.Input
                  label="First Name"
                  fluid
                  placeholder="First Name"
                  value={firstName}
                  onChange={(_, { value }) => setFirstName(value)}
                />
                <Form.Input
                  label="Last Name"
                  fluid
                  placeholder="last Name"
                  value={lastName}
                  onChange={(_, { value }) => setLastName(value)}
                />
              </Form.Group>
              <Form.Group widths="equal">
                <Form.Input
                  label="Email"
                  fluid
                  icon="mail"
                  iconPosition="left"
                  placeholder="Email"
                  type="email"
                  value={email}
                  onChange={(_, { value }) => setEmail(value)}
                />

                <Form.Input
                  label="Birthdate"
                  fluid
                  icon="calendar"
                  iconPosition="left"
                  placeholder="Birthdate"
                  type="date"
                  value={birthdate}
                  onChange={(_, { value }) => setBirthdate(value)}
                />
              </Form.Group>

              <Form.Input
                label="Type"
                fluid
                icon="settings"
                iconPosition="left"
                placeholder="Type"
                value={type}
                onChange={(_, { value }) => setType(value)}
              />

              <Button color="teal" fluid size="large">
                Create Account
              </Button>
            </Segment>
          </Form>
          <Message>
            Already have an account? <Link to="/login">Log in</Link>
          </Message>
          {message ? <Message {...message} /> : null}
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

export default Register;
