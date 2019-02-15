import React from "react";
import Item from "./Item";
import NewTeamForm from "./NewTeamForm";

const Teams = () => {
  const fields = [
    { header: "Name", name: "StaticName", filter: true },
    { header: "Hub", displayFn: team => team.HomeHub.Name, filter: true },
    {
      header: "Coach",
      displayFn: team => `${team.Coach.FirstName} ${team.Coach.LastName}`,
      filter: false
    },
    { header: "Mentors", displayFn: team => "Coming Soon", filter: false },
    { header: "School", displayFn: team => team.School.Name, filter: true },
    {
      header: "Founded",
      displayFn: team => (team.Founded ? team.Founded.substring(0, 10) : ""),
      filter: true
    }
  ];
  return <Item itemName="Team" fields={fields} NewItemForm={NewTeamForm} />;
};

export default Teams;
