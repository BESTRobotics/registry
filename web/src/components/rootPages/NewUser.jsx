import React, { useState, useEffect } from "react";
import {
  Card,
  Grid,
  Header,
  Button,
  Modal,
  Form
} from "semantic-ui-react";
import NewTeam from "../userForms/NewTeam";
import ProfileForm from "../userForms/ProfileForm";
import { logout } from "../../redux/login/reducer";
import { connect } from "react-redux";
import { getMyProfile, getMyStudents, registerStudents } from "../../redux/users/reducer";
import { getSeasons } from "../../redux/hubs/reducer";
import StudentsForm from "../userForms/StudentsForm";

const NewUser = ({
  myProfile,
  getMyProfile,
  getSeasons,
  getMyStudents,
  seasons,
  myStudents,
  registerStudents,
}) => {
  const [schoolModalOpen, setSchoolModalOpen] = useState(false);
  const [selectedStudents, setSelectedStudents] = useState([]);
  const [logoutModalOpen, setLogoutModalOpen] = useState(false);
  const [editingStudents, setEditingStudents] = useState(false);
  const [selectedSeason, setSelectedSeason] = useState(null);
  const [secret, setSecret] = useState("");

  useEffect(() => {
    myProfile || getMyProfile();
    (seasons && seasons.length) || getSeasons();
    (myStudents && myStudents.length) || getMyStudents();
  }, []);

  useEffect(() => {
    myStudents && setSelectedStudents(myStudents.map(s => s.ID))
  }, [myStudents])

  return (
    <Grid centered columns={2}>
      <Modal closeOnDimmerClick={false}
        open={myProfile && (!myProfile.FirstName || myProfile.FirstName === "")}
        closeOnEscape={false}
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
                {myStudents && myStudents.length && !editingStudents ? (
                  <Form onSubmit={() => registerStudents(selectedStudents, selectedSeason, secret)}>
                <Header>
                  Find your team and enter the secret code your teacher or coach
                  provided:
                </Header>
                {myStudents.map(s =>
                  <Form.Checkbox
                    checked={selectedStudents.includes(s.ID)}
                    label={`${s.FirstName} ${s.LastName}`}
                    onChange={() => setSelectedStudents(selectedStudents.includes(s.ID) ? selectedStudents.filter(f => f !== s.ID) : [...selectedStudents, s.ID])}
                  />)}
                  <Button onClick={() => setEditingStudents(true)}>Edit</Button>
                  <Form.Group inline>
                    <Form.Dropdown
                      placeholder="Select Season"
                      search
                      selection
                      value={selectedSeason}
                      onChange={(_, { value }) => setSelectedSeason(value)}
                      options={
                        seasons &&
                        seasons.map(s => ({
                          text: `${s.Name}`,
                          value: s.ID
                        }))
                      }
                    />{" "}
                    <Form.Input
                      icon="lock"
                      iconPosition="left"
                      value={secret}
                      onChange={(_, { value} ) => setSecret(value)}
                      action="Join Team"
                      placeholder="Secret"
                      disabled={!myStudents || !myStudents.length}
                    />
                  </Form.Group>
                </Form>
                ) : (
                  <StudentsForm done={() => setEditingStudents(false)}/>
                )}
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
                <Header>
                  If your school is not in the registry, you can a add one
                </Header>
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
                      onDone={(success) => {
                        setSchoolModalOpen(false);
                        if (success) {setLogoutModalOpen(true);}
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

const mapStateToProps = ({ usersReducer, hubsReducer }) => ({
  myProfile: usersReducer.myProfile,
  myStudents: usersReducer.myStudents,
  seasons: hubsReducer.seasons
});
const mapDispatchToProps = {
  logout: () => logout(),
  getMyProfile: () => getMyProfile.request(),
  getSeasons: () => getSeasons.request(),
  getMyStudents: () => getMyStudents.request(),
  registerStudents: (selectedStudents, selectedSeason, secret) => registerStudents.request(selectedStudents, selectedSeason, secret)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewUser);
