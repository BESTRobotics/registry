import React from "react";
import Item from "./Item";
import NewSchoolForm from "../itemForms/NewSchoolForm";

const Schools = ({ token }) => {
  const fields = [
    { header: "Name", name: "Name", filter: true },
    { header: "Address", name: "Address", filter: true },
    { header: "Website", name: "Website", filter: false }
  ];
  return (
    <Item
      itemName="School"
      fields={fields}
      NewItemForm={NewSchoolForm}
      token={token}
      trashcan={null}
    />
  );
};

export default Schools;
