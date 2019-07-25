import React, { useState } from "react";
import { connect } from "react-redux";
import { Button } from "semantic-ui-react";
import SingleStudentForm from "./SingleStudentForm";

const StudentsForm = ({ students, done }) => {
  const [newStudents, setNewStudents] = useState(1)
  return (
    <>
      {students &&
        students.map(student => <SingleStudentForm key={student.ID} student={student} />)}
        {Array(newStudents).fill(0).map(() => <SingleStudentForm />)}
        <Button icon="add" onClick={() => setNewStudents(newStudents + 1)}/>
        <Button icon="check" onClick={done}>Done</Button>
    </>
  );
};

const mapStateToProps = ({ usersReducer }) => ({
  students: usersReducer.myStudents
});

const mapDispatchToProps = {};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(StudentsForm);
