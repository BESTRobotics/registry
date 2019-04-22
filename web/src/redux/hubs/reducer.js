import { createActions, handleActions } from "redux-actions";

const defaultState = {
  myHubs: [],
  allHubs: [],
  seasons: [],
  allBrcHubs: {},
  myBrcHubs: {}
};

export const {
  getMyHubs,
  getAllHubs,
  getBrcHub,
  getSeasonBrcHubs,
  registerBrcHub,
  approveBrcHub,
  getSeasons
} = createActions({
  GET_MY_HUBS: {
    REQUEST: () => ({}),
    SUCCESS: hubs => ({ hubs }),
    FAILURE: error => ({ error })
  },
  GET_ALL_HUBS: {
    REQUEST: () => ({}),
    SUCCESS: hubs => ({ hubs }),
    FAILURE: error => ({ error })
  },
  GET_BRC_HUB: {
    REQUEST: id => ({ id }),
    SUCCESS: (id, seasons, brcHubs) => ({ id, seasons, brcHubs }),
    FAILURE: error => ({ error })
  },
  GET_SEASON_BRC_HUBS: {
    REQUEST: season => ({ season }),
    SUCCESS: (seasons, brcHubs) => ({ seasons, brcHubs }),
    FAILURE: error => ({ error })
  },
  REGISTER_BRC_HUB: {
    REQUEST: (id, season) => ({ id, season }),
    SUCCESS: (id, season, brcHub) => ({ id, season, brcHub }),
    FAILURE: error => ({ error })
  },
  APPROVE_BRC_HUB: {
    REQUEST: (id, season) => ({ id, season }),
    SUCCESS: (id, season) => ({ id, season }),
    FAILURE: error => ({ error })
  },
  GET_SEASONS: {
    REQUEST: () => ({}),
    SUCCESS: seasons => ({ seasons }),
    FAILURE: error => ({ error })
  }
});

const reducer = handleActions(
  {
    LOGOUT: () => defaultState,
    [getMyHubs.success]: (state, { payload: { hubs } }) => ({
      ...state,
      myHubs: hubs
    }),
    [getAllHubs.success]: (state, { payload: { hubs } }) => ({
      ...state,
      allHubs: hubs
    }),
    [getSeasons.success]: (state, { payload: { seasons } }) => ({
      ...state,
      seasons: seasons
    }),
    [getSeasonBrcHubs.success]: (state, { payload: { season, brcHubs } }) => ({
      ...state,
      allBrcHubs: {
        ...state.allBrcHubs,
        [season]: brcHubs
      }
    }),
    [getBrcHub.success]: (state, { payload: { id, seasons, brcHubs } }) => ({
      ...state,
      myBrcHubs: {
        ...state.myBrcHubs,
        seasons: seasons,
        [id]: seasons.map(s => ({
          ...s,
          brcHub: brcHubs.find(b => b.Season.ID === s.ID)
        }))
      }
    }),
    [registerBrcHub.success]: (state, { payload: { id, season, brcHub } }) => ({
      ...state,
      myBrcHubs: {
        ...state.myBrcHubs,
        [id]: state.myBrcHubs[id].map(s =>
          s.ID === season
            ? {
                ...s,
                brcHub
              }
            : s
        )
      }
    }),
    [approveBrcHub.success]: (state, { payload: { id, season } }) => ({
      ...state,
      allBrcHubs: {
        ...state.allBrcHubs,
        [season]: state.allBrcHubs[season].map(s =>
          s.ID === season
            ? {
                ...s,
                [id]: {
                  ...state.AllBrcHubs[season][id],
                  Meta: {
                    ...state.AllBrcHubs[season][id].Meta,
                    BRIApproved: true
                  }
                }
              }
            : s
        )
      }
    })
  },
  defaultState
);

export default reducer;
