import { takeEvery, select, call, put } from "redux-saga/effects";
import { getAllTeams, getMyTeams } from "./reducer";
import * as api from "../../api";

function* getAllTeamsSaga(action) {
  console.log("unimplemented");
  //do something
}

function* getMyTeamsSaga(action) {
  console.log("implemented");
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

// use them in parallel
export default [
  takeEvery(getAllTeams.request, getAllTeamsSaga),
  takeEvery(getMyTeams.request, getMyTeamsSaga)
];
