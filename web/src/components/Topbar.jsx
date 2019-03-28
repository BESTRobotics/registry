import React from "react";
import { Image, Menu } from "semantic-ui-react";
import logo from "../assets/logo.jpg";
import { NavLink } from "react-router-dom";
import { connect } from "react-redux";
import { logout as callLogout } from "../redux/login/reducer";

const Topbar = ({ logout, superAdmin, hubs, teams }) => {
  return (
    <Menu>
      <Menu.Item as={NavLink} to="/" header fitted>
        <Image size="tiny" src={logo} />
      </Menu.Item>
      {hubs ? (
        <Menu.Item as={NavLink} to="/hubs">
          My Hubs
        </Menu.Item>
      ) : null}
      {teams ? (
        <Menu.Item as={NavLink} to="/teams">
          My Teams
        </Menu.Item>
      ) : null}
      {superAdmin ? (
        <>
          <Menu.Item as={NavLink} to="/admin/hubs">
            Hubs
          </Menu.Item>
          <Menu.Item as={NavLink} to="/admin/teams">
            Teams
          </Menu.Item>
          <Menu.Item as={NavLink} to="/admin/seasons">
            Seasons
          </Menu.Item>
          <Menu.Item as={NavLink} to="/admin/users">
            Users
          </Menu.Item>
          <Menu.Item as={NavLink} to="/admin/events">
            Events
          </Menu.Item>
        </>
      ) : null}
      <Menu.Item position="right" onClick={logout}>
        Logout
      </Menu.Item>
    </Menu>
  );
};

const mapStateToProps = ({ loginReducer }) => ({
  superAdmin: loginReducer.superAdmin,
  hubs: loginReducer.hubs && loginReducer.hubs.length,
  teams: loginReducer.teams && loginReducer.teams.length
});

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(callLogout())
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Topbar);
