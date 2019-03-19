import React from "react";
import { Card, Dropdown, Grid, Header, Input, Button } from "semantic-ui-react";

const NewUser = () => {
  return (
    <Grid centered columns={2}>
      <Grid.Row>
        <Header as="h1">Welcome to the BEST Registry</Header>
      </Grid.Row>
      <Grid.Row>Help us get you set up</Grid.Row>
      <Grid.Row>
        <Grid.Column>
          <Card fluid color="red">
            <Card.Content>
              <Card.Header as={Header} size="huge">
                I am a Student or Parent in a Team
              </Card.Header>
              <Card.Description>
                <Header>
                  Find your team and enter the secret code your teacher or coach
                  provided:
                </Header>
                <Dropdown
                  placeholder="Select Team"
                  search
                  selection
                  options={[]}
                />{" "}
                <Input
                  icon="lock"
                  iconPosition="left"
                  action="Join Team"
                  placeholder="Secret"
                />
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
      <Grid.Row>
        <Grid.Column>
          <Card fluid color="orange">
            <Card.Content>
              <Card.Header as={Header} size="huge">
                I am a Teacher, Coach, or Administrator of a School
              </Card.Header>
              <Card.Description>
                <Header>Find your school or add a new one</Header>
                <Dropdown
                  placeholder="Select School"
                  search
                  selection
                  options={[]}
                />{" "}
                <Button.Group>
                  <Button>Join School</Button>
                  <Button>Add New School</Button>
                </Button.Group>
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
      <Grid.Row>
        <Grid.Column>
          <Card fluid color="yellow">
            <Card.Content>
              <Card.Header as={Header} size="huge">
                I am a Staff Member or Volunteer of a Hub
              </Card.Header>
              <Card.Description>
                <Header>Find your hub or add a new one</Header>
                <Dropdown
                  placeholder="Select hub"
                  search
                  selection
                  options={[]}
                />{" "}
                <Button.Group>
                  <Button>Join Hub</Button>
                  <Button>Add New Hub</Button>
                </Button.Group>
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

export default NewUser;
