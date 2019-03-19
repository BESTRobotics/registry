import axios from "axios";

const url = process.env.REACT_APP_API_URL;

export function fetchHub(id, token) {
  return axios.get(`http://${url}/v1/hubs/${id}`, {
    token: token
  });
}

export function fetchHubs(ids, token) {
  const requests = ids.map(id => fetchHub(id, token));
  return axios
    .all(requests)
    .then(axios.spread((...responses) => responses.map(res => res.data)));
}

export function fetchTeam(id, token) {
  return axios.get(`http://${url}/v1/teams/${id}`, {
    token: token
  });
}

export function fetchTeams(ids, token) {
  const requests = ids.map(id => fetchTeam(id, token));
  return axios
    .all(requests)
    .then(axios.spread((...responses) => responses.map(res => res.data)));
}
