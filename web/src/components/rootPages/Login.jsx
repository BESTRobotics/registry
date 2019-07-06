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
import logo from "../../assets/logo.jpg";
import axios from "axios";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import { setToken as callSetToken } from "../../redux/login/reducer";

const Login = ({ setToken }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const login = () => {
    axios
      .post(`${process.env.REACT_APP_API_URL}/v1/account/login/local`, {
        EMail: email,
        Password: password
      })
      .then(response => {
        setToken(response.data);
      })
      .catch(e => {
        if (e.response && e.response.status === 412) {
          setMessage({
            error: true,
            header: 'Account not validated',
            content: 'Please check your email for instructions on validating your account'
          });
        } else {
          setMessage({
            error: true,
            header: 'Login failed',
            content:
              e.response && e.response.data ? e.response.data.Message : e.message
          });
        }
      });
  };
  return (
    <Grid textAlign="center" verticalAlign="middle" style={{ height: "100%" }}>
      <Grid.Row>
        <Grid.Column style={{ maxWidth: 450 }}>
          <Header as="h2" color="teal" textAlign="center">
            <Image src={logo} />
            Log-in to your account
          </Header>
          <Form size="large" onSubmit={login}>
            <Segment stacked>
              <Form.Input
                fluid
                icon="user"
                iconPosition="left"
                placeholder="Email"
                type="email"
                value={email}
                onChange={(_, { value }) => setEmail(value)}
              />
              <Form.Input
                fluid
                icon="lock"
                iconPosition="left"
                placeholder="Password"
                type="password"
                value={password}
                onChange={(_, { value }) => setPassword(value)}
              />

              <Button color="teal" fluid size="large">
                Login
              </Button>
            </Segment>
          </Form>
          <Message>
            New account? <Link to="/register">Sign Up</Link>
          </Message>
          {message ? <Message {...message} /> : null}
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = () => ({});

const mapDispatchToProps = {
  setToken: token => callSetToken(token)
};
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login);
