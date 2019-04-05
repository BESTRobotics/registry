import React from "react";
import Item from "./Item";
import NewSeasonForm from "../itemForms/NewSeasonForm";

const Seasons = ({ token }) => {
  const fields = [
    { header: "Name", name: "Name", filter: true },
    { header: "State", name: "State" },
    { header: "Program", name: "Program" }
  ];
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

export default Seasons;
