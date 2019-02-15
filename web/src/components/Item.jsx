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

const Item = ({ itemName, fields, NewItemForm }) => {
  const [items, setItems] = useState([]);
  const [message, setMessage] = useState(null);
  const [search, setSearch] = useState("");
  const [page, setPage] = useState(0);
  const [newItemModalOpen, setNewItemModalOpen] = useState(false);
  const pageSize = 20;

  useEffect(() => {
    axios
      .get(
        `http://${process.env.REACT_APP_API_URL}/v1/${itemName.toLowerCase()}s`
      )
      .then(response => {
        setItems(response.data);
        setMessage(null);
      })
      .catch(e => {
        setMessage({
          error: true,
          header: `Problem getting ${itemName.toLowerCase()}s`,
          content: e.response.message || e.message
        });
      });
  }, []);

  const deleteItem = id => {
    axios
      .delete(
        `http://${process.env.REACT_APP_API_URL}/v1/${itemName.toLowerCase()}`
      )
      .then(() => {
        setItems(items.filter(i => i.ID !== id));
      })
      .catch(e => console.log(e));
  };

  const addItem = item => {
    setItems([...items, item]);
    setPage(Math.ceil(items.length / pageSize) - 1);
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
            trigger={<Button icon="add" />}
            onOpen={() => setNewItemModalOpen(true)}
            onClose={() => setNewItemModalOpen(false)}
            open={newItemModalOpen}
          >
            <Modal.Header>Add a new {itemName}</Modal.Header>
            <Modal.Content>
              <NewItemForm addToList={addItem} />
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
              <Table.HeaderCell key="delete" />
            </Table.Row>
          </Table.Header>
          {items && items.length ? (
            <Table.Body>
              {items
                .filter(
                  item =>
                    fields
                      .filter(f => f.filter)
                      .filter(f =>
                        (f.displayFn ? f.displayFn(item) : item[f.name])
                          .toLowerCase()
                          .includes(search)
                      ).length
                )
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
                    <Table.Cell collapsing key="delete">
                      <Button icon="trash" onClick={deleteItem(item.ID)} />
                    </Table.Cell>
                  </Table.Row>
                ))}
            </Table.Body>
          ) : (
            <FakeRows cols={4} />
          )}
          <Table.Footer>
            {!search ? (
              <Table.Row>
                <Table.HeaderCell
                  colSpan={fields.filter(f => f.header).length}
                  textAlign="right"
                >
                  <Pagination
                    activePage={page + 1}
                    onPageChange={(_, { activePage }) =>
                      setPage(activePage - 1)
                    }
                    totalPages={Math.ceil(items.length / pageSize)}
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
