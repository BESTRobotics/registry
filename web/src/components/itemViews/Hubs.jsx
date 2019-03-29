import React, { useEffect } from "react";
import { connect } from "react-redux";
import {
  Grid,
  Card,
  Header,
  Item,
  List,
  Button,
  Icon
} from "semantic-ui-react";
import { getMyHubs, getBrcHub, registerBrc } from "../../redux/hubs/reducer";
import FakeItemGroup from "./FakeItemGroup";
import { Link } from "react-router-dom";

const Hub = ({ hub, expanded, allBrcHubs, getBrcHub, registerBrc }) => {
  useEffect(() => {
    (allBrcHubs && allBrcHubs[hub.ID]) || getBrcHub(hub.ID);
  }, []);
  console.log(allBrcHubs);
  return (
    <Item>
      <Item.Content>
        <Item.Header as={Link} to={`/hubs/${hub.ID}`}>
          {hub.Name}
        </Item.Header>
        <Item.Meta>{hub.Description}</Item.Meta>
        {expanded ? (
          <Item.Description>
            <List>
              {allBrcHubs &&
                allBrcHubs[hub.ID] &&
                allBrcHubs[hub.ID].map(season => (
                  <List.Item key={season.ID}>
                    <Icon name="right triangle" />
                    <List.Content>
                      <List.Header
                        as={season.brcHub ? Link : null}
                        to={`/hubs/${hub.ID}/brc/${season.ID}`}
                      >
                        {season.Name}
                      </List.Header>
                      <List.Description>
                        {season.brcHub ? (
                          "Registered"
                        ) : season.Open ? (
                          <Button
                            as="a"
                            onClick={() => registerBrc(hub.ID, season.ID)}
                          >
                            Register Now
                          </Button>
                        ) : (
                          "Closed"
                        )}
                      </List.Description>
                    </List.Content>
                  </List.Item>
                ))}
            </List>
          </Item.Description>
        ) : null}
      </Item.Content>
    </Item>
  );
};

const Hubs = ({
  hubs,
  getMyHubs,
  match,
  hubsLength,
  allBrcHubs,
  getBrcHub,
  registerBrc
}) => {
  useEffect(() => {
    (hubs && hubs.length) || getMyHubs();
  }, []);
  return (
    <Grid columns={2} centered>
      <Grid.Row>
        <Grid.Column>
          <Card fluid color="orange">
            <Card.Content>
              <Card.Header as={Header} size="huge">
                My Hub{hubs && hubs.length > 1 ? "s" : ""}
              </Card.Header>
              <Card.Description>
                <Item.Group divided>
                  {hubs && hubs.length ? (
                    hubs.map(h => (
                      <Hub
                        key={h.ID}
                        match={match}
                        hub={h}
                        allBrcHubs={allBrcHubs}
                        getBrcHub={getBrcHub}
                        registerBrc={registerBrc}
                        expanded={
                          (match &&
                            match.params &&
                            match.params.id === String(h.ID)) ||
                          hubs.length === 1
                        }
                      />
                    ))
                  ) : (
                    <FakeItemGroup rows={hubsLength} />
                  )}
                </Item.Group>
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = ({ loginReducer, hubsReducer }) => ({
  hubsLength: loginReducer.hubs ? loginReducer.hubs.length : 0,
  hubs: hubsReducer.myHubs,
  allBrcHubs: hubsReducer.allBrcHubs
});

const mapDispatchToProps = {
  getBrcHub: id => getBrcHub.request(id),
  getMyHubs: () => getMyHubs.request(),
  registerBrc: (id, season) => registerBrc.request()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Hubs);
