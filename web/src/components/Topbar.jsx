import React from "react";
import { Image, Menu } from "semantic-ui-react";
import logo from "../assets/logo.jpg";
import { NavLink } from "react-router-dom";
import { connect } from "react-redux";
import { logout as callLogout } from "../redux/login/reducer";

const Topbar = ({ logout }) => {
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
      <Menu.Item as={NavLink} to="/events">
        Events
      </Menu.Item>
      <Menu.Item position="right" onClick={logout}>
        Logout
      </Menu.Item>
    </Menu>
  );
};

const mapStateToProps = ({ loginReducer }) => ({
  token: loginReducer.token,
  superAdmin: loginReducer.superAdmin,
  hubs: loginReducer.hubs,
  teams: loginReducer.teams
});

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(callLogout())
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Topbar);
