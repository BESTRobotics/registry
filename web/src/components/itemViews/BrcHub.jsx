import React, { useEffect } from "react";
import { connect } from "react-redux";
import { getBrcHub, registerBrcHub } from "../../redux/hubs/reducer";
import {
  Button,
  Icon,
  Header,
  Card,
  Grid,
  Divider,
  List
} from "semantic-ui-react";
import FakeItemGroup from "./FakeItemGroup";

const BrcDescription = ({ brcHub, register, id, season }) => {
  return brcHub.brcHub ? (
    <Card fluid>
      <Card.Content>
        <Card.Header>{`${brcHub.brcHub.Hub.Name}  |  ${
          brcHub.brcHub.Season.Name
        }`}</Card.Header>
        <Card.Meta>
          {brcHub.brcHub && brcHub.brcHub.Meta.BRIApproved ? (
            "Registration Approved"
          ) : (
            <>
              <Icon name="warning" /> Registration not yet approved
            </>
          )}
        </Card.Meta>
        <Card.Description>
          <Divider horizontal>
            <Header as="h4">
              <Icon name="calendar" />
              Events
            </Header>
          </Divider>
          <List divided verticalAlign="middle">
            {brcHub.brcHub.events
              ? brcHub.brcHub.events.map(e => (
                  <List.Item>
                    <List.Content>
                      <List.Header as="a">Event</List.Header>
                    </List.Content>
                  </List.Item>
                ))
              : "No events defined"}
          </List>
          <Divider horizontal>
            <Header as="h4">
              <Icon name="users" />
              Teams
            </Header>
          </Divider>
          <FakeItemGroup rows={3} />
        </Card.Description>
      </Card.Content>
    </Card>
  ) : (
    <Button onClick={() => register(id, season)}>Register</Button>
  );
};

const BrcHub = ({
  myBrcHubs,
  getBrcHub,
  registerBrc,
  match: {
    params: { id, season }
  }
}) => {
  useEffect(() => {
    (myBrcHubs && myBrcHubs[id]) || getBrcHub(id);
  }, []);

  return (
    <Grid columns={2} centered>
      <Grid.Row>
        <Grid.Column>
          {myBrcHubs && myBrcHubs[id] ? (
            <BrcDescription
              brcHub={myBrcHubs[id].find(s => String(s.ID) === season)}
              id={id}
              season={season}
              register={() => registerBrc(id)}
            />
          ) : (
            "Loading Hub Info ..."
          )}
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = ({ hubsReducer }) => ({
  myBrcHubs: hubsReducer.myBrcHubs
});

const mapDispatchToProps = {
  getBrcHub: id => getBrcHub.request(id),
  registerBrc: (id, season) => registerBrcHub.request(id, season)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(BrcHub);
