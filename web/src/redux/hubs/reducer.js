import { createActions, handleActions, combineActions } from "redux-actions";

const defaultState = {
  myHubs: [],
  allHubs: []
};

export const { getMyHubs, getAllHubs } = createActions({
  GET_MY_HUBS: {
    REQUEST: () => ({}),
    SUCCESS: hubs => ({ hubs }),
    FAILURE: error => ({ error })
  },
  GET_ALL_HUBS: {
    REQUEST: () => ({}),
    SUCCESS: hubs => ({ hubs }),
    FAILURE: error => ({ error })
  }
});

const reducer = handleActions(
  {
    [getMyHubs.request]: state => {
      console.log("HELLO");
      return state;
    },
    [getMyHubs.success]: (state, { payload: { hubs } }) => ({
      ...state,
      myHubs: hubs
    }),
    [getAllHubs.success]: (state, { payload: { hubs } }) => ({
      ...state,
      allHubs: hubs
    })
  },
  defaultState
);

export default reducer;
