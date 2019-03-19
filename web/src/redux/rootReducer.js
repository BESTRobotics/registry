import { combineReducers, createStore } from "redux";
import loginReducer from "./login/reducer";
import hubsReducer from "./hubs/reducer";
import teamsReducer from "./hubs/reducer";

export const rootReducer = combineReducers({
  loginReducer,
  hubsReducer,
  teamsReducer
});

const store = createStore(rootReducer);

store.subscribe(() => {
  const token = store.getState().loginReducer.token;
  if (token) {
    window.localStorage.setItem("token", token);
  } else {
    window.localStorage.removeItem("token");
  }
});

export default store;
