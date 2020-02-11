import React from "react";
import ReactDOM from "react-dom";
import GoDaddyDetails from "./providers/GoDaddy";
import {
  Button,
  Dropdown,
  InputGroup,
  FormControl,
  DropdownButton
} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';


ReactDOM.render(
  <GoDaddyDetails/>,
  document.getElementById('provider_details')
);

class Provider extends React.Component {
  render() {
    return (
      <div>
        <InputGroup className="mb-3">
          <DropdownButton
            as={InputGroup.Append}
            variant="outline-secondary"
            title="Provider"
            id="input-group-dropdown-2"
          >
            <Dropdown.Item href="#">Go Daddy</Dropdown.Item>
          </DropdownButton>
          <FormControl
            placeholder="Provider Name"
            aria-label="Provider Name"
            aria-describedby="basic-addon2"
          />

        </InputGroup>
      </div>
    )
  }
}

ReactDOM.render(
  <Provider/>,
  document.getElementById('provider_select')
);


/*
  submit data
 */
class SendData extends React.Component {
  render() {
    return (
      <Button
        variant="primary"
      >Submit
      </Button>
    )

  }
}

ReactDOM.render(
  <SendData/>,
  document.getElementById('submit')
);
