import { takeEvery, select, call, put } from "redux-saga/effects";
import { getAllHubs, getMyHubs } from "./reducer";
import * as api from "../../api";

function* getAllHubsSaga(action) {
  console.log("unimplemented");
  //do something
}

function* getMyHubsSaga(action) {
  console.log("implemented");
  const token = yield select(({ loginReducer }) => loginReducer.token);
  const hubIds = yield select(({ loginReducer }) => loginReducer.hubs);
  try {
    const hubs = yield call(api.fetchHubs, hubIds, token);
    yield put({ type: getMyHubs.success, payload: { hubs } });
  } catch (err) {
    console.error(err);
    yield put({ type: getMyHubs.failure, payload: { error: err } });
  }
}

// use them in parallel
export default [
  takeEvery(getAllHubs.request, getAllHubsSaga),
  takeEvery(getMyHubs.request, getMyHubsSaga)
];
