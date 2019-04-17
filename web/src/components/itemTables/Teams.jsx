import React from "react";
import Item from "./Item";
import NewTeamForm from "../itemForms/NewTeamForm";

const Teams = ({ token }) => {
  const trunc = (text, max) =>
    text.substr(0, max - 1) + (text.length > max ? "\u2026" : "");
  const fields = [
    { header: "Name", name: "StaticName", filter: true },
    { header: "Hub", displayFn: team => team.HomeHub.Name, filter: true },
    {
      header: "Coach",
      displayFn: team =>
        team.Coach
          ? trunc(
              team.Coach.map(m => `${m.FirstName} ${m.LastName}`).join(", "),
              25
            )
          : "",
      filter: false
    },
    {
      header: "Mentors",
      displayFn: team =>
        team.Mentors
          ? trunc(
              team.Mentors.map(m => `${m.FirstName} ${m.LastName}`).join(", "),
              25
            )
          : "",
      filter: false
    },
    { header: "School", name: "SchoolName", filter: true },
    { header: "Address", name: "SchoolAddress", filter: true },
    { header: "Website", name: "Website", filter: false },
    {
      header: "Founded",
      displayFn: team => (team.Founded ? team.Founded.substring(0, 10) : ""),
      filter: true
    }
  ];
  return (
    <Item
      itemName="Team"
      fields={fields}
      NewItemForm={NewTeamForm}
      token={token}
      trashcan="deactivate"
    />
  );
};

export default Teams;
