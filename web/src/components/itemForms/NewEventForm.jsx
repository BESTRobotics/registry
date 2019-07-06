import React, { useState, useEffect } from "react";
import axios from "axios";
import { Button, Form, Message, Modal, Header } from "semantic-ui-react";
import NewHubForm from "./NewHubForm";

const NewEventForm = ({ addToList, existingItem, token }) => {
  const headers = { authorization: token };
  const event = existingItem;
  const [name, setName] = useState(event ? event.Name : "");
  const [description, setDescription] = useState(
    event ? event.description : ""
  );
  const [location, setLocation] = useState(event ? event.Location : "");
  const [start, setStart] = useState(event ? event.Start : "");
  const [end, setEnd] = useState(event ? event.End : "");
  const [hub, setHub] = useState(event && event.Hub ? event.Hub.ID : "");
  const [id, setId] = useState(event ? event.ID : "");
  const [message, setMessage] = useState("");
  const [newHub, setNewHub] = useState("");
  const [hubs, setHubs] = useState([]);

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
    const newEvent = {
      Name: name,
      Description: description,
      Location: location,
      Start: start,
      End: end
    };
    let call = axios.post;
    let url = `${process.env.REACT_APP_API_URL}/v1/events`;
    if (id !== "") {
      newEvent.ID = id;
      call = axios.put;
      url = `${process.env.REACT_APP_API_URL}/v1/events/${id}`;
    }
    call(url, newEvent, { headers: headers })
      .then(response => {
        if (!newEvent.ID) {
          newEvent.ID = response.data.ID;
          setId(response.data.ID);
        }
        addToList(newEvent);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem creating event`,
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
        <Form.TextArea
          label="Description"
          value={description}
          onChange={(_, { value }) => setDescription(value)}
        />
        <Form.Input
          label="Location"
          value={location}
          onChange={(_, { value }) => setLocation(value)}
        />
        <Form.Input
          label="Start Time"
          type="datetime-local"
          value={start}
          onChange={(_, { value }) => setStart(value)}
        />
        <Form.Input
          type="datetime-local"
          label="End"
          value={end}
          onChange={(_, { value }) => setEnd(value)}
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
        <Button color="green">{id ? "Update Event" : "Add Event"}</Button>
      </Form>
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
    </React.Fragment>
  );
};

export default NewEventForm;
