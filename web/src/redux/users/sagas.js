import { takeEvery, select, call, put, all } from "redux-saga/effects";
import { updateMyProfile, getMyProfile, getMyStudents, updateMyStudent, registerStudents, getStudentRegistrations } from "./reducer";

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

function* updateMyStudentSaga({ payload: { student } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  const id = yield select(({ loginReducer }) => loginReducer.id);
  const studentToUpdate = {
    ID: student.id,
    FirstName: student.firstName,
    LastName: student.lastName,
    EMail: student.email,
    Race: student.race,
    Gender: student.gender,
  };

  try {
    const returnStudent = student.id ?
      yield call(
        api.updateStudent,
        id,
        studentToUpdate,
        token
      )
      :
      yield call(
        api.addStudent,
        id,
        studentToUpdate,
        token
      );
    yield put({
      type: updateMyStudent.success,
      payload: { student: returnStudent }
    });
  } catch (err) {
    console.error(err);
    yield put({ type: updateMyStudent.failure, payload: { error: err } });
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

function* registerStudentsSaga({ payload: { students, team, secret } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    yield all(students.map(s =>
      call(
        api.registerStudent,
        s,
        team,
        secret,
        token
      )
    ));
    yield put({ type: registerStudents.success, payload: {} });
  } catch (err) {
    yield put({ type: registerStudents.failure, payload: { error: err } });
  }
}

function* getStudentRegistrationsSaga({ payload: { students } }) {
  const token = yield select(({ loginReducer }) => loginReducer.token);
  try {
    let registrations = yield all(students.map(s =>
      call(
        api.teamsByStudent,
        s,
        token
      )
    ));
    yield put({ type: getStudentRegistrations.success, payload: { registrations } });
  } catch (err) {
    yield put({ type: registerStudents.failure, payload: { error: err } });
  }
}

export default [
  takeEvery(getMyProfile.request, getMyProfileSaga),
  takeEvery(updateMyProfile.request, updateMyProfileSaga),
  takeEvery(getMyStudents.request, getMyStudentsSaga),
  takeEvery(updateMyStudent.request, updateMyStudentSaga),
  takeEvery(registerStudents.request, registerStudentsSaga),
  takeEvery(getStudentRegistrations.request, getStudentRegistrationsSaga)
];
