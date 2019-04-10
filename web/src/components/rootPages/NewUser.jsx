import React, { useState, useEffect } from "react";
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
import ProfileForm from "../userForms/ProfileForm";
import { logout } from "../../redux/login/reducer";
import { connect } from "react-redux";
import { getMyProfile } from "../../redux/users/reducer";

const NewUser = ({ myProfile, getMyProfile }) => {
  const [schoolModalOpen, setSchoolModalOpen] = useState(false);
  const [logoutModalOpen, setLogoutModalOpen] = useState(false);

  useEffect(() => {
    myProfile || getMyProfile();
  }, []);
  console.log(myProfile);
  return (
    <Grid centered columns={2}>
      <Modal
        open={myProfile && (!myProfile.FirstName || myProfile.FirstName === "")}
        closeOnEscape={false}
        closeOnDimmerClick={false}
      >
        <Modal.Header>Complete your profile</Modal.Header>
        <Modal.Content>
          <ProfileForm />
        </Modal.Content>
      </Modal>
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

const mapStateToProps = ({ usersReducer }) => ({
  myProfile: usersReducer.myProfile
});
const mapDispatchToProps = {
  logout: () => logout(),
  getMyProfile: () => getMyProfile.request()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewUser);
