import React from "react";
import Item from "./Item";
import NewUserForm from "../itemForms/NewUserForm";

const Users = ({ token }) => {
  const fields = [
    // {
    //   header: "Name",
    //   displayFn: user => `${user.FirstName} ${user.LastName}`,
    //   filter: true
    // },
    { header: "Username", name: "Username", filter: true },
    { header: "Email", name: "EMail", filter: true }
    // {
    //   header: "Type",
    //   name: "Type",
    //   filter: true
    // },
    // {
    //   header: "Birthday",
    //   displayFn: user =>
    //     user.Birthdate ? user.Birthdate.substring(0, 10) : "",
    //   filter: false
    // }
  ];
  return (
    <Item
      itemName="User"
      fields={fields}
      NewItemForm={NewUserForm}
      token={token}
      trashcan={null}
    />
  );
};

export default Users;
