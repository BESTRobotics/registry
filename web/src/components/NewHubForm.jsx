import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, Form, Message, Modal, Header } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";
import PropTypes from "prop-types";

const NewHubForm = ({ addToList, existingItem }) => {
  const hub = existingItem;
  const [users, setUsers] = useState([]);
  const [name, setName] = useState(hub ? hub.Name : "");
  const [location, setLocation] = useState(hub ? hub.Location : "");
  const [founded, setFounded] = useState(
    hub ? hub.Founded.substring(0, 10) : ""
  );
  const [description, setDescription] = useState(hub ? hub.Description : "");
  const [id, setId] = useState(hub ? hub.ID : "");
  const [director, setDirector] = useState(
    hub && hub.Director ? hub.Director.ID : null
  );

  const [newUser, setNewUser] = useState("");
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
      Founded: founded !== "" ? new Date(founded).toISOString() : null
    };
    let call = axios.post;
    let url = `http://${process.env.REACT_APP_API_URL}/v1/hubs`;
    if (id !== "") {
      newHub.ID = id;
      call = axios.put;
      url = `http://${process.env.REACT_APP_API_URL}/v1/hubs/${id}/update`;
    }
    call(url, newHub)
      .then(response => {
        if (!newHub.ID) {
          newHub.ID = response.data.ID;
          setId(response.data.ID);
        }
        if (director !== "") {
          newHub.Director = users.filter(u => u.ID === director)[0];
          return axios.put(
            `http://${process.env.REACT_APP_API_URL}/v1/hubs/${
              newHub.ID
            }/director`,
            { ID: director }
          );
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
          allowAdditions
          loading={!users}
          options={users.map(u => ({
            text: `${u.FirstName} ${u.LastName}`,
            value: u.ID
          }))}
          selection
          value={director}
          onChange={(_, { value }) => setDirector(value)}
          onAddItem={(_, { value }) => setNewUser(value)}
        />
        <Button color="green">{hub.ID ? "Update Hub" : "Add Hub"}</Button>
      </Form>
      <Modal open={!!newUser} onClose={() => setNewUser("")}>
        <Header icon="user" content="Add New User" />
        <Modal.Content>
          <NewUserForm
            name={newUser}
            addToList={user => {
              setUsers([...users, user]);
              setDirector(user.ID);
              setNewUser("");
            }}
          />
        </Modal.Content>
      </Modal>
    </React.Fragment>
  );
};

export default NewHubForm;

NewHubForm.propTypes = {
  addToList: PropTypes.func.isRequired
};
