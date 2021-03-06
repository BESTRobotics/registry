import React, { useEffect, useState } from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";
import { getBrcHub, registerBrcHub } from "../../redux/hubs/reducer";
import { getAllTeams } from "../../redux/teams/reducer";
import {
  Button,
  Icon,
  Header,
  Card,
  Grid,
  Divider,
  List,
  Modal,
  Message
} from "semantic-ui-react";
import FakeItemGroup from "./FakeItemGroup";
import NewHubForm from "../itemForms/NewHubForm";

const Hub = ({
  hub,
  myBrcHubs,
  getBrcHub,
  token,
  registerBrc,
  allTeams,
  getAllTeams
}) => {
  const [hubModalOpen, setHubModalOpen] = useState(false);
  useEffect(() => {
    if (hub) {
      (myBrcHubs && hub && myBrcHubs[hub.ID]) || getBrcHub(hub.ID);
    }
  }, [hub]);
  useEffect(() => {
    (allTeams && allTeams.length) || getAllTeams();
  }, []);

  return hub ? (
    <Grid columns={2} centered>
      <Grid.Row>
        <Grid.Column>
          <Card fluid>
            <Card.Content>
              <Card.Header as={Header} size="huge">
                {hub.Name}
                <Modal
                  closeOnDimmerClick={false}
                  trigger={
                    <span onClick={() => setHubModalOpen(true)}>
                      <Icon
                        name="pencil"
                        size="small"
                        style={{ cursor: "pointer" }}
                      />
                    </span>
                  }
                  onOpen={() => setHubModalOpen(true)}
                  onClose={() => setHubModalOpen(false)}
                  open={!!hubModalOpen}
                >
                  <Modal.Header>Edit Hub</Modal.Header>
                  <Modal.Content>
                    <NewHubForm
                      addToList={() => setHubModalOpen(false)}
                      existingItem={hub}
                      token={token}
                    />
                  </Modal.Content>
                </Modal>
              </Card.Header>
              <Card.Meta>{hub.Description}</Card.Meta>
              <Card.Description>
                <Divider horizontal>
                  <Header as="h4">
                    <Icon name="info circle" />
                    Hub Info
                  </Header>
                </Divider>
                <b>Director:</b> {hub.Director.FirstName}{" "}
                {hub.Director.LastName}<br/>
                <b>Director Email:</b> <a href={`mailto:${hub.Director.EMail}`}>{hub.Director.EMail}</a>
                <br />
                <b>Admins:</b>{" "}
                {hub.Admins
                  ? hub.Admins.map(a => `${a.FirstName} ${a.LastName}`).join(
                      ","
                    )
                  : "none"}
                <br />
                <b>Location:</b> {hub.Location}
                <br />
                <b>Founded:</b> {hub.Founded && hub.Founded.substring(0, 10)}
                <br />
                <Divider horizontal>
                  <Header as="h4">
                    <Icon name="calendar" />
                    Seasons
                  </Header>
                </Divider>
                <List divided verticalAlign="middle">
                  {(myBrcHubs &&
                    myBrcHubs[hub.ID] &&
                    myBrcHubs[hub.ID].map(season => (
                      <List.Item key={season.ID}>
                        <List.Content>
                          <List.Header
                            as={season.brcHub ? Link : null}
                            to={`/hubs/${hub.ID}/brc/${season.ID}`}
                          >
                            {season.Name}
                          </List.Header>
                          <List.Description>
                            {season.brcHub ? (
                              season.brcHub.Meta.BRIApproved ? (
                                "Approved"
                              ) : (
                                "Pending Approval"
                              )
                            ) : season.State === "Open" ? (
                              <Button
                                compact
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
                    ))) || <FakeItemGroup rows={3} />}
                </List>
                <Divider horizontal>
                  <Header as="h4">
                    <Icon name="users" />
                    Teams
                  </Header>
                </Divider>
                {(allTeams &&
                  allTeams.length &&
                  allTeams
                    .filter(t => t.HomeHub.ID === hub.ID)
                    .map(team => (
                      <List.Item key={team.ID}>
                        <List.Content>
                          <List.Header as={Link} to={`/teams/${team.ID}`}>
                            {team.StaticName}
                          </List.Header>
                          <List.Description>{team.SchoolName}</List.Description>
                        </List.Content>
                      </List.Item>
                    ))) || <FakeItemGroup rows={3} />}
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  ) : (
    <Message>Hub not found</Message>
  );
};

const mapStateToProps = ({ hubsReducer, teamsReducer, loginReducer }) => ({
  myBrcHubs: hubsReducer.myBrcHubs,
  allTeams: teamsReducer.allTeams,
  token: loginReducer.token
});

const mapDispatchToProps = {
  getBrcHub: id => getBrcHub.request(id),
  getAllTeams: () => getAllTeams.request(),
  registerBrc: (id, season) => registerBrcHub.request(id, season)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Hub);
