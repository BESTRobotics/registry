import React, { Component } from "react";
import axios from "axios";

export default class Hubs extends Component {
  constructor(props) {
    super(props);
    this.state = {
      hubs: []
    };
  }

  componentDidMount() {
    axios
      .get(`http://${process.env.REACT_APP_API_URL}/v1/hubs`)
      .then(response => {
        console.log(response.data);
      })
      .catch(e => console.log(e));
  }
  render() {
    return (
      <div>
        Hello would you like a hub? {process.env.REACT_APP_API_URL} certainly
        does
        <br />
        {`http://${process.env.REACT_APP_API_URL}/v1/hubs`}
      </div>
    );
  }
}
