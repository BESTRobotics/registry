import React, { useEffect, useState } from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";
import {
  getBrcTeam,
  registerBrcTeam,
  getAllTeams
} from "../../redux/teams/reducer";
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
import NewTeamForm from "../itemForms/NewTeamForm";
import NewBRCTeamForm from "../userForms/NewBRCTeam";

const Team = ({
  team,
  allBrcTeams,
  getBrcTeam,
  token,
  registerBrc,
  allTeams,
  getAllTeams
}) => {
  const [teamModalOpen, setTeamModalOpen] = useState(false);
  const [brcModalOpen, setBrcModalOpen] = useState({});
  useEffect(() => {
    (allBrcTeams && team && allBrcTeams[team.ID]) || getBrcTeam(team.ID);
  }, []);

  return team ? (
    <Grid columns={2} centered>
      <Grid.Row>
        <Grid.Column>
          <Card fluid>
            <Card.Content>
              <Card.Header as={Header} size="huge">
                {team.StaticName}
                <Modal
                  trigger={
                    <span onClick={() => setTeamModalOpen(true)}>
                      <Icon
                        name="pencil"
                        size="small"
                        style={{ cursor: "pointer" }}
                      />
                    </span>
                  }
                  onOpen={() => setTeamModalOpen(true)}
                  onClose={() => setTeamModalOpen(false)}
                  open={!!teamModalOpen}
                >
                  <Modal.Header>New Team</Modal.Header>
                  <Modal.Content>
                    <NewTeamForm
                      addToList={() => setTeamModalOpen(false)}
                      existingItem={team}
                      token={token}
                    />
                  </Modal.Content>
                </Modal>
              </Card.Header>
              <Card.Meta>{team.SchoolName}</Card.Meta>
              <Card.Description>
                <Divider horizontal>
                  <Header as="h4">
                    <Icon name="info circle" />
                    Team Info
                  </Header>
                </Divider>
                <b>Coaches:</b>{" "}
                {team.Coaches
                  ? team.Coaches.map(a => `${a.FirstName} ${a.LastName}`).join(
                      ","
                    )
                  : "none"}
                <br />
                <b>Founded:</b> {team.Founded && team.Founded.substring(0, 10)}
                <br />
                <Divider horizontal>
                  <Header as="h4">
                    <Icon name="calendar" />
                    Seasons
                  </Header>
                </Divider>
                <List divided verticalAlign="middle">
                  {(allBrcTeams &&
                    allBrcTeams[team.ID] &&
                    allBrcTeams[team.ID].map(season => (
                      <List.Item key={season.ID}>
                        <List.Content>
                          <List.Header
                            as={season.brcTeam ? Link : null}
                            to={`/teams/${team.ID}/brc/${season.ID}`}
                          >
                            {season.Name}
                          </List.Header>
                          <List.Description>
                            {season.brcTeam ? (
                              "Registered"
                            ) : season.State === "Open" ? (
                              <Modal
                                open={brcModalOpen[season.ID]}
                                trigger={
                                  <Button compact>
                                    Register Now
                                  </Button>
                                }
                                onOpen={() => setBrcModalOpen({...brcModalOpen, [season.ID]: true})}
                                onClose={() => setBrcModalOpen({...brcModalOpen, [season.ID]: true})}
                              >
                                <Modal.Header>New BRCTeam</Modal.Header>
                                <Modal.Content>
                                  <NewBRCTeamForm
                                    onDone={() => setBrcModalOpen({...brcModalOpen, [season.ID]: false})}
                                    team={team.ID}
                                    season={season.ID}
                                  />
                                </Modal.Content>
                              </Modal>
                            ) : (
                              "Closed"
                            )}
                          </List.Description>
                        </List.Content>
                      </List.Item>
                    ))) || <FakeItemGroup rows={3} />}
                </List>
              </Card.Description>
            </Card.Content>
          </Card>
        </Grid.Column>
      </Grid.Row>
    </Grid>
  ) : (
    <Message>Team not found</Message>
  );
};

const mapStateToProps = ({ teamsReducer, hubReducer }) => ({
  allBrcTeams: teamsReducer.allBrcTeams,
  allTeams: teamsReducer.allTeams,
});

const mapDispatchToProps = {
  getBrcTeam: id => getBrcTeam.request(id),
  getAllTeams: () => getAllTeams.request(),
  registerBrc: (id, season) => registerBrcTeam.request(id, season)
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Team);
