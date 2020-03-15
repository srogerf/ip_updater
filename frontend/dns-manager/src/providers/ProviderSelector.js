// dropdown to select provider
import React from "react";
import {Dropdown, DropdownButton, FormControl, InputGroup} from "react-bootstrap";

class ProviderSelector extends React.Component {
  render() {
    return (
      <div>
        <InputGroup className="mb-4">
          <DropdownButton
            as={InputGroup.Append}
            variant="outline-secondary"
            title="Provider"
            id="input-group-dropdown-3"
          >
            <Dropdown.Item href="#">Go Daddy</Dropdown.Item>
          </DropdownButton>
          <FormControl
            placeholder="Provider Name"
            aria-label="Provider Name"
            aria-describedby="basic-addon1"
          />

        </InputGroup>
      </div>
    )
  }
}

export default ProviderSelector;
