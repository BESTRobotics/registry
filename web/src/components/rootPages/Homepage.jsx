import React from "react";
import { connect } from "react-redux";
import Hubs from "../itemViews/Hubs";
import Teams from "../itemViews/Teams";
import NewUser from "./NewUser";

const mapStateToProps = ({ loginReducer }) => ({
  hubsLength: loginReducer.hubs ? loginReducer.hubs.length : 0,
  teamsLength: loginReducer.teams ? loginReducer.teams.length : 0
});

const InnerHomepage = connect(mapStateToProps)(
  ({ hubsLength, teamsLength }) => {
    return (
      <React.Fragment>
        {hubsLength ? <Hubs /> : null}
        {teamsLength ? <Teams /> : null}
        {!hubsLength && !teamsLength && <NewUser />};
      </React.Fragment>
    );
  }
);

export default InnerHomepage;
