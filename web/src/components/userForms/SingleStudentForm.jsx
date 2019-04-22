import React, { useState } from "react";
import { connect } from "react-redux";
import { Form, Button } from "semantic-ui-react";
import { updateMyStudent } from "../../redux/users/reducer";

const SingleStudentForm = ({ updateStudent, student }) => {
  const [id, _] = useState(student ? student.id : null);
  const [firstName, setFirstName] = useState(student ? student.FirstName : "");
  const [lastName, setLastName] = useState(student ? student.LastName : "");
  const [birthdate, setBirthdate] = useState(
    student && student.Birthdate ? student.Birthdate.substring(0, 10) : ""
  );
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
            birthdate
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
            type="date"
            required
            label="Birthdate"
            value={birthdate}
            onChange={(_, { value }) => setBirthdate(value)}
          />
        </Form.Group>
        <Button color="green">Save Student</Button>
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
