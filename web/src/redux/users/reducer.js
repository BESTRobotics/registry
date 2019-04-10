import { createActions, handleActions } from "redux-actions";

const defaultState = {
  myProfile: null
};

export const { getMyProfile, updateMyProfile } = createActions({
  GET_MY_PROFILE: {
    REQUEST: () => ({}),
    SUCCESS: profile => ({ profile }),
    FAILURE: error => ({ error })
  },
  UPDATE_MY_PROFILE: {
    REQUEST: profile => ({ profile }),
    SUCCESS: profile => ({ profile }),
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
    })
  },
  defaultState
);

export default reducer;
