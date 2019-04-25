import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, Form, Message, Modal, Header } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";
import PropTypes from "prop-types";

const NewHubForm = ({ addToList, existingItem, token }) => {
  const headers = { authorization: token };
  const hub = existingItem;
  const [users, setUsers] = useState([]);
  const [name, setName] = useState(hub ? hub.Name : "");
  const [location, setLocation] = useState(hub ? hub.Location : "");
  const [founded, setFounded] = useState(
    hub && hub.Founded ? hub.Founded.substring(0, 10) : ""
  );
  const [description, setDescription] = useState(hub ? hub.Description : "");
  const [id, setId] = useState(hub ? hub.ID : "");
  const [director, setDirector] = useState(
    hub && hub.Director ? hub.Director.ID : null
  );

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

  const submitForm = () => {
    const newHub = {
      Name: name,
      Location: location,
      Description: description,
      Founded: founded !== "" ? new Date(founded).toISOString() : null,
      Director: { ID: director }
    };
    let url = `http://${process.env.REACT_APP_API_URL}/v1/hubs`;
    if (id !== "") {
      newHub.ID = id;
      url = `http://${process.env.REACT_APP_API_URL}/v1/hubs/${id}`;
    }
    axios
      .post(url, newHub, { headers: headers })
      .then(response => {
        if (!newHub.ID) {
          newHub.ID = response.data.ID;
          setId(response.data.ID);
        }
      })
      .then(() => {
        addToList(newHub);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem saving hub`,
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
        <Form.Input
          label="Location"
          value={location}
          onChange={(_, { value }) => setLocation(value)}
        />
        <Form.TextArea
          label="Description"
          value={description}
          onChange={(_, { value }) => setDescription(value)}
        />
        <Form.Input
          type="date"
          label="Founded"
          value={founded}
          onChange={(_, { value }) => setFounded(value)}
        />
        <Form.Dropdown
          label="Director"
          search
          loading={!users}
          options={users.map(u => ({
            text: `${u.FirstName} ${u.LastName}`,
            value: u.ID
          }))}
          selection
          value={director}
          onChange={(_, { value }) => setDirector(value)}
        />
        <Button color="green">{id ? "Update Hub" : "Add Hub"}</Button>
      </Form>
    </React.Fragment>
  );
};

export default NewHubForm;

NewHubForm.propTypes = {
  addToList: PropTypes.func.isRequired
};
