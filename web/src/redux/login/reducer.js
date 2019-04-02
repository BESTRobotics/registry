import { createActions, handleActions, combineActions } from "redux-actions";
import jwt_decode from "jwt-decode";
import { getAllHubs, getBrcHub, getMyHubs } from "../hubs/reducer";
import { getAllTeams, getMyTeams } from "../teams/reducer";

const initialToken = window.localStorage.getItem("token") || null;
const initialDecodedToken = initialToken && jwt_decode(initialToken);

const defaultState = {
  token: initialToken,
  superAdmin:
    initialDecodedToken &&
    initialDecodedToken.User.Capabilities &&
    initialDecodedToken.User.Capabilities.includes(0),
  hubs: initialDecodedToken ? initialDecodedToken.Hubs : [],
  teams: initialDecodedToken ? initialDecodedToken.Teams : [],
  message: null
};

export const { logout, setToken } = createActions({
  LOGOUT: () => ({}),
  SET_TOKEN: token => ({ token })
});

const reducer = handleActions(
  {
    [logout]: () => ({ token: null, superAdmin: false, hubs: [], teams: [] }),
    [setToken]: (state, { payload: { token } }) => {
      console.log(token);
      const decoded = jwt_decode(token);
      console.log(decoded);
      return {
        ...state,
        token: token,
        superAdmin:
          decoded.User.Capabilities && decoded.User.Capabilities.includes(0),
        hubs: decoded.Hubs,
        teams: decoded.Teams
      };
    },
    [combineActions(
      getAllHubs.failure,
      getBrcHub.failure,
      getMyHubs.failure,
      getAllTeams.failure,
      getMyTeams.failure
    )]: (state, { payload: { error } }) => {
      if (error.response && error.response.message) {
        return {
          ...state,
          message: {
            header: "An error occured",
            icon: "warning circle",
            content: error.response.message,
            error: true
          }
        };
      }
      return {
        ...state,
        message: {
          header: "An error occured",
          icon: "warning circle",
          content: error.message || "Connection Error",
          error: true
        }
      };
    }
  },
  defaultState
);

export default reducer;
