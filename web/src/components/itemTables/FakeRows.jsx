import React from "react";
import { Table, Placeholder } from "semantic-ui-react";

const FakeRows = ({ cols }) => {
  return (
    <Table.Body>
      {Array(4)
          .fill()
          .map((_, i) => (
            <Table.Row key={i}>
              {Array(cols)
                .fill()
                .map((_, j) => (
                  <Table.Cell key={j}>
                    <Placeholder>
                      <Placeholder.Header>
                        <Placeholder.Line />
                      </Placeholder.Header>
                    </Placeholder>
                  </Table.Cell>
                ))}
              </Table.Row>
          ))}
        </Table.Body>
  );
};

export default FakeRows;
