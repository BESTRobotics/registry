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
  const [emailError, setEmailError] = useState(null);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const login = () => {
    setEmailError(null);
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
  const reset = () => {
    setEmailError(null);
    if (!email || email === "") {
        setEmailError("Please enter email to reset password");
        setMessage({
          error: true,
          header: 'Enter email to reset password',
        });
      return
    }
    axios
      .get(`${process.env.REACT_APP_API_URL}/v1/account/local/reset/${email}`)
      .then(response => {
          setMessage({
            header: 'Password Reset Requested',
            content: 'Please check your email for instructions on resetting your password'
          });
      })
      .catch(e => {
        setMessage({
          error: true,
          header: 'Password Reset Failed',
          content:
          e.response && e.response.data ? e.response.data.Message : e.message
        });
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
                required
                icon="user"
                error={emailError}
                iconPosition="left"
                placeholder="Email"
                type="email"
                value={email}
                onChange={(_, { value }) => setEmail(value)}
              />
              <Form.Input
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
                Login
              </Button>
              <Button type="button" fluid size="small" onClick={reset}>Reset Password</Button>

            </Segment>
          </Form>
          <Message>
            New Account? <Link to="/register">Sign Up</Link>           </Message>
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
