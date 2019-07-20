import React, { useEffect, useState } from "react";
import axios from "axios";
import {
  Header,
  Grid,
  Table,
  Modal,
  Button,
  Input,
  Pagination,
  Message
} from "semantic-ui-react";
import FakeRows from "./FakeRows";
import PropTypes from "prop-types";

const Item = ({ trashcan, itemName, fields, NewItemForm, token }) => {
  const headers = { authorization: token };
  const [items, setItems] = useState([]);
  const [message, setMessage] = useState(null);
  const [search, setSearch] = useState("");
  const [page, setPage] = useState(0);
  const [newItemModalOpen, setNewItemModalOpen] = useState(null);
  const pageSize = 20;

  useEffect(() => {
    axios
      .get(
        `${process.env.REACT_APP_API_URL}/v1/${itemName.toLowerCase()}s`
      )
      .then(response => {
        if (response.data) {
          setItems(response.data);
          setMessage();
        }
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem getting ${itemName.toLowerCase()}s`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  }, []);

  const deactivateItem = id => {
    axios
      .put(
        `${
          process.env.REACT_APP_API_URL
        }/v1/${itemName.toLowerCase()}s/${id}/${trashcan}`,
        {},
        { headers: headers }
      )
      .then(() => {
        setItems(items.filter(i => i.ID !== id));
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem deleting ${itemName.toLowerCase()}s`,
          content:
            e.response && e.response.data ? e.response.data.Message : e.message
        });
      });
  };

  const addItem = item => {
    const existingItem = items.findIndex(i => i.ID === item.ID);
    if (existingItem === -1) {
      setPage(Math.floor((items.length + 1) / pageSize));
    }
    setItems(
      existingItem !== -1
        ? [
            ...items.slice(0, existingItem),
            item,
            ...items.slice(existingItem + 1)
          ]
        : [...items, item]
    );
    setNewItemModalOpen(false);
  };

  return (
    <Grid padded>
      <Grid.Row centered>{message ? <Message {...message} /> : null}</Grid.Row>
      <Grid.Row centered>
        <Header>{itemName} Management</Header>
      </Grid.Row>
      <Grid.Row centered>
        <Grid.Column width={3}>
          <Input
            icon="search"
            fluid
            placeholder="Search..."
            value={search}
            onChange={(_, { value }) => setSearch(value.toLowerCase())}
          />
        </Grid.Column>
        <Grid.Column width={2}>
          <Modal
            closeOnDimmerClick={false}
            trigger={<Button icon="add" />}
            onOpen={() => setNewItemModalOpen(true)}
            onClose={() => setNewItemModalOpen(null)}
            open={!!newItemModalOpen}
          >
            <Modal.Header>
              {newItemModalOpen && newItemModalOpen.ID ? "Edit" : "Add a new"}{" "}
              {itemName}
            </Modal.Header>
            <Modal.Content>
              <NewItemForm
                addToList={addItem}
                existingItem={
                  newItemModalOpen === true ? null : newItemModalOpen
                }
                token={token}
              />
            </Modal.Content>
          </Modal>
        </Grid.Column>
      </Grid.Row>
      <Grid.Row>
        <Table>
          <Table.Header>
            <Table.Row>
              {fields
                .filter(f => !!f.header)
                .map(f => (
                  <Table.HeaderCell key={f.header}>{f.header}</Table.HeaderCell>
                ))}
              <Table.HeaderCell key="edit" />
              {trashcan ? <Table.HeaderCell key="deactivate" /> : null}
            </Table.Row>
          </Table.Header>
          {items && items.length ? (
            <Table.Body>
              {items
                .filter(item => {
                  return fields
                    .filter(f => f.filter)
                    .filter(f => {
                      console.log(f);
                      return (f.displayFn ? f.displayFn(item) : item[f.name])
                        .toLowerCase()
                        .includes(search);
                    }).length;
                })
                .slice(page * pageSize, (page + 1) * pageSize)
                .map(item => (
                  <Table.Row key={item.ID}>
                    {fields
                      .filter(f => f.header)
                      .map(f =>
                        f.displayFn ? f.displayFn(item) : item[f.name]
                      )
                      .map((f, i) => (
                        <Table.Cell key={i}>{f}</Table.Cell>
                      ))}
                    <Table.Cell collapsing key="edit">
                      <Button
                        icon="edit"
                        onClick={() => setNewItemModalOpen(item)}
                      />
                    </Table.Cell>
                    {trashcan ? (
                      <Table.Cell collapsing key="deactivate">
                        <Button
                          icon="trash"
                          onClick={() => deactivateItem(item.ID)}
                        />
                      </Table.Cell>
                    ) : null}
                  </Table.Row>
                ))}
            </Table.Body>
          ) : (
            <FakeRows
              cols={fields.filter(f => f.header).length + (!!trashcan ? 2 : 1)}
            />
          )}
          <Table.Footer>
            {!search ? (
              <Table.Row>
                <Table.HeaderCell
                  colSpan={
                    fields.filter(f => f.header).length + (!!trashcan ? 2 : 1)
                  }
                  textAlign="right"
                >
                  <Pagination
                    activePage={page + 1}
                    onPageChange={(_, { activePage }) =>
                      setPage(activePage - 1)
                    }
                    totalPages={Math.ceil(items.length / pageSize) || 1}
                    prevItem={null}
                    nextItem={null}
                  />
                </Table.HeaderCell>
              </Table.Row>
            ) : null}
          </Table.Footer>
        </Table>
      </Grid.Row>
    </Grid>
  );
};

export default Item;

Item.propTypes = {
  itemName: PropTypes.string.isRequired, // Capital and Singular
  fields: PropTypes.arrayOf(
    PropTypes.shape({
      header: PropTypes.string, // The header if it should appear in the table
      name: PropTypes.string, // The field as it appears in the network request
      displayFn: PropTypes.func, // If displaying in table should be more than item[name].  Should take an item and return a string
      filter: PropTypes.bool // Whether the search box should consider this field in searching
    })
  ).isRequired,
  NewItemForm: PropTypes.func.isRequired
};
