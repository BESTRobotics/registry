import React, { useState, useEffect } from "react";
import axios from "axios";
import { Button, Form, Modal, Header, Message } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";
import NewHubForm from "./NewHubForm";
import NewSchoolForm from "./NewSchoolForm";
import PropTypes from "prop-types";

const NewTeamForm = ({ addToList, existingItem, token }) => {
  const headers = { authorization: token };
  const team = existingItem;
  const [users, setUsers] = useState([]);
  const [schools, setSchools] = useState([]);
  const [hubs, setHubs] = useState([]);

  const [id, setId] = useState(team ? team.ID : "");
  const [name, setName] = useState(team ? team.StaticName : "");
  const [hub, setHub] = useState(team ? team.HomeHub.ID : null);
  const [coach, setCoach] = useState(team ? team.Coach.ID : null);
  const [school, setSchool] = useState(team ? team.School.ID : null);
  const [founded, setFounded] = useState(
    team ? team.Founded.substring(0, 10) : ""
  );
  const [mentors, setMentors] = useState(
    team && team.Mentors ? team.Mentors.map(m => m.ID) : []
  );

  const [newUser, setNewUser] = useState("");
  const [newSchool, setNewSchool] = useState("");
  const [newHub, setNewHub] = useState("");
  const [message, setMessage] = useState(null);

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/users`)
      .then(response => {
        setUsers(response.data);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem getting users`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  }, []);

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/schools`)
      .then(response => {
        setSchools(response.data);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem getting schools`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  }, []);

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/hubs`)
      .then(response => {
        setHubs(response.data);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem getting hubs`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  }, []);

  const submitForm = () => {
    const newTeam = {
      StaticName: name,
      Founded: founded ? new Date(founded).toISOString() : null
    };
    let call = axios.post;
    let url = `http://${process.env.REACT_APP_API_URL}/v1/teams`;
    if (id !== "") {
      newTeam.ID = id;
      call = axios.put;
      url = `http://${process.env.REACT_APP_API_URL}/v1/teams/${id}`;
    }
    call(url, newTeam, { headers: headers })
      .then(response => {
        if (!newTeam.ID) {
          newTeam.ID = response.data.ID;
          setId(response.data.ID);
        }
        newTeam.HomeHub = hubs.filter(h => h.ID === hub)[0];
        newTeam.Coach = users.filter(u => u.ID === coach)[0];
        newTeam.School = schools.filter(s => s.ID === school)[0];
        newTeam.Mentors = users.filter(u => mentors.includes(u.ID));
        const { addMentors, subtractMentors } = (() => {
          if (existingItem && existingItem.Mentors) {
            const existingMentors = existingItem.Mentors.map(m => m.ID);
            const addMentors = mentors.filter(
              m => !existingMentors.includes(m)
            );
            const subtractMentors = existingMentors.filter(
              m => !mentors.includes(m)
            );
            return { addMentors, subtractMentors };
          } else {
            return { addMentors: mentors, subtractMentors: [] };
          }
        })();

        let adjustmentRequests = [];
        if (addMentors.length) {
          adjustmentRequests.push(
            ...addMentors.map(m =>
              axios.put(
                `http://${process.env.REACT_APP_API_URL}/v1/teams/${
                  newTeam.ID
                }/mentors`,
                { ID: m },
                { headers: headers }
              )
            )
          );
        }
        if (subtractMentors.length) {
          adjustmentRequests.push(
            ...subtractMentors.map(m =>
              axios.delete(
                `http://${process.env.REACT_APP_API_URL}/v1/teams/${
                  newTeam.ID
                }/mentors/${m}`,
                { headers: headers }
              )
            )
          );
        }
        return axios.all([
          axios.put(
            `http://${process.env.REACT_APP_API_URL}/v1/teams/${
              newTeam.ID
            }/home`,
            { ID: hub },
            { headers: headers }
          ),
          axios.put(
            `http://${process.env.REACT_APP_API_URL}/v1/teams/${
              newTeam.ID
            }/coach`,
            { ID: coach },
            { headers: headers }
          ),
          axios.put(
            `http://${process.env.REACT_APP_API_URL}/v1/teams/${
              newTeam.ID
            }/school`,
            { ID: school },
            { headers: headers }
          ),
          ...adjustmentRequests
        ]);
      })
      .then(() => {
        addToList(newTeam);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem saving team`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  };

  return (
    <React.Fragment>
      {message ? <Message {...message} /> : null}
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
          label="Mentors"
          search
          multiple
          allowAdditions
          loading={!users}
          options={users.map(u => ({
            text: `${u.FirstName} ${u.LastName}`,
            value: u.ID
          }))}
          selection
          value={mentors}
          onChange={(_, { value }) => setMentors(value)}
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
        <Button color="green">{id ? "Update Team" : "Add Team"}</Button>
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

NewTeamForm.propTypes = {
  addToList: PropTypes.func.isRequired
};
