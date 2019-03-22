import React from "react";
import { Item, Placeholder } from "semantic-ui-react";

const FakeItemGroup = ({ rows }) => {
  return (
    <Item.Group divided>
      {Array(rows)
        .fill(0)
        .map(() => (
          <Item>
            <Item.Content>
              <Placeholder>
                <Placeholder.Header>
                  <Placeholder.Line />
                </Placeholder.Header>
                <Placeholder.Paragraph>
                  <Placeholder.Line />
                  <Placeholder.Line />
                </Placeholder.Paragraph>
              </Placeholder>
            </Item.Content>
          </Item>
        ))}
    </Item.Group>
  );
};

export default FakeItemGroup;
