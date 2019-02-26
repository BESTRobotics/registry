import React, { useState } from "react";
import axios from "axios";
import { Button, Form, Message } from "semantic-ui-react";

const NewSchoolForm = ({ addToList, existingItem }) => {
  const school = existingItem;
  const [name, setName] = useState(school ? school.Name : "");
  const [address, setAddress] = useState(school ? school.Address : "");
  const [website, setWebsite] = useState(school ? school.Website : "");
  const [message, setMessage] = useState(null);
  const [id, setId] = useState(school ? school.ID : "");

  const submitForm = () => {
    const newSchool = {
      Name: name,
      Address: address,
      Website: website
    };
    let call = axios.post;
    let url = `http://${process.env.REACT_APP_API_URL}/v1/hubs`;
    if (id !== "") {
      newSchool.ID = id;
      call = axios.put;
      url = `http://${process.env.REACT_APP_API_URL}/v1/hubs/${id}/update`;
    }
    call(url, newSchool)
      .then(response => {
        if (!newSchool.ID) {
          newSchool.ID = response.data.ID;
          setId(response.data.ID);
        }
        addToList(newSchool);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem getting users`,
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
          label="Address"
          value={address}
          onChange={(_, { value }) => setAddress(value)}
        />
        <Form.Input
          label="Website"
          type="url"
          value={website}
          onChange={(_, { value }) => setWebsite(value)}
        />
        <Button color="green">
          {school && school.ID ? "Update School" : "Add School"}
        </Button>
      </Form>
    </React.Fragment>
  );
};

export default NewSchoolForm;
