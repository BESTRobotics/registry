import axios from "axios";

const url = process.env.REACT_APP_API_URL;

export function getHub(id, token) {
  axios.get(`http://${url}/v1/hubs/${id}`, {
    token: token
  });
}

export function getHubs(ids, token) {
  const requests = ids.map(id => getHub(id, token));
  return axios.all(requests).then(
    axios.spread((...responses) => {
      responses.map(res => res.data);
    })
  );
}

export function getTeam(id, token) {
  axios.get(`http://${url}/v1/teams/${id}`, {
    token: token
  });
}

export function getTeams(ids, token) {
  const requests = ids.map(id => getTeam(id, token));
  return axios.all(requests).then(
    axios.spread((...responses) => {
      responses.map(res => res.data);
    })
  );
}
