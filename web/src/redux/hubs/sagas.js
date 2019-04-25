import { takeEvery, select, all, call, put } from "redux-saga/effects";
import {
  getBrcHub,
  getSeasons,
  getSeasonBrcHubs,
  getAllHubs,
  getMyHubs,
  registerBrcHub,
  approveBrcHub,
  saveSeason
} from "./reducer";
import * as api from "../../api";

function* getAllHubsSaga(action) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const hubs = yield call(api.fetchAllHubs, token);
    yield put({ type: getAllHubs.success, payload: { hubs } });
  } catch (err) {
    console.error(err);
    yield put({ type: getAllHubs.failure, payload: { error: err } });
  }
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

function* getSeasonsSaga() {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const seasons = yield call(api.fetchSeasons, token);
    yield put({ type: getSeasons.success, payload: { seasons } });
  } catch (err) {
    console.error(err);
    yield put({ type: getSeasons.failure, payload: { error: err } });
  }
}

function* getSeasonBrcHubsSaga({ payload: { season } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const brcHubs = yield call(api.fetchSeasonBrcHubs, season, token);
    yield put({ type: getSeasonBrcHubs.success, payload: { brcHubs, season } });
  } catch (err) {
    console.error(err);
    yield put({ type: getSeasonBrcHubs.failure, payload: { error: err } });
  }
}

function* registerBrcHubSaga({ payload: { id, season } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const brcHub = yield call(api.registerBrcHub, id, season, token);
    yield put({
      type: registerBrcHub.success,
      payload: { id, season, brcHub }
    });
  } catch (err) {
    console.error(err);
    yield put({ type: registerBrcHub.failure, payload: { error: err } });
  }
}

function* approveBrcHubSaga({ payload: { hubid, brchubid, season } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    yield call(api.approveBrcHub, hubid, brchubid, season, token);
    yield put({
      type: approveBrcHub.success,
      payload: { id: brchubid, season }
    });
  } catch (err) {
    console.error(err);
    yield put({ type: approveBrcHub.failure, payload: { error: err } });
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

function* saveSeasonSaga({ payload: { season } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  let id;
  let newSeason;
  try {
    if (season.ID && season.ID !== "") {
      id = season.ID;
      newSeason = season;
      yield call(api.updateSeason, season.ID, season, token);
    } else {
      newSeason = yield call(api.newSeason, season, token);
      console.log(newSeason);
      id = newSeason.ID;
    }
    yield put({
      type: saveSeason.success,
      payload: { id, season: newSeason }
    });
  } catch (err) {
    console.error(err);
    yield put({ type: saveSeason.failure, payload: { error: err } });
  }
}

// use them in parallel
export default [
  takeEvery(getAllHubs.request, getAllHubsSaga),
  takeEvery(getMyHubs.request, getMyHubsSaga),
  takeEvery(getBrcHub.request, getBrcHubSaga),
  takeEvery(getSeasonBrcHubs.request, getSeasonBrcHubsSaga),
  takeEvery(approveBrcHub.request, approveBrcHubSaga),
  takeEvery(registerBrcHub.request, registerBrcHubSaga),
  takeEvery(getSeasons.request, getSeasonsSaga),
  takeEvery(saveSeason.request, saveSeasonSaga)
];
