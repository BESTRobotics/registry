import React, { useState } from "react";
import axios from "axios";
import { Button, Form } from "semantic-ui-react";

const NewSchoolForm = ({ addSchool }) => {
  const [name, setName] = useState("");
  const [address, setAddress] = useState("");
  const [website, setWebsite] = useState("");

  const submitForm = () => {
    const newSchool = {
      Name: name,
      Address: address,
      Website: website
    };
    axios
      .post(`http://${process.env.REACT_APP_API_URL}/v1/schools`, newSchool)
      .then(response => {
        console.log(response.data);
        addSchool(newSchool);
      })
      .catch(e => console.log(e));
  };

  return (
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
      <Button color="green">Add School</Button>
    </Form>
  );
};

export default NewSchoolForm;
