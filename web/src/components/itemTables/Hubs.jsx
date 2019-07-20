import React from "react";
import Item from "./Item";
import NewHubForm from "../itemForms/NewHubForm";
import { connect } from "react-redux";

const Hubs = ({ token }) => {
  const fields = [
    { header: "Name", name: "Name", filter: true },
    { header: "Location", name: "Location", filter: true },
    {
      header: "Director",
      displayFn: h =>
      h.Director ? `${h.Director.FirstName} ${h.Director.LastName} <${h.Director.EMail}>` : "",
      filter: false
    },
    {
      header: "Description",
      name: "Description",
      filter: false
    }
  ];
  return (
    <Item
      itemName="Hub"
      fields={fields}
      NewItemForm={NewHubForm}
      token={token}
      trashcan="deactivate"
    />
  );
};

const mapStateToProps = ({ loginReducer }) => ({ token: loginReducer.token });

export default connect(mapStateToProps)(Hubs);
