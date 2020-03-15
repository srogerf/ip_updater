import React from "react";
import ReactDOM from "react-dom";
import GoDaddyDetails from "./providers/GoDaddy";
import ProviderSelector from "./providers/ProviderSelector";
import {
  Button
} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

// render the details input panel
ReactDOM.render(
  <GoDaddyDetails/>,
  document.getElementById('provider_details')
);

ReactDOM.render(
  <ProviderSelector/>,
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
