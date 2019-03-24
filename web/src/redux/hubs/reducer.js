import { createActions, handleActions, combineActions } from "redux-actions";

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
    SUCCESS: (id, brcHub) => ({ id, brcHub }),
    FAILURE: error => ({ error })
  },
  REGISTER_BRC: {
    REQUEST: id => ({ id }),
    SUCCESS: (id, brcHub) => ({ id, brcHub }),
    FAILURE: error => ({ error })
  }
});

const reducer = handleActions(
  {
    [getMyHubs.request]: state => {
      return state;
    },
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
    [getBrcHub.success]: (state, { payload: { id, brcHub } }) => ({
      ...state,
      allBrcHubs: {
        ...state.allBrcHubs,
        [id]: brcHub
      }
    })
  },
  defaultState
);

export default reducer;
