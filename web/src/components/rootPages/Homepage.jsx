import React from "react";
import { connect } from "react-redux";
import Hubs from "../itemViews/Hubs";
import Teams from "../itemViews/Teams";
import NewUser from "./NewUser";
import { Header } from "semantic-ui-react";
import Adminpage from "./Adminpage";

const mapStateToProps = ({ loginReducer }) => ({
  superAdmin: loginReducer.superAdmin,
  hubsLength: loginReducer.hubs ? loginReducer.hubs.length : 0,
  teamsLength: loginReducer.teams ? loginReducer.teams.length : 0
});

const InnerHomepage = connect(mapStateToProps)(
  ({ hubsLength, teamsLength, superAdmin }) => {
    return (
      <React.Fragment>
        {superAdmin && !hubsLength && !teamsLength ? <Adminpage /> : null}
        {hubsLength === 1 && teamsLength && <Header>My Hub</Header>}
        {hubsLength ? <Hubs /> : null}
        {teamsLength === 1 && hubsLength && <Header>My Team</Header>}
        {teamsLength ? <Teams /> : null}
        {!hubsLength && !teamsLength && <NewUser />};
      </React.Fragment>
    );
  }
);

export default InnerHomepage;
