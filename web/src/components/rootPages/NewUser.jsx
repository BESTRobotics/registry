import React, { useState, useEffect } from "react";
import {
  Card,
  Dropdown,
  Grid,
  Header,
  Input,
  Button,
  Modal,
  Form
} from "semantic-ui-react";
import NewTeam from "../userForms/NewTeam";
import ProfileForm from "../userForms/ProfileForm";
import { logout } from "../../redux/login/reducer";
import { connect } from "react-redux";
import { getMyProfile, getMyStudents } from "../../redux/users/reducer";
import { getAllTeams } from "../../redux/teams/reducer";
import StudentsForm from "../userForms/StudentsForm";

const NewUser = ({
  myProfile,
  getMyProfile,
  getAllTeams,
  getMyStudents,
  teams,
  myStudents
}) => {
  const [schoolModalOpen, setSchoolModalOpen] = useState(false);
  const [logoutModalOpen, setLogoutModalOpen] = useState(false);

  useEffect(() => {
    myProfile || getMyProfile();
    (teams && teams.length) || getAllTeams();
    (myStudents && myStudents.length) || getMyStudents();
  }, []);
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
                {myStudents && myStudents.length ? (
                  "GOOD JOB YOU HAVE STUDENTS"
                ) : (
                  <StudentsForm />
                )}
                <Form onSubmit={console.log}>
                  <Form.Group inline>
                    <Form.Dropdown
                      placeholder="Select Team"
                      search
                      selection
                      options={
                        teams &&
                        teams.map(t => ({
                          text: `${t.SchoolName} — ${t.StaticName}`,
                          value: t.ID
                        }))
                      }
                    />{" "}
                    <Form.Input
                      icon="lock"
                      iconPosition="left"
                      action="Join Team"
                      placeholder="Secret"
                      disabled={!myStudents || !myStudents.length}
                    />
                  </Form.Group>
                </Form>
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

const mapStateToProps = ({ usersReducer, teamsReducer }) => ({
  myProfile: usersReducer.myProfile,
  myStudents: usersReducer.myStudents,
  teams: teamsReducer.allTeams
});
const mapDispatchToProps = {
  logout: () => logout(),
  getMyProfile: () => getMyProfile.request(),
  getAllTeams: () => getAllTeams.request(),
  getMyStudents: () => getMyStudents.request()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewUser);
