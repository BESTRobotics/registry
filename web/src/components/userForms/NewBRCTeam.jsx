import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import { Button, Form } from "semantic-ui-react";
import { getAllHubs } from "../../redux/hubs/reducer";
import { registerBrcTeam } from "../../redux/teams/reducer";

const NewBRCTeam = ({ onDone, team, season, registerBrcTeam}) => {
  const [state, setState] = useState("");
  const [uil, setUil] = useState("");
  const [joinKey, setJoinKey] = useState("");

  return (
    <React.Fragment>
      <Form
        onSubmit={() => {
          registerBrcTeam(team, season, {
            state,
            uil: uil === "" ? undefined : uil,
            joinKey,
          });
          onDone(true);
        }}
      >
        <Form.Input
          label="State"
          required
          value={state}
          onChange={(_, { value }) => setState(value)}
        />
        <Form.Input
          label="UIL Deivision"
          value={uil}
          onChange={(_, { value }) => setUil(value)}
        />
        <Form.Input
          label="Secret Join Key"
          required
          value={joinKey}
          onChange={(_, { value }) => setJoinKey(value)}
        />
        <Button color="green">Register Team</Button>
        <Button type="button" secondary onClick={() => onDone(false)}>Cancel</Button>
      </Form>
    </React.Fragment>
  );
};

const mapStateToProps = ({}) => ({
});

const mapDispatchToProps = {
  getAllHubs: () => getAllHubs.request(),
  registerBrcTeam: (team, season, brcTeam) => registerBrcTeam.request(team, season, brcTeam)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewBRCTeam);
