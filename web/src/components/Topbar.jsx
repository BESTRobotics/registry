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
      {superAdmin || hubs ? (
        <Menu.Item as={NavLink} to="/hubs">
          Hubs
        </Menu.Item>
      ) : null}
      {superAdmin || teams ? (
        <Menu.Item as={NavLink} to="/teams">
          Teams
        </Menu.Item>
      ) : null}
      {superAdmin ? (
        <>
          <Menu.Item as={NavLink} to="/schools">
            Schools
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
