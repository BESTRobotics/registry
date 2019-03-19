import React, { useState } from "react";
import axios from "axios";
import { Button, Form, Message } from "semantic-ui-react";
import PropTypes from "prop-types";

const NewSeasonForm = ({ addToList, existingItem }) => {
  const season = existingItem;
  const [name, setName] = useState(season ? season.Name : "");
  const [message, setMessage] = useState(null);
  const [id, setId] = useState(season ? season.ID : "");

  const submitForm = () => {
    const newSeason = {
      Name: name
    };
    let call = axios.post;
    let url = `http://${process.env.REACT_APP_API_URL}/v1/seasons`;
    if (id !== "") {
      newSeason.ID = id;
      call = axios.put;
      url = `http://${process.env.REACT_APP_API_URL}/v1/seasons/${id}/update`;
    }
    call(url, newSeason)
      .then(response => {
        if (!newSeason.ID) {
          newSeason.ID = response.data.ID;
          setId(response.data.ID);
        }
      })
      .then(() => {
        addToList(newSeason);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem saving season`,
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
        <Button color="green">{id ? "Update Season" : "Add Season"}</Button>
      </Form>
    </React.Fragment>
  );
};

export default NewSeasonForm;

NewSeasonForm.propTypes = {
  addToList: PropTypes.func.isRequired
};
