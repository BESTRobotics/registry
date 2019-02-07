import React from "react";
import { Image, Menu } from "semantic-ui-react";
import logo from "../assets/logo.jpg";
import { NavLink } from "react-router-dom";

const Topbar = () => {
  return (
    <Menu>
      <Menu.Item header fitted>
        <Image size="tiny" src={logo} />
      </Menu.Item>
      <Menu.Item as={NavLink} to="/hubs">
        Hubs
      </Menu.Item>
      <Menu.Item as={NavLink} to="/schools">
        Schools
      </Menu.Item>
      <Menu.Item as={NavLink} to="/teams">
        Teams
      </Menu.Item>
      <Menu.Item as={NavLink} to="/seasons">
        Seasons
      </Menu.Item>
      <Menu.Item as={NavLink} to="/users">
        Users
      </Menu.Item>
    </Menu>
  );
};

export default Topbar;
