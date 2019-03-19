import { createActions, handleActions, combineActions } from "redux-actions";

const defaultState = {
  myTeams: [],
  allTeams: []
};

export const { getMyTeams, getAllTeams } = createActions({
  GET_MY_TEAMS: {
    REQUEST: () => ({}),
    SUCCESS: hubs => ({ hubs }),
    FAILURE: error => ({ error })
  },
  GET_ALL_TEAMS: {
    REQUEST: () => ({}),
    SUCCESS: hubs => ({ hubs }),
    FAILURE: error => ({ error })
  }
});

const reducer = handleActions(
  {
    [getMyTeams.success]: (state, { payload: { teams } }) => ({
      ...state,
      myTeams: teams
    }),
    [getAllTeams.success]: (state, { payload: { teams } }) => ({
      ...state,
      allTeams: teams
    })
  },
  defaultState
);

export default reducer;
