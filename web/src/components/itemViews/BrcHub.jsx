import React, { useEffect } from "react";
import { connect } from "react-redux";
import { getBrcHub, registerBrc } from "../../redux/hubs/reducer";
import { Button, Icon } from "semantic-ui-react";

const BrcDescription = ({ brcHub, register }) => {
  console.log(brcHub);
  return brcHub.Message ? (
    <Button onClick={register}>Register for current season</Button>
  ) : (
    <div>
      {brcHub.brcHub.Meta.BRIApproved || (
        <div>
          <Icon name="warning" /> Registration not yet approved
        </div>
      )}
    </div>
  );
};

const BrcHub = ({
  allBrcHubs,
  getBrcHub,
  registerBrc,
  match: {
    params: { id }
  }
}) => {
  useEffect(() => {
    (allBrcHubs && allBrcHubs[id]) || getBrcHub(id);
  }, []);

  return (
    <div>
      {allBrcHubs && allBrcHubs[id] ? (
        <BrcDescription
          brcHub={allBrcHubs[id]}
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
