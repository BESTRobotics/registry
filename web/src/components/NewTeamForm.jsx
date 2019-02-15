import React, { useState, useEffect } from "react";
import axios from "axios";
import { Button, Form, Modal, Header } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";
import NewHubForm from "./NewHubForm";
import NewSchoolForm from "./NewSchoolForm";

const NewTeamForm = ({ addToList }) => {
  const [users, setUsers] = useState([]);
  const [schools, setSchools] = useState([]);
  const [hubs, setHubs] = useState([]);

  const [name, setName] = useState("");
  const [hub, setHub] = useState(null);
  const [coach, setCoach] = useState(null);
  const [school, setSchool] = useState(null);
  const [founded, setFounded] = useState("");

  const [newUser, setNewUser] = useState("");
  const [newSchool, setNewSchool] = useState("");
  const [newHub, setNewHub] = useState("");

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/users`)
      .then(response => {
        setUsers(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/schools`)
      .then(response => {
        setSchools(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/hubs`)
      .then(response => {
        setHubs(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  const submitForm = () => {
    const newTeam = {
      StaticName: name,
      Founded: founded ? new Date(founded).toISOString() : null
    };
    axios
      .post(`http://${process.env.REACT_APP_API_URL}/v1/teams`, newTeam)
      .then(response => {
        newTeam.ID = response.data.ID;
        newTeam.HomeHub = hubs.filter(h => h.ID === hub)[0];
        newTeam.Coach = users.filter(u => u.ID === coach)[0];
        newTeam.School = schools.filter(s => s.ID === school)[0];
        return axios.all([
          axios.put(
            `http://${process.env.REACT_APP_API_URL}/v1/teams/${
              response.data.ID
            }/home`,
            { ID: hub }
          ),
          axios.put(
            `http://${process.env.REACT_APP_API_URL}/v1/teams/${
              response.data.ID
            }/coach`,
            { ID: coach }
          ),
          axios.put(
            `http://${process.env.REACT_APP_API_URL}/v1/teams/${
              response.data.ID
            }/school`,
            { ID: school }
          )
        ]);
      })
      .then(() => {
        addToList(newTeam);
      })
      .catch(e => console.log(e));
  };

  return (
    <React.Fragment>
      <Form onSubmit={submitForm}>
        <Form.Input
          label="Name"
          value={name}
          onChange={(_, { value }) => setName(value)}
        />
        <Form.Dropdown
          label="Hub"
          search
          allowAdditions
          loading={!hubs}
          options={hubs.map(h => ({
            text: h.Name,
            value: h.ID
          }))}
          selection
          value={hub}
          onChange={(_, { value }) => setHub(value)}
          onAddItem={(_, { value }) => setNewHub(value)}
        />
        <Form.Dropdown
          label="Coach"
          search
          allowAdditions
          loading={!users}
          options={users.map(u => ({
            text: `${u.FirstName} ${u.LastName}`,
            value: u.ID
          }))}
          selection
          value={coach}
          onChange={(_, { value }) => setCoach(value)}
          onAddItem={(_, { value }) => setNewUser(value)}
        />
        <Form.Dropdown
          label="School"
          search
          allowAdditions
          loading={!schools}
          options={schools.map(s => ({
            text: s.Name,
            value: s.ID
          }))}
          selection
          value={school}
          onChange={(_, { value }) => setSchool(value)}
          onAddItem={(_, { value }) => setNewSchool(value)}
        />
        <Form.Input
          type="date"
          label="Founded"
          value={founded}
          onChange={(_, { value }) => setFounded(value)}
        />
        <Button color="green">Add Team</Button>
      </Form>
      <Modal open={!!newUser} onClose={() => setNewUser("")}>
        <Header icon="user" content="Add New User" />
        <Modal.Content>
          <NewUserForm
            name={newUser}
            addToList={user => {
              setUsers([...users, user]);
              setCoach(user.ID);
              setNewUser("");
            }}
          />
        </Modal.Content>
      </Modal>

      <Modal open={!!newHub} onClose={() => setNewHub("")}>
        <Header icon="map marker" content="Add New Hub" />
        <Modal.Content>
          <NewHubForm
            name={newHub}
            addToList={hub => {
              setHubs([...hubs, hub]);
              setHub(hub.ID);
              setNewHub("");
            }}
          />
        </Modal.Content>
      </Modal>

      <Modal open={!!newSchool} onClose={() => setNewSchool("")}>
        <Header icon="house" content="Add New School" />
        <Modal.Content>
          <NewSchoolForm
            name={newSchool}
            addToList={school => {
              setSchools([...schools, school]);
              setSchool(school.ID);
              setNewSchool("");
            }}
          />
        </Modal.Content>
      </Modal>
    </React.Fragment>
  );
};

export default NewTeamForm;
