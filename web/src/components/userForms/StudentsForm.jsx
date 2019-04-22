import React, { useState } from "react";
import { connect } from "react-redux";
import { Form, Button } from "semantic-ui-react";
import { updateMyStudents } from "../../redux/users/reducer";
import SingleStudentForm from "./SingleStudentForm";

const StudentsForm = ({ students }) => {
  return (
    <>
      {students &&
        students.map(student => <SingleStudentForm student={student} />)}
      <SingleStudentForm student={null} />
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
