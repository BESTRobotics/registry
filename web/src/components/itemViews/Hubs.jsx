import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Grid, Card, Header, Item, Icon } from "semantic-ui-react";
import { getMyHubs } from "../../redux/hubs/reducer";

const Hubs = ({ hubs, getMyHubs }) => {
  useEffect(() => {
    (hubs && hubs.length) || getMyHubs();
  }, []);

  return (
    <Grid.Row>
      <Grid.Column>
        <Card fluid color="orange">
          <Card.Content>
            <Card.Header as={Header} size="huge">
              My Hubs
            </Card.Header>
            <Card.Description>
              <Item.Group divided>
                {hubs && hubs.length
                  ? hubs.map(h => (
                      <Item key={h.ID}>
                        <Item.Content>
                          <Item.Header>{h.Name}</Item.Header>
                          <Item.Meta>{h.Description}</Item.Meta>
                          <Item.Description />
                        </Item.Content>
                      </Item>
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
