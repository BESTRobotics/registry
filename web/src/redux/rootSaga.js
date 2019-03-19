import { all } from "redux-saga/effects";
import hubsSagas from "./hubs/sagas";
import teamsSagas from "./teams/sagas";

export default function* rootSaga() {
  yield all([...hubsSagas, ...teamsSagas]);
}
