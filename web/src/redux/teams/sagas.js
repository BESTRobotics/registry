import { takeEvery, select, all, call, put } from "redux-saga/effects";
import {
  getAllTeams,
  getMyTeams,
  getBrcTeam,
  registerBrcTeam,
  registerNewTeam
} from "./reducer";
import * as api from "../../api";

function* getAllTeamsSaga(action) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const teams = yield call(api.fetchAllTeams, token);
    yield put({ type: getAllTeams.success, payload: { teams } });
  } catch (err) {
    console.error(err);
    yield put({ type: getAllTeams.failure, payload: { error: err } });
  }
}

function* getBrcTeamSaga({ payload: { id } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const [seasons, brcTeams] = yield all([
      call(api.fetchSeasons, token),
      call(api.fetchBrcTeams, id, token)
    ]);
    yield put({ type: getBrcTeam.success, payload: { brcTeams, seasons, id } });
  } catch (err) {
    console.error(err);
    yield put({ type: getBrcTeam.failure, payload: { error: err } });
  }
}

function* getMyTeamsSaga(action) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  const teamIds = yield select(({ loginReducer }) => loginReducer.teams);
  try {
    const teams = yield call(api.fetchTeams, teamIds, token);
    yield put({ type: getMyTeams.success, payload: { teams } });
  } catch (err) {
    console.error(err);
    yield put({ type: getMyTeams.failure, payload: { error: err } });
  }
}

function* registerBrcTeamSaga({ payload: { id, season } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const brcHub = yield call(api.registerBrcTeam, id, season, token);
    yield put({
      type: registerBrcTeam.success,
      payload: { id, season, brcHub }
    });
  } catch (err) {
    console.error(err);
    yield put({ type: registerBrcTeam.failure, payload: { error: err } });
  }
}

function* registerNewTeamSaga({ payload: { team } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    const teamToRegister = {
      StaticName: team.name,
      SchoolName: team.schoolName,
      SchoolAddress: team.schoolAddress,
      Website: team.website,
      Founded: team.founded ? new Date(team.founded).toISOString() : null
    };
    const newTeam = yield call(api.registerNewTeam, teamToRegister, token);
    console.log(team.hub);
    if (team.hub) {
      yield call(api.setTeamHub, newTeam.id, team.hub, token);
    }
    yield put({
      type: registerNewTeam.success,
      payload: { team: newTeam }
    });
  } catch (err) {
    console.error(err);
    yield put({ type: registerNewTeam.failure, payload: { error: err } });
  }
}

// use them in parallel
export default [
  takeEvery(getAllTeams.request, getAllTeamsSaga),
  takeEvery(getMyTeams.request, getMyTeamsSaga),
  takeEvery(getBrcTeam.request, getBrcTeamSaga),
  takeEvery(registerBrcTeam.request, registerBrcTeamSaga),
  takeEvery(registerNewTeam.request, registerNewTeamSaga)
];
