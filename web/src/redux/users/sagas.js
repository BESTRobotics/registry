import { takeEvery, select, call, put } from "redux-saga/effects";
import { updateMyProfile, getMyProfile, getMyStudents } from "./reducer";

import * as api from "../../api";

function* getMyStudentsSaga(action) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  const id = yield select(({ loginReducer }) => loginReducer.id);
  try {
    const students = yield call(api.fetchStudents, id, token);
    yield put({ type: getMyStudents.success, payload: { students } });
  } catch (err) {
    console.error(err);
    yield put({ type: getMyStudents.failure, payload: { error: err } });
  }
}

function* getMyProfileSaga(action) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  const id = yield select(({ loginReducer }) => loginReducer.id);
  try {
    const profile = yield call(api.fetchProfile, id, token);
    yield put({ type: getMyProfile.success, payload: { profile } });
  } catch (err) {
    console.error(err);
    yield put({ type: getMyProfile.failure, payload: { error: err } });
  }
}

function* updateMyProfileSaga({ payload: { profile } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  const id = yield select(({ loginReducer }) => loginReducer.id);
  const profileToUpdate = {
    FirstName: profile.firstName,
    LastName: profile.lastName,
    Birthdate: profile.birthdate
      ? new Date(profile.birthdate).toISOString()
      : null
  };

  try {
    const returnProfile = yield call(
      api.updateProfile,
      id,
      profileToUpdate,
      token
    );
    yield put({
      type: getMyProfile.success,
      payload: { profile: returnProfile }
    });
  } catch (err) {
    console.error(err);
    yield put({ type: getMyProfile.failure, payload: { error: err } });
  }
}
export default [
  takeEvery(getMyProfile.request, getMyProfileSaga),
  takeEvery(updateMyProfile.request, updateMyProfileSaga),
  takeEvery(getMyStudents.request, getMyStudentsSaga)
];
