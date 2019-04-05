import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Grid, Card, Header, Item } from "semantic-ui-react";
import { getMyHubs, getBrcHub, registerBrcHub } from "../../redux/hubs/reducer";
import FakeItemGroup from "./FakeItemGroup";
import { Link } from "react-router-dom";
import Hub from "./Hub.jsx";

const Hubs = ({ hubs, getMyHubs, hubsLength, match }) => {
  useEffect(() => {
    (hubs && hubs.length) || getMyHubs();
  }, []);
  if (hubs && hubs.length) {
    if (match && match.params && match.params.id) {
      const hub = hubs.find(h => String(h.ID) === match.params.id);
      return <Hub hub={hub} />;
    }
    if (hubs.length === 1) {
      return <Hub hub={hubs[0]} />;
    }
  }
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
                    hubs.map(hub => (
                      <Item>
                        <Item.Content>
                          <Item.Header as={Link} to={`/hubs/${hub.ID}`}>
                            {hub.Name}
                          </Item.Header>
                          <Item.Meta>{hub.Description}</Item.Meta>
                        </Item.Content>
                      </Item>
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
  registerBrc: (id, season) => registerBrcHub.request(id, season)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Hubs);
