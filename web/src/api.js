import axios from "axios";

const url = process.env.REACT_APP_API_URL;

export function fetchHub(id, token) {
  return axios.get(`http://${url}/v1/hubs/${id}`, {
    headers: { Authorization: token }
  });
}

export function fetchSeasons(token) {
  return axios
    .get(`http://${url}/v1/seasons`, {
      headers: { Authorization: token }
    })
    .then(s => s.data);
}

export function fetchBrcHubs(id, token) {
  return axios
    .get(`http://${url}/v1/hubs/${id}/brc`, {
      headers: { Authorization: token }
    })
    .then(h => h.data);
}

export function registerBrcHub(id, season, token) {
  return axios
    .post(
      `http://${url}/v1/hubs/${id}/brc/${season}`,
      {},
      {
        headers: { Authorization: token }
      }
    )
    .then(h => h.data);
}

export function fetchHubs(ids, token) {
  const requests = ids.map(id => fetchHub(id, token));
  return axios
    .all(requests)
    .then(axios.spread((...responses) => responses.map(res => res.data)));
}

export function fetchTeam(id, token) {
  return axios.get(`http://${url}/v1/teams/${id}`, {
    headers: { Authorization: token }
  });
}

export function fetchTeams(ids, token) {
  const requests = ids.map(id => fetchTeam(id, token));
  return axios
    .all(requests)
    .then(axios.spread((...responses) => responses.map(res => res.data)));
}
