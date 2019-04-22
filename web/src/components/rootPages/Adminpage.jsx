import React, { useEffect } from "react";
import { connect } from "react-redux";
import { Grid, Card, Header, Item, Button } from "semantic-ui-react";
import {
  getBrcHub,
  registerBrcHub,
  getAllHubs,
  getSeasonBrcHubs,
  getSeasons,
  approveBrcHub
} from "../../redux/hubs/reducer";
import FakeItemGroup from "../itemViews/FakeItemGroup";
import { Link } from "react-router-dom";

const AdminPage = ({
  brcHubs,
  seasons,
  getSeasonBrcHubs,
  getSeasons,
  approveBrcHub
}) => {
  useEffect(() => {
    (seasons && seasons.length) || getSeasons();
    seasons &&
      seasons.forEach(season => {
        (brcHubs && brcHubs[season.ID]) || getSeasonBrcHubs(season.ID);
      });
  }, [seasons]);
  console.log(seasons);
  return (
    <Grid columns={2} centered>
      <Grid.Row>
        <Grid.Column>
          {brcHubs ? (
            Object.keys(brcHubs)
              .reverse()
              .map(season => (
                <Card fluid color="orange" key={season}>
                  <Card.Content>
                    <Card.Header as={Header} size="huge">
                      {seasons.find(s => s.ID.toString() === season).Name}
                    </Card.Header>
                    <Card.Description>
                      <Item.Group divided>
                        {brcHubs[season].map(brcHub => {
                          return (
                            <Item key={brcHub.Hub.ID}>
                              <Item.Content>
                                <Item.Header
                                  as={Link}
                                  to={`/hubs/${brcHub.Hub.ID}`}
                                >
                                  {brcHub.Hub.Name}
                                </Item.Header>
                                <Item.Meta>{brcHub.Hub.Location}</Item.Meta>
                                {brcHub.Meta.BRIApproved ? (
                                  "Approved"
                                ) : (
                                  <Button
                                    onClick={() =>
                                      approveBrcHub(
                                        brcHub.Hub.ID,
                                        brcHub.Season.ID
                                      )
                                    }
                                  >
                                    Approve
                                  </Button>
                                )}
                              </Item.Content>
                            </Item>
                          );
                        })}
                      </Item.Group>
                    </Card.Description>
                  </Card.Content>
                </Card>
              ))
          ) : (
            <FakeItemGroup rows={3} />
          )}
        </Grid.Column>
      </Grid.Row>
    </Grid>
  );
};

const mapStateToProps = ({ loginReducer, hubsReducer }) => ({
  brcHubs: hubsReducer.allBrcHubs,
  seasons: hubsReducer.seasons
});

const mapDispatchToProps = {
  getSeasonBrcHubs: season => getSeasonBrcHubs.request(season),
  approveBrcHub: (id, season) => approveBrcHub.request(id, season),
  getSeasons: () => getSeasons.request()
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AdminPage);
