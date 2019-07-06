import React, { useState, useEffect } from "react";
import axios from "axios";
import { Button, Form, Modal, Header, Message } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";
import PropTypes from "prop-types";

const NewTeamForm = ({ addToList, existingItem, token }) => {
  const headers = { authorization: token };
  const team = existingItem;
  const [users, setUsers] = useState([]);
  const [hubs, setHubs] = useState([]);

  const [id, setId] = useState(team ? team.ID : "");
  const [name, setName] = useState(team ? team.StaticName : "");
  const [schoolName, setSchoolName] = useState(team ? team.SchoolName : "");
  const [schoolAddress, setSchoolAddress] = useState(
    team ? team.SchoolAddress : ""
  );
  const [website, setWebsite] = useState(team ? team.Website : "");
  const [hub, setHub] = useState(team ? team.HomeHub.ID : null);
  const [coaches, setCoaches] = useState(
    team && team.Coaches ? team.Coaches.map(c => c.ID) : []
  );
  const [founded, setFounded] = useState(
    team ? team.Founded.substring(0, 10) : ""
  );
  const [newUser, setNewUser] = useState("");
  const [message, setMessage] = useState(null);

  useEffect(() => {
    axios
      .get(`${process.env.REACT_APP_API_URL}/v1/users`)
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
      .get(`${process.env.REACT_APP_API_URL}/v1/hubs`)
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
      SchoolName: schoolName,
      SchoolAddress: schoolAddress,
      Website: website,
      Founded: founded ? new Date(founded).toISOString() : null,
      HomeHub: { ID: hub },
      Coaches: coaches.map(id => ({ ID: id }))
    };
    let url = `${process.env.REACT_APP_API_URL}/v1/teams`;
    if (id !== "") {
      newTeam.ID = id;
      url = `${process.env.REACT_APP_API_URL}/v1/teams/${id}`;
    }
    axios
      .post(url, newTeam, { headers: headers })
      .then(response => {
        if (!newTeam.ID) {
          newTeam.ID = response.data.ID;
          setId(response.data.ID);
        }
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
          label="Team Name"
          value={name}
          onChange={(_, { value }) => setName(value)}
        />
        <Form.Input
          label="School Name"
          value={schoolName}
          onChange={(_, { value }) => setSchoolName(value)}
        />
        <Form.TextArea
          label="School Address"
          value={schoolAddress}
          onChange={(_, { value }) => setSchoolAddress(value)}
        />
        <Form.Input
          label="Website"
          type="url"
          value={website}
          onChange={(_, { value }) => setWebsite(value)}
        />
        <Form.Dropdown
          label="Hub"
          search
          loading={!hubs}
          options={hubs.map(h => ({
            text: `${h.Name} — ${h.Location}`,
            value: h.ID
          }))}
          selection
          value={hub}
          onChange={(_, { value }) => setHub(value)}
        />
        <Form.Dropdown
          label="Coaches"
          search
          multiple
          loading={!users}
          options={users.map(u => ({
            text: `${u.FirstName} ${u.LastName}`,
            value: u.ID
          }))}
          selection
          value={coaches}
          onChange={(_, { value }) => setCoaches(value)}
        />
        <Form.Input
          type="date"
          label="Founded"
          value={founded}
          onChange={(_, { value }) => setFounded(value)}
        />
        <Button color="green">{id ? "Update Team" : "Add Team"}</Button>
      </Form>
    </React.Fragment>
  );
};

export default NewTeamForm;

NewTeamForm.propTypes = {
  addToList: PropTypes.func.isRequired
};
