import React from "react";
import Item from "./Item";
import NewEventForm from "./NewEventForm";

const Users = ({ token }) => {
  const fields = [
    { header: "Name", name: "Name", filter: true },
    { header: "Description", name: "Description", filter: true },
    { header: "Location", name: "Location", filter: true },
    { header: "Start", name: "StartTime", filter: false },
    { header: "End", name: "EndTime", filter: false },
    {
      header: "Hub",
      displayFn: event => (event.Hub ? event.Hub.Name : ""),
      filter: true
    }
  ];
  return (
    <Item
      itemName="Event"
      fields={fields}
      NewItemForm={NewEventForm}
      token={token}
      trashcan={null}
    />
  );
};

export default Users;
