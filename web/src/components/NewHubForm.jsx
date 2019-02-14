import React, { useEffect, useState } from "react";
import axios from "axios";
import { Button, Form, Modal, Header } from "semantic-ui-react";
import NewUserForm from "./NewUserForm";

const NewHubForm = ({ addHub }) => {
  const [users, setUsers] = useState([]);
  const [name, setName] = useState("");
  const [location, setLocation] = useState("");
  const [founded, setFounded] = useState("");
  const [description, setDescription] = useState("");
  const [director, setDirector] = useState(null);

  const [newUser, setNewUser] = useState("");

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/users`)
      .then(response => {
        setUsers(response.data);
      })
      .catch(e => console.log(e));
  }, []);

  const submitForm = () => {
    const newHub = {
      Name: name,
      Location: location,
      Description: description,
      Founded: new Date(founded).toISOString()
    };
    axios
      .post(`http://${process.env.REACT_APP_API_URL}/v1/hubs`, newHub)
      .then(response => {
        newHub.ID = response.data.ID;
        newHub.Director = users.filter(u => u.ID === director)[0];
        return axios.put(
          `http://${process.env.REACT_APP_API_URL}/v1/hubs/${
            response.data.ID
          }/director`,
          { ID: director }
        );
      })
      .then(() => {
        addHub(newHub);
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
        <Button color="green">Add Hub</Button>
      </Form>
      <Modal open={!!newUser} onClose={() => setNewUser("")}>
        <Header icon="user" content="Add New User" />
        <Modal.Content>
          <NewUserForm
            name={newUser}
            addUser={user => {
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
