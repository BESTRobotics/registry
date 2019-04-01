import { createActions, handleActions, combineActions } from "redux-actions";
import { logout } from "../login/reducer";

const defaultState = {
  myTeams: [],
  allTeams: []
};

export const { getMyTeams, getAllTeams } = createActions({
  GET_MY_TEAMS: {
    REQUEST: () => ({}),
    SUCCESS: teams => ({ teams }),
    FAILURE: error => ({ error })
  },
  GET_ALL_TEAMS: {
    REQUEST: () => ({}),
    SUCCESS: teams => ({ teams }),
    FAILURE: error => ({ error })
  }
});

const reducer = handleActions(
  {
    [logout]: () => defaultState,
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
