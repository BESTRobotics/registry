import { takeEvery, select, all, call, put } from "redux-saga/effects";
import { getBrcHub, getAllHubs, getMyHubs, registerBrc } from "./reducer";
import * as api from "../../api";

function* getAllHubsSaga(action) {
  console.log("unimplemented");
  //do something
}

function* getBrcHubSaga({ payload: { id } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const [seasons, brcHubs] = yield all([
      call(api.fetchSeasons, token),
      call(api.fetchBrcHubs, id, token)
    ]);
    yield put({ type: getBrcHub.success, payload: { brcHubs, seasons, id } });
  } catch (err) {
    console.error(err);
    yield put({ type: getBrcHub.failure, payload: { error: err } });
  }
}

function* registerBrcHubSaga({ payload: { id, season } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const brcHub = yield call(api.registerBrcHub, id, season, token);
    yield put({ type: registerBrc.success, payload: { id, season, brcHub } });
  } catch (err) {
    console.error(err);
    yield put({ type: registerBrc.failure, payload: { error: err } });
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
