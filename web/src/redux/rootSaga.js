import { spawn, call } from "redux-saga/effects";
import hubsSagas from "./hubs/sagas";
import teamsSagas from "./teams/sagas";

export function* rootSaga() {
  const sagas = [...hubsSagas, ...teamsSagas];

  yield sagas.map(saga =>
    spawn(function*() {
      while (true) {
        try {
          yield call(saga);
          break;
        } catch (e) {
          console.log(e);
        }
      }
    })
  );
}
