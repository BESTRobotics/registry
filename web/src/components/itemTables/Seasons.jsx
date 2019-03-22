import React from "react";
import Item from "./Item";
import NewSeasonForm from "../itemForms/NewSeasonForm";

const Hubs = ({ token }) => {
  const fields = [{ header: "Name", name: "Name", filter: true }];
  return (
    <Item
      itemName="Season"
      fields={fields}
      NewItemForm={NewSeasonForm}
      token={token}
      trashcan="archive"
    />
  );
};

export default Hubs;
