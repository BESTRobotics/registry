import React from "react";
import Item from "./Item";
import NewHubForm from "./NewHubForm";

const Hubs = ({ token }) => {
  const fields = [
    { header: "Name", name: "Name", filter: true },
    { header: "Location", name: "Location", filter: true },
    {
      header: "Director",
      displayFn: h => `${h.Director.FirstName} ${h.Director.LastName}`,
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
    />
  );
};

export default Hubs;
