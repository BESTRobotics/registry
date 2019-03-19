import { createActions, handleActions, combineActions } from "redux-actions";
import jwt_decode from "jwt-decode";

const initialToken = window.localStorage.getItem("token") || null;
const initialDecodedToken = initialToken && jwt_decode(initialToken);

const defaultState = {
  token: initialToken,
  superAdmin:
    initialDecodedToken &&
    initialDecodedToken.User.Capabilities &&
    initialDecodedToken.User.Capabilities.includes(0),
  hubs: initialDecodedToken ? initialDecodedToken.Hubs : [],
  teams: initialDecodedToken ? initialDecodedToken.Teams : []
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
    }
  },
  defaultState
);

export default reducer;
