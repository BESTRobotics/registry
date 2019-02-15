import React from "react";
import Item from "./Item";
import NewSchoolForm from "./NewSchoolForm";

const Schools = () => {
  const fields = [
    { header: "Name", name: "Name", filter: true },
    { header: "Address", name: "Address", filter: true },
    { header: "Website", name: "Website", filter: false }
  ];
  return <Item itemName="School" fields={fields} NewItemForm={NewSchoolForm} />;
};

export default Schools;
