import { combineReducers, createStore, applyMiddleware } from "redux";
import createSagaMiddleware from "redux-saga";
import loginReducer from "./login/reducer";
import hubsReducer from "./hubs/reducer";
import teamsReducer from "./teams/reducer";
import usersReducer from "./users/reducer";
import rootSaga from "./rootSaga";

export const rootReducer = combineReducers({
  hubsReducer,
  teamsReducer,
  usersReducer,
  loginReducer
});

const sagaMiddleware = createSagaMiddleware();

const store = createStore(rootReducer, applyMiddleware(sagaMiddleware));

store.subscribe(() => {
  const token = store.getState().loginReducer.token;
  if (token) {
    window.localStorage.setItem("token", token);
  } else {
    window.localStorage.removeItem("token");
  }
});

sagaMiddleware.run(rootSaga);

export default store;
