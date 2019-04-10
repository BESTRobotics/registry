import React, { useState } from "react";
import { connect } from "react-redux";
import { Form, Button } from "semantic-ui-react";
import { updateMyProfile } from "../../redux/users/reducer";

const ProfileForm = ({ updateMyProfile }) => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [birthdate, setBirthdate] = useState("");
  return (
    <React.Fragment>
      <Form
        onSubmit={() => {
          updateMyProfile({
            firstName,
            lastName,
            birthdate
          });
        }}
      >
        <Form.Input
          label="First Name"
          required
          value={firstName}
          onChange={(_, { value }) => setFirstName(value)}
        />
        <Form.Input
          label="Last Name"
          required
          value={lastName}
          onChange={(_, { value }) => setLastName(value)}
        />
        <Form.Input
          type="date"
          required
          label="Birthdate"
          value={birthdate}
          onChange={(_, { value }) => setBirthdate(value)}
        />
        <Button color="green">Update Profile</Button>
      </Form>
    </React.Fragment>
  );
};

const mapStateToProps = ({ usersReducer }) => ({
  profile: usersReducer.myProfile
});

const mapDispatchToProps = {
  updateMyProfile: profile => updateMyProfile.request(profile)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ProfileForm);
