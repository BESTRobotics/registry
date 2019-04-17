import React, { useState } from "react";
import {
  Grid,
  Header,
  Segment,
  Button,
  Form,
  Image,
  Message,
  Modal
} from "semantic-ui-react";
import { Link } from "react-router-dom";
import logo from "../../assets/logo.jpg";
import axios from "axios";

const Register = ({ history }) => {
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [message, setMessage] = useState("");
  const [portalOpen, setModalOpen] = useState(false);
  const register = () => {
    const newUser = {
      EMail: email
    };
    axios
      .post(
        `http://${process.env.REACT_APP_API_URL}/v1/account/register/local`,
        { U: newUser, Password: password }
      )
      .then(() => {
        setModalOpen(true);
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
              <Form.Input
                label="Email"
                fluid
                required
                icon="mail"
                iconPosition="left"
                placeholder="Email"
                type="email"
                value={email}
                onChange={(_, { value }) => setEmail(value)}
              />
              <Form.Input
                label="Password"
                fluid
                required
                icon="lock"
                iconPosition="left"
                placeholder="Password"
                type="password"
                value={password}
                onChange={(_, { value }) => setPassword(value)}
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
          <Modal
            size="tiny"
            onClose={() => {
              setModalOpen(false);
              history.push("/login");
            }}
            open={portalOpen}
          >
            <Modal.Header>Registration Successful</Modal.Header>
            <Modal.Content>
              <p>Please check your email to complete sign-up</p>
            </Modal.Content>

            <Modal.Actions>
              <Button
                content="Okay"
                positive
                onClick={() => {
                  setModalOpen(false);
                  history.push("/login");
                }}
              />
            </Modal.Actions>
          </Modal>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

export default Register;
