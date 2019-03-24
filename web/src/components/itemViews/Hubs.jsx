import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Grid, Card, Header, Item } from "semantic-ui-react";
import { getMyHubs } from "../../redux/hubs/reducer";
import { NavLink } from "react-router-dom";
import BrcHub from "./BrcHub";

const Hub = ({ hub, expanded }) => {
  return (
    <Item as={NavLink} to={`/hubs/${hub.ID}`}>
      <Item.Content>
        <Item.Header>{hub.Name}</Item.Header>
        <Item.Meta>{hub.Description}</Item.Meta>
        {expanded ? (
          <Item.Description>
            <BrcHub id={hub.ID} />
          </Item.Description>
        ) : null}
      </Item.Content>
    </Item>
  );
};

const Hubs = ({ hubs, getMyHubs, match }) => {
  useEffect(() => {
    (hubs && hubs.length) || getMyHubs();
  }, []);
  return (
    <Grid.Row>
      <Grid.Column>
        <Card fluid color="orange">
          <Card.Content>
            <Card.Header as={Header} size="huge">
              My Hub{hubs && hubs.length > 1 ? "s" : ""}
            </Card.Header>
            <Card.Description>
              <Item.Group divided>
                {hubs && hubs.length
                  ? hubs.map(h => (
                      <Hub
                        key={h.ID}
                        match={match}
                        hub={h}
                        expanded={
                          (match.params && match.params.id === "8") ||
                          hubs.length === 1
                        }
                      />
                    ))
                  : "Loading"}
              </Item.Group>
            </Card.Description>
          </Card.Content>
        </Card>
      </Grid.Column>
    </Grid.Row>
  );
};

const mapStateToProps = ({ loginReducer, hubsReducer }) => ({
  hubsLength: loginReducer.hubs ? loginReducer.hubs.length : 0,
  hubs: hubsReducer.myHubs
});

const mapDispatchToProps = {
  getMyHubs: () => getMyHubs.request()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Hubs);
