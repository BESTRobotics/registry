import { createActions, handleActions } from "redux-actions";

const defaultState = {
  myHubs: [],
  allHubs: {},
  allBrcHubs: {}
};

export const { getMyHubs, getAllHubs, getBrcHub, registerBrc } = createActions({
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
  REGISTER_BRC: {
    REQUEST: (id, season) => ({ id, season }),
    SUCCESS: (id, season, brcHub) => ({ id, season, brcHub }),
    FAILURE: error => ({ error })
  }
});

const reducer = handleActions(
  {
    LOGOUT: () => defaultState,
    [getMyHubs.request]: state => state,
    [getMyHubs.success]: (state, { payload: { hubs } }) => ({
      ...state,
      myHubs: hubs,
      allHubs: {
        ...state.allHubs,
        ...hubs.reduce((o, h) => {
          o[h.id] = h;
          return o;
        }, {})
      }
    }),
    [getAllHubs.success]: (state, { payload: { hubs } }) => ({
      ...state,
      allHubs: hubs
    }),
    [getBrcHub.success]: (state, { payload: { id, seasons, brcHubs } }) => ({
      ...state,
      allBrcHubs: {
        ...state.allBrcHubs,
        [id]: seasons.map(s => ({
          ...s,
          brcHub: brcHubs.find(b => b.Season.ID === s.ID)
        }))
      }
    }),
    [registerBrc.success]: (state, { payload: { id, season, brcHub } }) => ({
      ...state,
      allBrcHubs: {
        ...state.allBrcHubs,
        [id]: state.allBrcHubs[id].map(s =>
          s.ID === season
            ? {
                ...s,
                brcHub
              }
            : s
        )
      }
    })
  },
  defaultState
);

export default reducer;
