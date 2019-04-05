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
import NewTeam from "../userForms/NewTeam";
import { logout } from "../../redux/login/reducer";
import { connect } from "react-redux";

const NewUser = () => {
  const [schoolModalOpen, setSchoolModalOpen] = useState(false);
  const [logoutModalOpen, setLogoutModalOpen] = useState(false);
  return (
    <Grid centered columns={2}>
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
                    <NewTeam
                      onDone={() => {
                        setSchoolModalOpen(false);
                        setLogoutModalOpen(true);
                      }}
                    />
                  </Modal.Content>
                </Modal>
                <Modal
                  size="tiny"
                  onClose={() => {
                    setLogoutModalOpen(false);
                    logout();
                  }}
                  open={logoutModalOpen}
                >
                  <Modal.Header>Registration Successful</Modal.Header>
                  <Modal.Content>
                    <p>To update your team ownership, please log in again</p>
                  </Modal.Content>

                  <Modal.Actions>
                    <Button
                      content="Okay"
                      positive
                      onClick={() => {
                        setLogoutModalOpen(false);
                        logout();
                      }}
                    />
                  </Modal.Actions>
                </Modal>
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = () => ({});
const mapDispatchToProps = {
  logout: () => logout()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewUser);
