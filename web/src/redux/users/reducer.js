import { createActions, handleActions } from "redux-actions";

const defaultState = {
  myProfile: null,
  myStudents: null
};

export const {
  getMyProfile,
  updateMyProfile,
  getMyStudents,
  updateMyStudent
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
  }
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
    })
  },
  defaultState
);

export default reducer;
