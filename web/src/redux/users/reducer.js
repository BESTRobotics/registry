import { createActions, handleActions } from "redux-actions";

const defaultState = {
  myProfile: null,
  myStudents: null
};

export const {
  getMyProfile,
  updateMyProfile,
  getMyStudents,
  updateMyStudent,
  addEmptyStudent,
  registerStudents
} = createActions({
  GET_MY_PROFILE: {
    REQUEST: () => ({}),
    SUCCESS: profile => ({ profile }),
    FAILURE: error => ({ error })
  },
  UPDATE_MY_PROFILE: {
    REQUEST: profile => ({ profile }),
    SUCCESS: profile => ({ profile }),
    FAILURE: error => ({ error })
  },
  GET_MY_STUDENTS: {
    REQUEST: () => ({}),
    SUCCESS: students => ({ students }),
    FAILURE: error => ({ error })
  },
  UPDATE_MY_STUDENT: {
    REQUEST: student => ({ student }),
    SUCCESS: student => ({ student }),
    FAILURE: error => ({ error })
  },
  REGISTER_STUDENTS: {
    REQUEST: (students, team, secret) => ({ students, team, secret }),
    SUCCESS: student => ({ student }),
    FAILURE: error => ({ error })
  },
});

const reducer = handleActions(
  {
    LOGOUT: () => defaultState,
    [getMyProfile.success]: (state, { payload: { profile } }) => ({
      ...state,
      myProfile: profile
    }),
    [updateMyProfile.success]: (state, { payload: { profile } }) => ({
      ...state,
      myProfile: profile
    }),
    [getMyStudents.success]: (state, { payload: { students } }) => ({
      ...state,
      myStudents: students
    }),
    [updateMyStudent.success]: (state, { payload: { student } }) => {
      const studentIndex = state.myStudents.findIndex(s => s.ID === student.ID)
      const myStudents = (studentIndex !== -1) ? [...state.myStudents.slice(0, studentIndex), student, ...state.myStudents.slice(studentIndex + 1)] : [...state.myStudents, student]
      return ({
        ...state,
        myStudents
      })
    }
  },
  defaultState
);

export default reducer;
