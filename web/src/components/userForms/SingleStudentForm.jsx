import React, { useState } from "react";
import { connect } from "react-redux";
import { Form, Button } from "semantic-ui-react";
import { updateMyStudent } from "../../redux/users/reducer";

const SingleStudentForm = ({ updateStudent, student}) => {
  const [id, _] = useState(student ? student.ID : null);
  const [firstName, setFirstName] = useState(student ? student.FirstName : "");
  const [lastName, setLastName] = useState(student ? student.LastName : "");
  const [email, setEmail] = useState(student ? student.Email : "");
  const [race, setRace] = useState(student ? student.Race : null);
  const [gender, setGender] = useState(student ? student.Race : null);
  return (
    <React.Fragment>
      <Form
        onSubmit={() => {
          updateStudent({
            id,
            firstName,
            lastName,
            email
          });
        }}
      >
        <Form.Group inline>
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
            type="email"
            label="Email"
            value={email}
            onChange={(_, { value }) => setEmail(value)}
          />
          <Button color="green">Save Student</Button>
        </Form.Group>
      </Form>
    </React.Fragment>
  );
};

const mapStateToProps = ({}) => ({});

const mapDispatchToProps = {
  updateStudent: student => updateMyStudent.request(student)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SingleStudentForm);
