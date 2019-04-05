import { createActions, handleActions } from "redux-actions";

const defaultState = {
  myTeams: [],
  allTeams: [],
  allBrcTeams: []
};

export const {
  getMyTeams,
  getAllTeams,
  getBrcTeam,
  registerNewTeam,
  registerBrcTeam
} = createActions({
  GET_MY_TEAMS: {
    REQUEST: () => ({}),
    SUCCESS: teams => ({ teams }),
    FAILURE: error => ({ error })
  },
  GET_ALL_TEAMS: {
    REQUEST: () => ({}),
    SUCCESS: teams => ({ teams }),
    FAILURE: error => ({ error })
  },
  GET_BRC_TEAM: {
    REQUEST: id => ({ id }),
    SUCCESS: (id, seasons, brcTeams) => ({ id, seasons, brcTeams }),
    FAILURE: error => ({ error })
  },
  REGISTER_NEW_TEAM: {
    REQUEST: team => ({ team }),
    SUCCESS: team => ({ team }),
    FAILURE: error => ({ error })
  },
  REGISTER_BRC_TEAM: {
    REQUEST: (id, season) => ({ id, season }),
    SUCCESS: (id, season, brcTeam) => ({ id, season, brcTeam }),
    FAILURE: error => ({ error })
  }
});

const reducer = handleActions(
  {
    LOGOUT: () => defaultState,
    [getMyTeams.success]: (state, { payload: { teams } }) => ({
      ...state,
      myTeams: teams
    }),
    [getAllTeams.success]: (state, { payload: { teams } }) => ({
      ...state,
      allTeams: teams
    }),
    [getBrcTeam.success]: (state, { payload: { id, seasons, brcTeams } }) => ({
      ...state,
      allBrcTeams: {
        ...state.allBrcTeams,
        [id]: seasons.map(s => ({
          ...s,
          brcTeam: brcTeams.find(b => b.Season.ID === s.ID)
        }))
      }
    }),
    [registerBrcTeam.success]: (
      state,
      { payload: { id, season, brcTeam } }
    ) => ({
      ...state,
      allBrcTeams: {
        ...state.allBrcTeams,
        [id]: state.allBrcTeams[id].map(s =>
          s.ID === season
            ? {
                ...s,
                brcTeam
              }
            : s
        )
      }
    })
  },
  defaultState
);

export default reducer;
