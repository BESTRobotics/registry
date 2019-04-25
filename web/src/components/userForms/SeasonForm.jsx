import React, { useState } from "react";
import { connect } from "react-redux";
import { Button, Form } from "semantic-ui-react";
import PropTypes from "prop-types";
import { saveSeason } from "../../redux/hubs/reducer";

const SeasonForm = ({ existingItem, saveSeason }) => {
  const season = existingItem;
  const [name, setName] = useState(season ? season.Name : "");
  const [state, setState] = useState(season ? season.State : "Closed");
  const [program, setProgram] = useState(season ? season.Program : 0);
  const [id, setId] = useState(season ? season.ID : null);

  return (
    <Form onSubmit={() => saveSeason({ id, name, state, program })}>
      <Form.Input
        label="Name"
        value={name}
        onChange={(_, { value }) => setName(value)}
      />
      <Form.Checkbox
        label="Open"
        checked={state === "Open"}
        onChange={() => setState(state !== "Open" ? "Open" : "Closed")}
      />
      <Form.Input
        label="Program"
        type="number"
        step="1"
        value={program}
        onChange={(_, { value }) => setProgram(value)}
      />
      <Button color="green">
        {id && id !== "" ? "Update Season" : "Add Season"}
      </Button>
    </Form>
  );
};

const mapStateToProps = ({}) => ({});

const mapDispatchToProps = {
  saveSeason: season => saveSeason.request(season)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SeasonForm);
