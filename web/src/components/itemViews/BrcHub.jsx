import React, { useEffect } from "react";
import { connect } from "react-redux";
import { getBrcHub, registerBrc } from "../../redux/hubs/reducer";
import { Button, Icon, Header } from "semantic-ui-react";

const BrcDescription = ({ brcHub, register }) => {
  console.log(brcHub);
  return (
    <>
      <Header>{`${brcHub.brcHub.Hub.Name} | ${
        brcHub.brcHub.Season.Name
      }`}</Header>
      {brcHub.Message ? (
        <Button onClick={register}>Register for current season</Button>
      ) : (
        <div>
          {brcHub.brcHub.Meta.BRIApproved ? (
            "Registration Approved"
          ) : (
            <>
              <Icon name="warning" /> Registration not yet approved
            </>
          )}
        </div>
      )}
    </>
  );
};

const BrcHub = ({
  allBrcHubs,
  getBrcHub,
  registerBrc,
  match: {
    params: { id, season }
  }
}) => {
  useEffect(() => {
    (allBrcHubs && allBrcHubs[id]) || getBrcHub(id);
  }, []);

  console.log(allBrcHubs[id] && allBrcHubs[id].map(s => s.ID));
  console.log(season);

  return (
    <div>
      {allBrcHubs && allBrcHubs[id] ? (
        <BrcDescription
          brcHub={allBrcHubs[id].find(s => String(s.ID) === season)}
          register={() => registerBrc(id)}
        />
      ) : (
        "Loading Enrollment ..."
      )}
    </div>
  );
};

const mapStateToProps = ({ hubsReducer }) => ({
  allBrcHubs: hubsReducer.allBrcHubs
});

const mapDispatchToProps = {
  getBrcHub: id => getBrcHub.request(id),
  registerBrc: id => registerBrc.request(id)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(BrcHub);
