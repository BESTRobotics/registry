import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import { Button, Form } from "semantic-ui-react";
import { getAllHubs } from "../../redux/hubs/reducer";
import { registerNewTeam } from "../../redux/teams/reducer";

const NewTeam = ({ onDone, hubs, registerNewTeam, getAllHubs }) => {
  const [name, setName] = useState("");
  const [schoolName, setSchoolName] = useState("");
  const [schoolAddress, setSchoolAddress] = useState("");
  const [website, setWebsite] = useState("");
  const [hub, setHub] = useState(null);
  const [founded, setFounded] = useState("");

  useEffect(() => {
    (hubs && hubs.length) || getAllHubs();
  }, []);

  return (
    <React.Fragment>
      <Form
        onSubmit={() => {
          registerNewTeam({
            name,
            schoolName,
            schoolAddress,
            website,
            hub,
            founded
          });
          onDone(true);
        }}
      >
        <Form.Input
          label="Team Name"
          required
          value={name}
          onChange={(_, { value }) => setName(value)}
        />
        <Form.Input
          label="School Name"
          required
          value={schoolName}
          onChange={(_, { value }) => setSchoolName(value)}
        />
        <Form.TextArea
          label="School Address"
          required
          value={schoolAddress}
          onChange={(_, { value }) => setSchoolAddress(value)}
        />
        <Form.Input
          label="Website"
          type="url"
          value={website}
          onChange={(_, { value }) => setWebsite(value)}
        />
        <Form.Dropdown
          label="Home Hub"
          search
          required
          loading={!hubs.length}
          options={
            hubs &&
            hubs.map(h => ({
              text: `${h.Name} — ${h.Location}`,
              value: h.ID
            }))
          }
          selection
          value={hub}
          onChange={(_, { value }) => setHub(value)}
        />
        <Form.Input
          type="date"
          label="Founded"
          value={founded}
          onChange={(_, { value }) => setFounded(value)}
        />
        <Button color="green">Register Team</Button>
        <Button type="button" secondary onClick={() => onDone(false)}>Cancel</Button>
      </Form>
    </React.Fragment>
  );
};

const mapStateToProps = ({ hubsReducer }) => ({
  hubs: hubsReducer.allHubs
});

const mapDispatchToProps = {
  getAllHubs: () => getAllHubs.request(),
  registerNewTeam: team => registerNewTeam.request(team)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NewTeam);
