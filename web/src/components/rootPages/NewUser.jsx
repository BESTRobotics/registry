import React, { useState } from "react";
import {
  Card,
  Dropdown,
  Grid,
  Header,
  Input,
  Button,
  Modal
} from "semantic-ui-react";
import NewTeamForm from "../itemForms/NewTeamForm";
import { connect } from "react-redux";

const NewUser = ({ token }) => {
  const [schoolModalOpen, setSchoolModalOpen] = useState(false);
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
                <Header>If your school doesn't exist, you can a new one</Header>
                <Modal
                  trigger={
                    <Button onClick={() => setSchoolModalOpen(true)}>
                      Add a new team
                    </Button>
                  }
                  onOpen={() => setSchoolModalOpen(true)}
                  onClose={() => setSchoolModalOpen(false)}
                  open={!!schoolModalOpen}
                >
                  <Modal.Header>New Team</Modal.Header>
                  <Modal.Content>
                    <NewTeamForm
                      addToList={() => setSchoolModalOpen(false)}
                      token={token}
                    />
                  </Modal.Content>
                </Modal>
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = ({ loginReducer }) => ({ token: loginReducer.token });

export default connect(mapStateToProps)(NewUser);
