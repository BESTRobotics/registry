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
  getSeasons,
  saveSeason
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
    REQUEST: (hubid, brchubid, season) => ({ hubid, brchubid, season }),
    SUCCESS: (id, season) => ({ id, season }),
    FAILURE: error => ({ error })
  },
  GET_SEASONS: {
    REQUEST: () => ({}),
    SUCCESS: seasons => ({ seasons }),
    FAILURE: error => ({ error })
  },
  SAVE_SEASON: {
    REQUEST: season => ({ season }),
    SUCCESS: season => ({ season }),
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
    [approveBrcHub.success]: (
      state,
      { payload: { hubid, brchubid, season } }
    ) => {
      return {
        ...state,
        allBrcHubs: {
          ...state.allBrcHubs,
          [season]: state.allBrcHubs[season].map(s =>
            s.ID === hubid
              ? {
                  ...s,
                  [brchubid]: {
                    ...state.AllBrcHubs[season][hubid],
                    Meta: {
                      ...state.AllBrcHubs[season][hubid].Meta,
                      BRIApproved: true
                    }
                  }
                }
              : s
          )
        }
      };
    },
    [saveSeason.success]: (state, { payload: { id, season } }) => {
      console.log(id);
      console.log(season);
      let seasons = state.seasons;
      const seasonIndex = state.seasons.findIndex(s => s.ID === id);
      if (seasonIndex === -1) {
        seasons = [...seasons, season];
      } else {
        seasons = [
          ...seasons.slice(0, seasonIndex - 1),
          season,
          ...season.slice(seasonIndex)
        ];
      }
      return {
        ...state,
        seasons
      };
    }
  },
  defaultState
);

export default reducer;
