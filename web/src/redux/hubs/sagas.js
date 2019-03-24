import { takeEvery, select, call, put } from "redux-saga/effects";
import { getBrcHub, getAllHubs, getMyHubs, registerBrc } from "./reducer";
import * as api from "../../api";

function* getAllHubsSaga(action) {
  console.log("unimplemented");
  //do something
}

function* getBrcHubSaga({ payload: { id } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const seasons = yield call(api.fetchSeasons, token);
    const brcHub = yield call(api.fetchBrcHub, id, seasons[0].ID, token);
    yield put({ type: getBrcHub.success, payload: { brcHub, id } });
  } catch (err) {
    console.error(err);
    yield put({ type: getBrcHub.failure, payload: { error: err } });
  }
}

function* registerBrcHubSaga({ payload: { id } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const seasons = yield call(api.fetchSeasons, token);
    const brcHub = yield call(api.registerBrcHub, id, seasons[0].ID, token);
    yield put({ type: getBrcHub.success, payload: { brcHub, id } });
  } catch (err) {
    console.error(err);
    yield put({ type: getBrcHub.failure, payload: { error: err } });
  }
}

function* getMyHubsSaga(action) {
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
  takeEvery(getMyHubs.request, getMyHubsSaga),
  takeEvery(getBrcHub.request, getBrcHubSaga),
  takeEvery(registerBrc.request, registerBrcHubSaga)
];
